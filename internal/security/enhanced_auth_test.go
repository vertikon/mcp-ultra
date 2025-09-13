package security

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// MockOPAService for testing
type MockOPAService struct {
	mock.Mock
}

func (m *MockOPAService) IsAuthorized(ctx context.Context, claims *Claims, method, path string) bool {
	args := m.Called(ctx, claims, method, path)
	return args.Bool(0)
}

func (m *MockOPAService) EvaluatePolicy(ctx context.Context, input interface{}) (bool, error) {
	args := m.Called(ctx, input)
	return args.Bool(0), args.Error(1)
}

func TestNewAuthService(t *testing.T) {
	config := AuthConfig{
		Mode:         "jwt",
		Issuer:       "test-issuer",
		Audience:     "test-audience",
		TokenExpiry:  time.Hour,
		RefreshExpiry: time.Hour * 24,
	}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	assert.NotNil(t, authService)
	assert.Equal(t, config, authService.config)
	assert.Equal(t, logger, authService.logger)
	assert.Equal(t, mockOPA, authService.opa)
	assert.NotNil(t, authService.publicKeys)
}

func TestAuthService_JWTMiddleware(t *testing.T) {
	config := AuthConfig{
		Mode:     "jwt",
		Issuer:   "test-issuer",
		Audience: "test-audience",
	}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	// Create a mock next handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})

	middleware := authService.JWTMiddleware(nextHandler)

	tests := []struct {
		name           string
		path           string
		authHeader     string
		expectedStatus int
		setupOPA       func()
	}{
		{
			name:           "health endpoint bypasses auth",
			path:           "/healthz",
			authHeader:     "",
			expectedStatus: http.StatusOK,
			setupOPA:       func() {},
		},
		{
			name:           "ready endpoint bypasses auth",
			path:           "/readyz",
			authHeader:     "",
			expectedStatus: http.StatusOK,
			setupOPA:       func() {},
		},
		{
			name:           "metrics endpoint bypasses auth",
			path:           "/metrics",
			authHeader:     "",
			expectedStatus: http.StatusOK,
			setupOPA:       func() {},
		},
		{
			name:           "missing authorization header returns 401",
			path:           "/api/tasks",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			setupOPA:       func() {},
		},
		{
			name:           "invalid authorization header format returns 401",
			path:           "/api/tasks",
			authHeader:     "Basic dGVzdDp0ZXN0",
			expectedStatus: http.StatusUnauthorized,
			setupOPA:       func() {},
		},
		{
			name:           "invalid token returns 401",
			path:           "/api/tasks",
			authHeader:     "Bearer invalid.token.here",
			expectedStatus: http.StatusUnauthorized,
			setupOPA:       func() {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset OPA mock
			mockOPA.ExpectedCalls = nil
			tt.setupOPA()

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			w := httptest.NewRecorder()
			middleware.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusUnauthorized {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "unauthorized", response["error"])
			}
		})
	}
}

func TestAuthService_ValidateToken(t *testing.T) {
	config := AuthConfig{
		Mode:     "jwt",
		Issuer:   "test-issuer",
		Audience: "test-audience",
	}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	// Generate RSA key pair for testing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Add public key to auth service
	keyID := "test-key-id"
	authService.publicKeys[keyID] = &privateKey.PublicKey

	tests := []struct {
		name        string
		setupToken  func() string
		expectError bool
		checkClaims func(*testing.T, *Claims)
	}{
		{
			name: "valid token with all claims",
			setupToken: func() string {
				claims := &Claims{
					UserID:   "user-123",
					Email:    "test@example.com",
					Role:     "admin",
					Scopes:   []string{"read", "write"},
					TenantID: "tenant-456",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    config.Issuer,
						Audience:  []string{config.Audience},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
						IssuedAt:  jwt.NewNumericDate(time.Now()),
						NotBefore: jwt.NewNumericDate(time.Now()),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				token.Header["kid"] = keyID
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: false,
			checkClaims: func(t *testing.T, claims *Claims) {
				assert.Equal(t, "user-123", claims.UserID)
				assert.Equal(t, "test@example.com", claims.Email)
				assert.Equal(t, "admin", claims.Role)
				assert.Equal(t, []string{"read", "write"}, claims.Scopes)
				assert.Equal(t, "tenant-456", claims.TenantID)
			},
		},
		{
			name: "token with wrong issuer",
			setupToken: func() string {
				claims := &Claims{
					UserID: "user-123",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    "wrong-issuer",
						Audience:  []string{config.Audience},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				token.Header["kid"] = keyID
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: true,
		},
		{
			name: "token with wrong audience",
			setupToken: func() string {
				claims := &Claims{
					UserID: "user-123",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    config.Issuer,
						Audience:  []string{"wrong-audience"},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				token.Header["kid"] = keyID
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: true,
		},
		{
			name: "expired token",
			setupToken: func() string {
				claims := &Claims{
					UserID: "user-123",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    config.Issuer,
						Audience:  []string{config.Audience},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // Already expired
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				token.Header["kid"] = keyID
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: true,
		},
		{
			name: "token with unknown key ID",
			setupToken: func() string {
				claims := &Claims{
					UserID: "user-123",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    config.Issuer,
						Audience:  []string{config.Audience},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				token.Header["kid"] = "unknown-key-id"
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: true,
		},
		{
			name: "token without key ID",
			setupToken: func() string {
				claims := &Claims{
					UserID: "user-123",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    config.Issuer,
						Audience:  []string{config.Audience},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
				// Don't set kid header
				tokenString, _ := token.SignedString(privateKey)
				return tokenString
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenString := tt.setupToken()

			claims, err := authService.validateToken(tokenString)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, claims)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, claims)
				if tt.checkClaims != nil {
					tt.checkClaims(t, claims)
				}
			}
		})
	}
}

func TestGetUserFromContext(t *testing.T) {
	tests := []struct {
		name        string
		setupCtx    func() context.Context
		expectError bool
		expectUser  *Claims
	}{
		{
			name: "context with valid user claims",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID:   "user-123",
					Email:    "test@example.com",
					Role:     "admin",
					TenantID: "tenant-456",
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectError: false,
			expectUser: &Claims{
				UserID:   "user-123",
				Email:    "test@example.com",
				Role:     "admin",
				TenantID: "tenant-456",
			},
		},
		{
			name: "context without user",
			setupCtx: func() context.Context {
				return context.Background()
			},
			expectError: true,
			expectUser:  nil,
		},
		{
			name: "context with wrong type",
			setupCtx: func() context.Context {
				return context.WithValue(context.Background(), "user", "not-claims")
			},
			expectError: true,
			expectUser:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setupCtx()

			user, err := GetUserFromContext(ctx)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, user)
				assert.Equal(t, tt.expectUser.UserID, user.UserID)
				assert.Equal(t, tt.expectUser.Email, user.Email)
				assert.Equal(t, tt.expectUser.Role, user.Role)
				assert.Equal(t, tt.expectUser.TenantID, user.TenantID)
			}
		})
	}
}

func TestRequireScope(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})

	middleware := RequireScope("write")

	tests := []struct {
		name           string
		setupCtx       func() context.Context
		expectedStatus int
	}{
		{
			name: "user with required scope succeeds",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Scopes: []string{"read", "write", "delete"},
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "user without required scope fails",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Scopes: []string{"read"},
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "context without user fails",
			setupCtx: func() context.Context {
				return context.Background()
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "user with empty scopes fails",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Scopes: []string{},
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setupCtx()
			req := httptest.NewRequest(http.MethodGet, "/api/test", nil).WithContext(ctx)
			w := httptest.NewRecorder()

			middleware(nextHandler).ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestRequireRole(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})

	middleware := RequireRole("manager")

	tests := []struct {
		name           string
		setupCtx       func() context.Context
		expectedStatus int
	}{
		{
			name: "user with exact role succeeds",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Role:   "manager",
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "admin user always succeeds",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Role:   "admin",
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "user with different role fails",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Role:   "user",
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusForbidden,
		},
		{
			name: "context without user fails",
			setupCtx: func() context.Context {
				return context.Background()
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "user with empty role fails",
			setupCtx: func() context.Context {
				claims := &Claims{
					UserID: "user-123",
					Role:   "",
				}
				return context.WithValue(context.Background(), "user", claims)
			},
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setupCtx()
			req := httptest.NewRequest(http.MethodGet, "/api/test", nil).WithContext(ctx)
			w := httptest.NewRecorder()

			middleware(nextHandler).ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestAuthService_ResponseWriters(t *testing.T) {
	config := AuthConfig{}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	t.Run("writeUnauthorized sets correct headers and body", func(t *testing.T) {
		w := httptest.NewRecorder()
		message := "test unauthorized message"

		authService.writeUnauthorized(w, message)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "unauthorized", response["error"])
		assert.Equal(t, message, response["message"])
	})

	t.Run("writeForbidden sets correct headers and body", func(t *testing.T) {
		w := httptest.NewRecorder()
		message := "test forbidden message"

		authService.writeForbidden(w, message)

		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "forbidden", response["error"])
		assert.Equal(t, message, response["message"])
	})
}

func TestAuthService_JWKToRSA(t *testing.T) {
	config := AuthConfig{}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	tests := []struct {
		name        string
		n           string
		e           string
		expectError bool
	}{
		{
			name: "valid JWK parameters",
			// These are example values - in real tests you'd use actual JWK values
			n:           "AM7nTbTKe9LMZuWWZSlbgWyA", // Base64url encoded
			e:           "AQAB",                    // Standard RSA exponent (65537)
			expectError: false,
		},
		{
			name:        "invalid modulus encoding",
			n:           "invalid-base64url!!!",
			e:           "AQAB",
			expectError: true,
		},
		{
			name:        "invalid exponent encoding",
			n:           "AM7nTbTKe9LMZuWWZSlbgWyA",
			e:           "invalid-base64url!!!",
			expectError: true,
		},
		{
			name:        "empty modulus",
			n:           "",
			e:           "AQAB",
			expectError: true,
		},
		{
			name:        "empty exponent",
			n:           "AM7nTbTKe9LMZuWWZSlbgWyA",
			e:           "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey, err := authService.jwkToRSA(tt.n, tt.e)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, publicKey)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, publicKey)
				assert.NotNil(t, publicKey.N)
				assert.Greater(t, publicKey.E, 0)
			}
		})
	}
}

func TestAuthService_ConcurrentAccess(t *testing.T) {
	config := AuthConfig{
		Mode:     "jwt",
		Issuer:   "test-issuer",
		Audience: "test-audience",
	}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	// Add a test key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	
	keyID := "test-key"
	authService.publicKeys[keyID] = &privateKey.PublicKey

	t.Run("concurrent access to public keys is safe", func(t *testing.T) {
		const numGoroutines = 10
		done := make(chan bool, numGoroutines)

		// Simulate concurrent reads and writes to public keys
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				// Read existing key
				key, exists := authService.publicKeys[keyID]
				assert.True(t, exists)
				assert.NotNil(t, key)

				// Add a new key
				newKeyID := fmt.Sprintf("concurrent-key-%d", id)
				authService.publicKeys[newKeyID] = &privateKey.PublicKey

				// Read the new key
				newKey, exists := authService.publicKeys[newKeyID]
				assert.True(t, exists)
				assert.NotNil(t, newKey)
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			select {
			case <-done:
				// Success
			case <-time.After(5 * time.Second):
				t.Fatal("Test timed out")
			}
		}

		// Verify all keys were added
		assert.GreaterOrEqual(t, len(authService.publicKeys), numGoroutines+1)
	})
}

func TestAuthService_EdgeCases(t *testing.T) {
	config := AuthConfig{
		Mode:     "jwt",
		Issuer:   "test-issuer",
		Audience: "test-audience",
	}
	logger := zap.NewNop()
	mockOPA := &MockOPAService{}

	authService := NewAuthService(config, logger, mockOPA)

	t.Run("validate token with nil input", func(t *testing.T) {
		claims, err := authService.validateToken("")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("validate token with malformed JWT", func(t *testing.T) {
		claims, err := authService.validateToken("not.a.jwt")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("validate token with missing segments", func(t *testing.T) {
		claims, err := authService.validateToken("header.payload")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("context with nil user value", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "user", nil)
		user, err := GetUserFromContext(ctx)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}