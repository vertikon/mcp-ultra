package compliance

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func createTestComplianceFramework(t *testing.T) *ComplianceFramework {
	config := ComplianceConfig{
		Enabled:       true,
		DefaultRegion: "BR",
		PIIDetection: PIIDetectionConfig{
			Enabled:           true,
			ScanFields:        []string{"email", "phone", "cpf", "name"},
			ClassificationAPI: "local",
			Confidence:        0.8,
			AutoMask:          true,
		},
		Consent: ConsentConfig{
			Enabled:         true,
			DefaultPurposes: []string{"processing", "analytics"},
			TTL:             2 * time.Hour,
			GranularLevel:   "field",
		},
		DataRetention: DataRetentionConfig{
			Enabled:       true,
			DefaultPeriod: 365 * 24 * time.Hour, // 1 year
			CategoryPeriods: map[string]time.Duration{
				"personal": 2 * 365 * 24 * time.Hour, // 2 years
				"session":  30 * 24 * time.Hour,      // 30 days
			},
			AutoDelete:      true,
			BackupRetention: 7 * 365 * 24 * time.Hour, // 7 years
		},
		AuditLogging: AuditLoggingConfig{
			Enabled:           true,
			DetailLevel:       "full",
			RetentionPeriod:   5 * 365 * 24 * time.Hour, // 5 years
			EncryptionEnabled: true,
			ExternalLogging:   false,
			ExternalEndpoint:  "",
		},
		LGPD: LGPDConfig{
			Enabled:          true,
			DPOContact:       "dpo@example.com",
			LegalBasis:       "consent",
			DataCategories:   []string{"personal", "sensitive"},
			SharedThirdParty: []string{"analytics-provider"},
		},
		GDPR: GDPRConfig{
			Enabled:             true,
			DPOContact:          "dpo@example.com",
			LegalBasis:          "consent",
			DataCategories:      []string{"personal", "sensitive"},
			CrossBorderTransfer: true,
			AdequacyDecisions:   []string{"US", "CA"},
		},
		Anonymization: AnonymizationConfig{
			Enabled:    true,
			Methods:    []string{"hashing", "generalization"},
			HashSalt:   "test-salt",
			Reversible: false,
			KAnonymity: 5,
			Algorithms: map[string]string{
				"email": "hash",
				"phone": "mask",
			},
		},
		DataRights: DataRightsConfig{
			Enabled:              true,
			ResponseTime:         30 * 24 * time.Hour, // 30 days
			AutoFulfillment:      false,
			VerificationRequired: true,
			NotificationChannels: []string{"email", "sms"},
		},
	}

	logger := zaptest.NewLogger(t)
	framework, err := NewComplianceFramework(config, logger)
	require.NoError(t, err)
	require.NotNil(t, framework)

	return framework
}

func TestComplianceFramework_Creation(t *testing.T) {
	framework := createTestComplianceFramework(t)
	assert.NotNil(t, framework)
	assert.True(t, framework.config.Enabled)
	assert.Equal(t, "BR", framework.config.DefaultRegion)
}

func TestComplianceFramework_PIIDetection(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	testData := map[string]interface{}{
		"name":  "João Silva",
		"email": "joao@example.com",
		"phone": "+5511999999999",
		"cpf":   "123.456.789-00",
		"age":   30,
	}

	result, err := framework.ScanForPII(ctx, testData)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Should detect PII fields
	assert.Contains(t, result.DetectedFields, "email")
	assert.Contains(t, result.DetectedFields, "phone")
	assert.Contains(t, result.DetectedFields, "name")
	assert.Contains(t, result.DetectedFields, "cpf")

	// Age should not be detected as PII
	assert.NotContains(t, result.DetectedFields, "age")
}

func TestComplianceFramework_ConsentManagement(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()
	purposes := []string{"processing", "analytics"}

	// Record consent
	err := framework.RecordConsent(ctx, userID, purposes, "web")
	assert.NoError(t, err)

	// Check consent
	hasConsent, err := framework.HasConsent(ctx, userID, "processing")
	assert.NoError(t, err)
	assert.True(t, hasConsent)

	// Check consent for ungranted purpose
	hasConsent, err = framework.HasConsent(ctx, userID, "marketing")
	assert.NoError(t, err)
	assert.False(t, hasConsent)

	// Withdraw consent
	err = framework.WithdrawConsent(ctx, userID, []string{"analytics"})
	assert.NoError(t, err)

	// Verify consent withdrawn
	hasConsent, err = framework.HasConsent(ctx, userID, "analytics")
	assert.NoError(t, err)
	assert.False(t, hasConsent)

	// Processing consent should still exist
	hasConsent, err = framework.HasConsent(ctx, userID, "processing")
	assert.NoError(t, err)
	assert.True(t, hasConsent)
}

func TestComplianceFramework_DataRetention(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()
	dataCategory := "personal"

	// Record data creation
	err := framework.RecordDataCreation(ctx, userID, dataCategory, map[string]interface{}{
		"name":  "Test User",
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	// Check retention policy
	policy, err := framework.GetRetentionPolicy(ctx, dataCategory)
	assert.NoError(t, err)
	assert.NotNil(t, policy)
	assert.Equal(t, 2*365*24*time.Hour, policy.RetentionPeriod)

	// Check if data should be deleted (shouldn't be for recent data)
	shouldDelete, err := framework.ShouldDeleteData(ctx, userID, dataCategory)
	assert.NoError(t, err)
	assert.False(t, shouldDelete)
}

func TestComplianceFramework_DataRights_Access(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()

	// Record some data
	err := framework.RecordDataCreation(ctx, userID, "personal", map[string]interface{}{
		"name":  "Test User",
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	// Process data access request
	request := DataAccessRequest{
		UserID:    userID,
		RequestID: uuid.New(),
		Purpose:   "data_portability",
		Format:    "json",
	}

	result, err := framework.ProcessDataAccessRequest(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Data)
	assert.Equal(t, "completed", result.Status)
}

func TestComplianceFramework_DataRights_Deletion(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()

	// Record some data
	err := framework.RecordDataCreation(ctx, userID, "personal", map[string]interface{}{
		"name":  "Test User",
		"email": "test@example.com",
	})
	assert.NoError(t, err)

	// Process deletion request
	request := DataDeletionRequest{
		UserID:    userID,
		RequestID: uuid.New(),
		Reason:    "user_request",
		Scope:     "all",
	}

	result, err := framework.ProcessDataDeletionRequest(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "completed", result.Status)
	assert.True(t, result.DeletedRecords > 0)
}

func TestComplianceFramework_Anonymization(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	originalData := map[string]interface{}{
		"name":  "João Silva",
		"email": "joao@example.com",
		"phone": "+5511999999999",
		"age":   30,
	}

	anonymizedData, err := framework.AnonymizeData(ctx, originalData, "pseudonymization")
	assert.NoError(t, err)
	assert.NotNil(t, anonymizedData)

	// Sensitive fields should be anonymized
	assert.NotEqual(t, originalData["email"], anonymizedData["email"])
	assert.NotEqual(t, originalData["phone"], anonymizedData["phone"])
	assert.NotEqual(t, originalData["name"], anonymizedData["name"])

	// Non-sensitive fields should remain the same
	assert.Equal(t, originalData["age"], anonymizedData["age"])
}

func TestComplianceFramework_AuditLogging(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()
	action := "data_access"
	details := map[string]interface{}{
		"requested_fields": []string{"name", "email"},
		"reason":          "compliance_request",
	}

	// Log audit event
	err := framework.LogAuditEvent(ctx, userID, action, details)
	assert.NoError(t, err)

	// Retrieve audit logs
	logs, err := framework.GetAuditLogs(ctx, AuditFilter{
		UserID: &userID,
		Action: &action,
		From:   time.Now().Add(-time.Hour),
		To:     time.Now().Add(time.Hour),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)
	assert.Equal(t, userID, logs[0].UserID)
	assert.Equal(t, action, logs[0].Action)
}

func TestComplianceFramework_GetComplianceStatus(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	status, err := framework.GetComplianceStatus(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, status)
	assert.True(t, status.Enabled)
	assert.Equal(t, "BR", status.Region)
	assert.NotEmpty(t, status.EnabledFeatures)
	assert.Contains(t, status.EnabledFeatures, "PII_DETECTION")
	assert.Contains(t, status.EnabledFeatures, "CONSENT_MANAGEMENT")
	assert.Contains(t, status.EnabledFeatures, "DATA_RETENTION")
	assert.Contains(t, status.EnabledFeatures, "AUDIT_LOGGING")
	assert.Contains(t, status.EnabledFeatures, "LGPD")
	assert.Contains(t, status.EnabledFeatures, "GDPR")
}

func TestComplianceFramework_ValidateCompliance(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	userID := uuid.New()

	// Record consent first
	err := framework.RecordConsent(ctx, userID, []string{"processing"}, "web")
	assert.NoError(t, err)

	// Validate compliance for data processing
	result, err := framework.ValidateCompliance(ctx, ComplianceValidationRequest{
		UserID:      userID,
		Action:      "process_data",
		DataType:    "personal",
		Purpose:     "processing",
		LegalBasis:  "consent",
		Jurisdiction: "BR",
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Compliant)
	assert.Empty(t, result.Violations)

	// Test validation without consent
	result, err = framework.ValidateCompliance(ctx, ComplianceValidationRequest{
		UserID:      uuid.New(), // Different user without consent
		Action:      "process_data",
		DataType:    "personal",
		Purpose:     "processing",
		LegalBasis:  "consent",
		Jurisdiction: "BR",
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Compliant)
	assert.NotEmpty(t, result.Violations)
}

func TestComplianceFramework_ConcurrentOperations(t *testing.T) {
	framework := createTestComplianceFramework(t)
	ctx := context.Background()

	numOperations := 50
	done := make(chan bool, numOperations)

	// Run concurrent consent operations
	for i := 0; i < numOperations; i++ {
		go func(i int) {
			userID := uuid.New()
			purposes := []string{"processing", "analytics"}

			err := framework.RecordConsent(ctx, userID, purposes, "web")
			assert.NoError(t, err)

			hasConsent, err := framework.HasConsent(ctx, userID, "processing")
			assert.NoError(t, err)
			assert.True(t, hasConsent)

			done <- true
		}(i)
	}

	// Wait for all operations to complete
	for i := 0; i < numOperations; i++ {
		<-done
	}
}

func TestComplianceFramework_ConfigValidation(t *testing.T) {
	logger := zaptest.NewLogger(t)

	// Test with invalid config (disabled PIIDetection but AutoMask enabled)
	invalidConfig := ComplianceConfig{
		Enabled: true,
		PIIDetection: PIIDetectionConfig{
			Enabled:  false,
			AutoMask: true, // This should cause validation to fail
		},
	}

	framework, err := NewComplianceFramework(invalidConfig, logger)
	// Should handle gracefully or return meaningful error
	if err != nil {
		assert.Contains(t, err.Error(), "invalid configuration")
	} else {
		assert.NotNil(t, framework)
		// Framework should adjust config to be valid
		assert.False(t, framework.config.PIIDetection.AutoMask)
	}
}