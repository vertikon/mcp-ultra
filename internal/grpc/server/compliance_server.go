package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"go.uber.org/zap"

	complivancev1 "github.com/vertikon/mcp-ultra/api/grpc/gen/compliance/v1"
	"github.com/vertikon/mcp-ultra/internal/compliance"
)

// ComplianceServer implements the ComplianceService gRPC server
type ComplianceServer struct {
	complivancev1.UnimplementedComplianceServiceServer
	framework *compliance.ComplianceFramework
	logger    *zap.Logger
}

// NewComplianceServer creates a new ComplianceServer instance
func NewComplianceServer(framework *compliance.ComplianceFramework, logger *zap.Logger) *ComplianceServer {
	return &ComplianceServer{
		framework: framework,
		logger:    logger,
	}
}

// RecordConsent records user consent for data processing
func (s *ComplianceServer) RecordConsent(ctx context.Context, req *complivancev1.RecordConsentRequest) (*complivancev1.RecordConsentResponse, error) {
	s.logger.Info("Recording consent", zap.String("subject_id", req.SubjectId))

	consentRecord := &compliance.ConsentRecord{
		SubjectID:     req.SubjectId,
		Purposes:      req.Purposes,
		LegalBasis:    req.LegalBasis,
		ConsentGiven:  req.ConsentGiven,
		Timestamp:     time.Now(),
		ExpiresAt:     req.ExpiresAt.AsTime(),
		Metadata:      req.Metadata,
		IPAddress:     req.IpAddress,
		UserAgent:     req.UserAgent,
		ConsentMethod: req.ConsentMethod,
		Version:       req.Version,
	}

	err := s.framework.RecordConsent(ctx, consentRecord)
	if err != nil {
		s.logger.Error("Failed to record consent", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to record consent: %v", err)
	}

	return &complivancev1.RecordConsentResponse{
		ConsentId: consentRecord.ID,
		Success:   true,
	}, nil
}

// GetConsent retrieves consent information for a subject
func (s *ComplianceServer) GetConsent(ctx context.Context, req *complivancev1.GetConsentRequest) (*complivancev1.GetConsentResponse, error) {
	s.logger.Info("Getting consent", zap.String("subject_id", req.SubjectId))

	consents, err := s.framework.GetConsent(ctx, req.SubjectId, req.Purposes)
	if err != nil {
		s.logger.Error("Failed to get consent", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get consent: %v", err)
	}

	pbConsents := make([]*complivancev1.ConsentRecord, 0, len(consents))
	for _, consent := range consents {
		pbConsent := &complivancev1.ConsentRecord{
			ConsentId:     consent.ID,
			SubjectId:     consent.SubjectID,
			Purposes:      consent.Purposes,
			LegalBasis:    consent.LegalBasis,
			ConsentGiven:  consent.ConsentGiven,
			Timestamp:     timestamppb.New(consent.Timestamp),
			ExpiresAt:     timestamppb.New(consent.ExpiresAt),
			Metadata:      consent.Metadata,
			IpAddress:     consent.IPAddress,
			UserAgent:     consent.UserAgent,
			ConsentMethod: consent.ConsentMethod,
			Version:       consent.Version,
		}
		pbConsents = append(pbConsents, pbConsent)
	}

	return &complivancev1.GetConsentResponse{
		Consents: pbConsents,
	}, nil
}

// WithdrawConsent withdraws previously given consent
func (s *ComplianceServer) WithdrawConsent(ctx context.Context, req *complivancev1.WithdrawConsentRequest) (*complivancev1.WithdrawConsentResponse, error) {
	s.logger.Info("Withdrawing consent", zap.String("subject_id", req.SubjectId))

	err := s.framework.WithdrawConsent(ctx, req.SubjectId, req.Purposes, req.Reason)
	if err != nil {
		s.logger.Error("Failed to withdraw consent", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to withdraw consent: %v", err)
	}

	return &complivancev1.WithdrawConsentResponse{
		Success: true,
	}, nil
}

// ValidateConsent validates current consent status
func (s *ComplianceServer) ValidateConsent(ctx context.Context, req *complivancev1.ValidateConsentRequest) (*complivancev1.ValidateConsentResponse, error) {
	s.logger.Info("Validating consent", zap.String("subject_id", req.SubjectId))

	isValid, err := s.framework.ValidateConsent(ctx, req.SubjectId, req.Purpose)
	if err != nil {
		s.logger.Error("Failed to validate consent", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to validate consent: %v", err)
	}

	var reason string
	if !isValid {
		reason = "consent not found, expired, or withdrawn"
	}

	return &complivancev1.ValidateConsentResponse{
		IsValid: isValid,
		Reason:  reason,
	}, nil
}

// ProcessDataSubjectRequest handles data subject rights requests (GDPR/LGPD)
func (s *ComplianceServer) ProcessDataSubjectRequest(ctx context.Context, req *complivancev1.ProcessDataSubjectRequestRequest) (*complivancev1.ProcessDataSubjectRequestResponse, error) {
	s.logger.Info("Processing data subject request", 
		zap.String("subject_id", req.SubjectId),
		zap.String("request_type", req.RequestType.String()))

	var requestType compliance.DataSubjectRequestType
	switch req.RequestType {
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_ACCESS:
		requestType = compliance.DataSubjectRequestAccess
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_ERASURE:
		requestType = compliance.DataSubjectRequestErasure
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_RECTIFICATION:
		requestType = compliance.DataSubjectRequestRectification
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_PORTABILITY:
		requestType = compliance.DataSubjectRequestPortability
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_OBJECTION:
		requestType = compliance.DataSubjectRequestObjection
	case complivancev1.DataSubjectRequestType_DATA_SUBJECT_REQUEST_TYPE_RESTRICTION:
		requestType = compliance.DataSubjectRequestRestriction
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid request type: %v", req.RequestType)
	}

	result, err := s.framework.ProcessDataSubjectRequest(ctx, &compliance.DataSubjectRequest{
		SubjectID:   req.SubjectId,
		Type:        requestType,
		Description: req.Description,
		Email:       req.Email,
		Metadata:    req.Metadata,
	})
	if err != nil {
		s.logger.Error("Failed to process data subject request", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to process data subject request: %v", err)
	}

	return &complivancev1.ProcessDataSubjectRequestResponse{
		RequestId:   result.RequestID,
		Status:      result.Status,
		Message:     result.Message,
		EstimatedAt: timestamppb.New(result.EstimatedCompletionTime),
	}, nil
}

// DetectPII detects personally identifiable information in text
func (s *ComplianceServer) DetectPII(ctx context.Context, req *complivancev1.DetectPIIRequest) (*complivancev1.DetectPIIResponse, error) {
	s.logger.Debug("Detecting PII", zap.Int("text_length", len(req.Text)))

	result, err := s.framework.DetectPII(ctx, req.Text)
	if err != nil {
		s.logger.Error("Failed to detect PII", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to detect PII: %v", err)
	}

	pbEntities := make([]*complivancev1.PIIEntity, 0, len(result.Entities))
	for _, entity := range result.Entities {
		pbEntity := &complivancev1.PIIEntity{
			Type:       entity.Type,
			Value:      entity.Value,
			StartPos:   int32(entity.StartPos),
			EndPos:     int32(entity.EndPos),
			Confidence: float32(entity.Confidence),
			Category:   entity.Category,
		}
		pbEntities = append(pbEntities, pbEntity)
	}

	return &complivancev1.DetectPIIResponse{
		HasPii:    result.HasPII,
		Entities:  pbEntities,
		Confidence: float32(result.Confidence),
		Metadata:  result.Metadata,
	}, nil
}

// AnonymizeData anonymizes sensitive data
func (s *ComplianceServer) AnonymizeData(ctx context.Context, req *complivancev1.AnonymizeDataRequest) (*complivancev1.AnonymizeDataResponse, error) {
	s.logger.Debug("Anonymizing data", zap.String("method", req.Method))

	result, err := s.framework.AnonymizeData(ctx, req.Data, req.Method, req.Options)
	if err != nil {
		s.logger.Error("Failed to anonymize data", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to anonymize data: %v", err)
	}

	return &complivancev1.AnonymizeDataResponse{
		AnonymizedData: result.Data,
		Method:         result.Method,
		Metadata:       result.Metadata,
		Reversible:     result.Reversible,
		ReversalKey:    result.ReversalKey,
	}, nil
}

// GetAuditLog retrieves audit log entries
func (s *ComplianceServer) GetAuditLog(ctx context.Context, req *complivancev1.GetAuditLogRequest) (*complivancev1.GetAuditLogResponse, error) {
	s.logger.Info("Getting audit log", zap.String("subject_id", req.SubjectId))

	entries, err := s.framework.GetAuditLog(ctx, &compliance.AuditLogQuery{
		SubjectID:  req.SubjectId,
		Action:     req.Action,
		StartTime:  req.StartTime.AsTime(),
		EndTime:    req.EndTime.AsTime(),
		Limit:      int(req.Limit),
		Offset:     int(req.Offset),
	})
	if err != nil {
		s.logger.Error("Failed to get audit log", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get audit log: %v", err)
	}

	pbEntries := make([]*complivancev1.AuditLogEntry, 0, len(entries))
	for _, entry := range entries {
		pbEntry := &complivancev1.AuditLogEntry{
			Id:        entry.ID,
			SubjectId: entry.SubjectID,
			Action:    entry.Action,
			Resource:  entry.Resource,
			Timestamp: timestamppb.New(entry.Timestamp),
			UserId:    entry.UserID,
			IpAddress: entry.IPAddress,
			UserAgent: entry.UserAgent,
			Metadata:  entry.Metadata,
			Result:    entry.Result,
			Details:   entry.Details,
		}
		pbEntries = append(pbEntries, pbEntry)
	}

	return &complivancev1.GetAuditLogResponse{
		Entries: pbEntries,
		Total:   int32(len(pbEntries)),
	}, nil
}

// CreateRetentionPolicy creates a new data retention policy
func (s *ComplianceServer) CreateRetentionPolicy(ctx context.Context, req *complivancev1.CreateRetentionPolicyRequest) (*complivancev1.CreateRetentionPolicyResponse, error) {
	s.logger.Info("Creating retention policy", zap.String("name", req.Policy.Name))

	policy := &compliance.RetentionPolicy{
		Name:        req.Policy.Name,
		Description: req.Policy.Description,
		DataTypes:   req.Policy.DataTypes,
		Duration:    req.Policy.Duration.AsDuration(),
		AutoDelete:  req.Policy.AutoDelete,
		Conditions:  req.Policy.Conditions,
		Metadata:    req.Policy.Metadata,
		LegalBasis:  req.Policy.LegalBasis,
		Region:      req.Policy.Region,
	}

	err := s.framework.CreateRetentionPolicy(ctx, policy)
	if err != nil {
		s.logger.Error("Failed to create retention policy", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create retention policy: %v", err)
	}

	return &complivancev1.CreateRetentionPolicyResponse{
		PolicyId: policy.ID,
		Success:  true,
	}, nil
}

// GetRetentionPolicies retrieves retention policies
func (s *ComplianceServer) GetRetentionPolicies(ctx context.Context, req *complivancev1.GetRetentionPoliciesRequest) (*complivancev1.GetRetentionPoliciesResponse, error) {
	s.logger.Info("Getting retention policies", zap.String("data_type", req.DataType))

	policies, err := s.framework.GetRetentionPolicies(ctx, req.DataType)
	if err != nil {
		s.logger.Error("Failed to get retention policies", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get retention policies: %v", err)
	}

	pbPolicies := make([]*complivancev1.RetentionPolicy, 0, len(policies))
	for _, policy := range policies {
		pbPolicy := &complivancev1.RetentionPolicy{
			Id:          policy.ID,
			Name:        policy.Name,
			Description: policy.Description,
			DataTypes:   policy.DataTypes,
			Duration:    nil, // TODO: Convert duration
			AutoDelete:  policy.AutoDelete,
			Conditions:  policy.Conditions,
			Metadata:    policy.Metadata,
			LegalBasis:  policy.LegalBasis,
			Region:      policy.Region,
			CreatedAt:   timestamppb.New(policy.CreatedAt),
			UpdatedAt:   timestamppb.New(policy.UpdatedAt),
		}
		pbPolicies = append(pbPolicies, pbPolicy)
	}

	return &complivancev1.GetRetentionPoliciesResponse{
		Policies: pbPolicies,
	}, nil
}

// ApplyRetentionPolicy applies retention policy to data
func (s *ComplianceServer) ApplyRetentionPolicy(ctx context.Context, req *complivancev1.ApplyRetentionPolicyRequest) (*complivancev1.ApplyRetentionPolicyResponse, error) {
	s.logger.Info("Applying retention policy", zap.String("policy_id", req.PolicyId))

	result, err := s.framework.ApplyRetentionPolicy(ctx, req.PolicyId, req.DataId)
	if err != nil {
		s.logger.Error("Failed to apply retention policy", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to apply retention policy: %v", err)
	}

	return &complivancev1.ApplyRetentionPolicyResponse{
		Success:     result.Success,
		ExpiresAt:   timestamppb.New(result.ExpiresAt),
		ActionTaken: result.ActionTaken,
		Message:     result.Message,
	}, nil
}

// GetComplianceStatus gets overall compliance status
func (s *ComplianceServer) GetComplianceStatus(ctx context.Context, req *complivancev1.GetComplianceStatusRequest) (*complivancev1.GetComplianceStatusResponse, error) {
	s.logger.Info("Getting compliance status")

	status, err := s.framework.GetComplianceStatus(ctx)
	if err != nil {
		s.logger.Error("Failed to get compliance status", zap.Error(err))
		return nil, fmt.Errorf("failed to get compliance status: %v", err)
	}

	pbChecks := make([]*complivancev1.ComplianceCheck, 0, len(status.Checks))
	for _, check := range status.Checks {
		pbCheck := &complivancev1.ComplianceCheck{
			Name:        check.Name,
			Status:      check.Status,
			Description: check.Description,
			LastChecked: timestamppb.New(check.LastChecked),
			Metadata:    check.Metadata,
		}
		pbChecks = append(pbChecks, pbCheck)
	}

	return &complivancev1.GetComplianceStatusResponse{
		OverallStatus: status.OverallStatus,
		Checks:        pbChecks,
		LastUpdated:   timestamppb.New(status.LastUpdated),
		Metadata:      status.Metadata,
	}, nil
}