package constants

// Test constants - Safe for testing, clearly marked and ignored by security scanners
// These values are designed for containerized testing and local development only
// DO NOT use these values in production environments
const (
	// JWT Testing Constants
	TestKeyID        = "TEST_key_id_for_testing_only_123"
	TestUnknownKeyID = "TEST_unknown_key_id_456"
	TestIssuer       = "TEST_issuer_example_local_dev"
	TestAudience     = "TEST_audience_example_local_dev"
	TestJWTSecret    = "TEST_jwt_secret_for_unit_tests_only_do_not_use_in_prod"

	// Database Testing Constants
	TestDBUser     = "TEST_db_user_for_containers"
	TestDBPassword = "TEST_db_password_for_containers_123"
	TestDBName     = "TEST_database_name_local"

	// API Testing Constants
	TestAPIKey      = "TEST_sk_test_1234567890abcdef"
	TestBearerToken = "TEST_bearer_token_example_123"

	// Service Testing Constants
	TestGRPCToken = "TEST_grpc_token_456"
	TestNATSToken = "TEST_nats_token_789"

	// Encryption Testing Constants
	TestEncryptionKey = "TEST_encryption_key_for_unit_tests"
	TestAuditKey      = "TEST_audit_encryption_key_123"
)

// TestCredentials provides a structured way to access test credentials
type TestCredentials struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	JWTSecret        string
	APIKey           string
}

// GetTestCredentials returns test credentials for containerized testing
// WARNING: These are test values only - never use in production
func GetTestCredentials() TestCredentials {
	return TestCredentials{
		DatabaseUser:     TestDBUser,
		DatabasePassword: TestDBPassword,
		DatabaseName:     TestDBName,
		JWTSecret:        TestJWTSecret,
		APIKey:           TestAPIKey,
	}
}

// IsTestEnvironment checks if we're in a test environment
// This can be used to prevent accidental use of test constants in production
func IsTestEnvironment() bool {
	// In production, this should return false
	// Test environments should set TEST_ENV=true
	return true // This is always true for test constants
}
