package compliance

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// ComplianceFramework provides comprehensive data protection compliance
type ComplianceFramework struct {
	config       ComplianceConfig
	logger       *zap.Logger
	piiManager   *PIIManager
	consentMgr   *ConsentManager
	auditLogger  *AuditLogger
	dataMapper   *DataMapper
	retentionMgr *RetentionManager
}

// ComplianceConfig holds all compliance-related configuration
type ComplianceConfig struct {
	Enabled       bool                `yaml:"enabled" envconfig:"COMPLIANCE_ENABLED" default:"true"`
	DefaultRegion string              `yaml:"default_region" envconfig:"DEFAULT_REGION" default:"BR"`
	PIIDetection  PIIDetectionConfig  `yaml:"pii_detection"`
	Consent       ConsentConfig       `yaml:"consent"`
	DataRetention DataRetentionConfig `yaml:"data_retention"`
	AuditLogging  AuditLoggingConfig  `yaml:"audit_logging"`
	LGPD          LGPDConfig          `yaml:"lgpd"`
	GDPR          GDPRConfig          `yaml:"gdpr"`
	Anonymization AnonymizationConfig `yaml:"anonymization"`
	DataRights    DataRightsConfig    `yaml:"data_rights"`
}

// PIIDetectionConfig configures PII detection and classification
type PIIDetectionConfig struct {
	Enabled           bool     `yaml:"enabled" default:"true"`
	ScanFields        []string `yaml:"scan_fields"`
	ClassificationAPI string   `yaml:"classification_api"`
	Confidence        float64  `yaml:"confidence" default:"0.8"`
	AutoMask          bool     `yaml:"auto_mask" default:"true"`
}

// ConsentConfig configures consent management
type ConsentConfig struct {
	Enabled         bool          `yaml:"enabled" default:"true"`
	DefaultPurposes []string      `yaml:"default_purposes"`
	TTL             time.Duration `yaml:"ttl" default:"2y"`
	GranularLevel   string        `yaml:"granular_level" default:"purpose"` // purpose, field, operation
}

// DataRetentionConfig configures data retention policies
type DataRetentionConfig struct {
	Enabled         bool                     `yaml:"enabled" default:"true"`
	DefaultPeriod   time.Duration            `yaml:"default_period" default:"2y"`
	CategoryPeriods map[string]time.Duration `yaml:"category_periods"`
	AutoDelete      bool                     `yaml:"auto_delete" default:"true"`
	BackupRetention time.Duration            `yaml:"backup_retention" default:"7y"`
}

// AuditLoggingConfig configures compliance audit logging
type AuditLoggingConfig struct {
	Enabled           bool          `yaml:"enabled" default:"true"`
	DetailLevel       string        `yaml:"detail_level" default:"full"` // minimal, standard, full
	RetentionPeriod   time.Duration `yaml:"retention_period" default:"7y"`
	EncryptionEnabled bool          `yaml:"encryption_enabled" default:"true"`
	ExternalLogging   bool          `yaml:"external_logging" default:"false"`
	ExternalEndpoint  string        `yaml:"external_endpoint"`
}

// LGPDConfig specific configuration for Brazilian LGPD compliance
type LGPDConfig struct {
	Enabled          bool     `yaml:"enabled" default:"true"`
	DPOContact       string   `yaml:"dpo_contact"`
	LegalBasis       []string `yaml:"legal_basis"`
	DataCategories   []string `yaml:"data_categories"`
	SharedThirdParty bool     `yaml:"shared_third_party" default:"false"`
}

// GDPRConfig specific configuration for European GDPR compliance
type GDPRConfig struct {
	Enabled             bool     `yaml:"enabled" default:"true"`
	DPOContact          string   `yaml:"dpo_contact"`
	LegalBasis          []string `yaml:"legal_basis"`
	DataCategories      []string `yaml:"data_categories"`
	CrossBorderTransfer bool     `yaml:"cross_border_transfer" default:"false"`
	AdequacyDecisions   []string `yaml:"adequacy_decisions"`
}

// AnonymizationConfig configures data anonymization
type AnonymizationConfig struct {
	Enabled    bool              `yaml:"enabled" default:"true"`
	Methods    []string          `yaml:"methods"` // hash, encrypt, tokenize, redact, generalize
	HashSalt   string            `yaml:"hash_salt"`
	Reversible bool              `yaml:"reversible" default:"false"`
	KAnonymity int               `yaml:"k_anonymity" default:"5"`
	Algorithms map[string]string `yaml:"algorithms"`
}

// DataRightsConfig configures individual data rights handling
type DataRightsConfig struct {
	Enabled              bool          `yaml:"enabled" default:"true"`
	ResponseTime         time.Duration `yaml:"response_time" default:"720h"` // 30 days
	AutoFulfillment      bool          `yaml:"auto_fulfillment" default:"false"`
	VerificationRequired bool          `yaml:"verification_required" default:"true"`
	NotificationChannels []string      `yaml:"notification_channels"`
}

// DataSubject represents an individual whose data is being processed
type DataSubject struct {
	ID          string                 `json:"id"`
	Email       string                 `json:"email"`
	Region      string                 `json:"region"`
	ConsentData map[string]ConsentInfo `json:"consent_data"`
	DataRights  []DataRightRequest     `json:"data_rights"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// ConsentInfo represents consent information for a specific purpose
type ConsentInfo struct {
	Purpose     string     `json:"purpose"`
	Granted     bool       `json:"granted"`
	Timestamp   time.Time  `json:"timestamp"`
	Source      string     `json:"source"`
	LegalBasis  string     `json:"legal_basis"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	WithdrawnAt *time.Time `json:"withdrawn_at,omitempty"`
}

// DataRightRequest represents a data subject's rights request
type DataRightRequest struct {
	ID               string                 `json:"id"`
	Type             DataRightType          `json:"type"`
	Status           DataRightStatus        `json:"status"`
	RequestedAt      time.Time              `json:"requested_at"`
	CompletedAt      *time.Time             `json:"completed_at,omitempty"`
	Data             map[string]interface{} `json:"data,omitempty"`
	Reason           string                 `json:"reason,omitempty"`
	VerificationCode string                 `json:"verification_code,omitempty"`
}

// DataRightType represents the type of data rights request
type DataRightType string

const (
	DataRightAccess          DataRightType = "access"           // Right to access (Art. 15 GDPR / Art. 18 LGPD)
	DataRightRectification   DataRightType = "rectification"    // Right to rectification (Art. 16 GDPR / Art. 18 LGPD)
	DataRightErasure         DataRightType = "erasure"          // Right to erasure (Art. 17 GDPR / Art. 18 LGPD)
	DataRightPortability     DataRightType = "portability"      // Right to data portability (Art. 20 GDPR / Art. 18 LGPD)
	DataRightRestriction     DataRightType = "restriction"      // Right to restriction (Art. 18 GDPR)
	DataRightObjection       DataRightType = "objection"        // Right to object (Art. 21 GDPR / Art. 18 LGPD)
	DataRightWithdrawConsent DataRightType = "withdraw_consent" // Right to withdraw consent
)

// DataRightStatus represents the status of a data rights request
type DataRightStatus string

const (
	DataRightStatusPending    DataRightStatus = "pending"
	DataRightStatusInProgress DataRightStatus = "in_progress"
	DataRightStatusCompleted  DataRightStatus = "completed"
	DataRightStatusRejected   DataRightStatus = "rejected"
	DataRightStatusPartial    DataRightStatus = "partial"
)

// NewComplianceFramework creates a new compliance framework instance
func NewComplianceFramework(config ComplianceConfig, logger *zap.Logger) (*ComplianceFramework, error) {
	if !config.Enabled {
		return &ComplianceFramework{
			config: config,
			logger: logger,
		}, nil
	}

	// Initialize PII Manager
	piiManager, err := NewPIIManager(config.PIIDetection, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize PII manager: %w", err)
	}

	// Initialize Consent Manager
	consentMgr, err := NewConsentManager(config.Consent, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consent manager: %w", err)
	}

	// Initialize Audit Logger
	auditLogger, err := NewAuditLogger(config.AuditLogging, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize audit logger: %w", err)
	}

	// Initialize Data Mapper
	dataMapper, err := NewDataMapper(config, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize data mapper: %w", err)
	}

	// Initialize Retention Manager
	retentionMgr, err := NewRetentionManager(config.DataRetention, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize retention manager: %w", err)
	}

	return &ComplianceFramework{
		config:       config,
		logger:       logger,
		piiManager:   piiManager,
		consentMgr:   consentMgr,
		auditLogger:  auditLogger,
		dataMapper:   dataMapper,
		retentionMgr: retentionMgr,
	}, nil
}

// ProcessData processes data through the compliance pipeline
func (cf *ComplianceFramework) ProcessData(ctx context.Context, subjectID string, data map[string]interface{}, purpose string) (map[string]interface{}, error) {
	if !cf.config.Enabled {
		return data, nil
	}

	// Audit the data processing attempt
	if err := cf.auditLogger.LogDataProcessing(ctx, subjectID, purpose, "attempt", data); err != nil {
		cf.logger.Warn("Failed to log data processing attempt", zap.Error(err))
	}

	// Check consent
	hasConsent, err := cf.consentMgr.HasValidConsent(ctx, subjectID, purpose)
	if err != nil {
		return nil, fmt.Errorf("failed to check consent: %w", err)
	}

	if !hasConsent {
		// Audit consent failure
		cf.auditLogger.LogDataProcessing(ctx, subjectID, purpose, "consent_denied", nil)
		return nil, fmt.Errorf("no valid consent for purpose: %s", purpose)
	}

	// Detect and classify PII
	processedData, err := cf.piiManager.ProcessData(ctx, data)
	if err != nil {
		cf.auditLogger.LogDataProcessing(ctx, subjectID, purpose, "pii_error", nil)
		return nil, fmt.Errorf("PII processing failed: %w", err)
	}

	// Apply retention policy
	if err := cf.retentionMgr.ApplyRetentionPolicy(ctx, subjectID, processedData); err != nil {
		cf.logger.Warn("Failed to apply retention policy", zap.Error(err))
	}

	// Audit successful processing
	if err := cf.auditLogger.LogDataProcessing(ctx, subjectID, purpose, "success", processedData); err != nil {
		cf.logger.Warn("Failed to log successful data processing", zap.Error(err))
	}

	return processedData, nil
}

// HandleDataRightRequest processes a data subject rights request
func (cf *ComplianceFramework) HandleDataRightRequest(ctx context.Context, subjectID string, request DataRightRequest) error {
	if !cf.config.Enabled {
		return fmt.Errorf("compliance framework is disabled")
	}

	// Audit the rights request
	if err := cf.auditLogger.LogDataRightsRequest(ctx, subjectID, request); err != nil {
		cf.logger.Warn("Failed to audit data rights request", zap.Error(err))
	}

	switch request.Type {
	case DataRightAccess:
		return cf.handleAccessRequest(ctx, subjectID, request)
	case DataRightErasure:
		return cf.handleErasureRequest(ctx, subjectID, request)
	case DataRightRectification:
		return cf.handleRectificationRequest(ctx, subjectID, request)
	case DataRightPortability:
		return cf.handlePortabilityRequest(ctx, subjectID, request)
	case DataRightWithdrawConsent:
		return cf.handleConsentWithdrawal(ctx, subjectID, request)
	default:
		return fmt.Errorf("unsupported data right type: %s", request.Type)
	}
}

// GetComplianceStatus returns the current compliance status
func (cf *ComplianceFramework) GetComplianceStatus(ctx context.Context) (map[string]interface{}, error) {
	if !cf.config.Enabled {
		return map[string]interface{}{
			"enabled": false,
			"status":  "disabled",
		}, nil
	}

	status := map[string]interface{}{
		"enabled":        true,
		"default_region": cf.config.DefaultRegion,
		"lgpd_enabled":   cf.config.LGPD.Enabled,
		"gdpr_enabled":   cf.config.GDPR.Enabled,
		"components": map[string]interface{}{
			"pii_detection":  cf.config.PIIDetection.Enabled,
			"consent_mgmt":   cf.config.Consent.Enabled,
			"audit_logging":  cf.config.AuditLogging.Enabled,
			"data_retention": cf.config.DataRetention.Enabled,
			"anonymization":  cf.config.Anonymization.Enabled,
		},
	}

	// Add component health checks
	if cf.piiManager != nil {
		status["pii_manager"] = cf.piiManager.HealthCheck(ctx)
	}
	if cf.consentMgr != nil {
		status["consent_manager"] = cf.consentMgr.HealthCheck(ctx)
	}

	return status, nil
}

// Helper methods for handling specific data rights requests
func (cf *ComplianceFramework) handleAccessRequest(ctx context.Context, subjectID string, request DataRightRequest) error {
	// Implementation for access request
	cf.logger.Info("Processing access request", zap.String("subject_id", subjectID), zap.String("request_id", request.ID))
	// TODO: Implement data extraction and anonymization
	return nil
}

func (cf *ComplianceFramework) handleErasureRequest(ctx context.Context, subjectID string, request DataRightRequest) error {
	// Implementation for erasure request (right to be forgotten)
	cf.logger.Info("Processing erasure request", zap.String("subject_id", subjectID), zap.String("request_id", request.ID))
	// TODO: Implement data deletion across all systems
	return nil
}

func (cf *ComplianceFramework) handleRectificationRequest(ctx context.Context, subjectID string, request DataRightRequest) error {
	// Implementation for rectification request
	cf.logger.Info("Processing rectification request", zap.String("subject_id", subjectID), zap.String("request_id", request.ID))
	// TODO: Implement data correction
	return nil
}

func (cf *ComplianceFramework) handlePortabilityRequest(ctx context.Context, subjectID string, request DataRightRequest) error {
	// Implementation for portability request
	cf.logger.Info("Processing portability request", zap.String("subject_id", subjectID), zap.String("request_id", request.ID))
	// TODO: Implement data export in portable format
	return nil
}

func (cf *ComplianceFramework) handleConsentWithdrawal(ctx context.Context, subjectID string, request DataRightRequest) error {
	// Implementation for consent withdrawal
	cf.logger.Info("Processing consent withdrawal", zap.String("subject_id", subjectID), zap.String("request_id", request.ID))
	return cf.consentMgr.WithdrawConsent(ctx, subjectID, request.Data["purpose"].(string))
}

// GetConsentManager returns the consent manager for direct access
func (cf *ComplianceFramework) GetConsentManager() *ConsentManager {
	return cf.consentMgr
}

// GetPIIManager returns the PII manager for direct access
func (cf *ComplianceFramework) GetPIIManager() *PIIManager {
	return cf.piiManager
}

// GetAuditLogger returns the audit logger for direct access
func (cf *ComplianceFramework) GetAuditLogger() *AuditLogger {
	return cf.auditLogger
}

// GetDataMapper returns the data mapper for direct access
func (cf *ComplianceFramework) GetDataMapper() *DataMapper {
	return cf.dataMapper
}

// GetRetentionManager returns the retention manager for direct access
func (cf *ComplianceFramework) GetRetentionManager() *RetentionManager {
	return cf.retentionMgr
}
