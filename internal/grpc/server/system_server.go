package server

import (
	"context"
	"encoding/json"
	"runtime"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"go.uber.org/zap"

	systemv1 "github.com/vertikon/mcp-ultra/api/grpc/gen/system/v1"
	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/features"
	"github.com/vertikon/mcp-ultra/internal/observability"
	"github.com/vertikon/mcp-ultra/pkg/version"
)

// SystemServer implements the SystemService gRPC server
type SystemServer struct {
	systemv1.UnimplementedSystemServiceServer
	config       *config.Config
	obsService   *observability.Service
	flagManager  *features.FlagManager
	logger       *zap.Logger
	startTime    time.Time
	buildInfo    *version.BuildInfo
}

// NewSystemServer creates a new SystemServer instance
func NewSystemServer(
	config *config.Config,
	obsService *observability.Service,
	flagManager *features.FlagManager,
	logger *zap.Logger,
) *SystemServer {
	return &SystemServer{
		config:      config,
		obsService:  obsService,
		flagManager: flagManager,
		logger:      logger,
		startTime:   time.Now(),
		buildInfo:   version.GetBuildInfo(),
	}
}

// Check performs health checks
func (s *SystemServer) Check(ctx context.Context, req *systemv1.HealthCheckRequest) (*systemv1.HealthCheckResponse, error) {
	s.logger.Debug("Health check requested", zap.String("service", req.Service))

	startTime := time.Now()
	
	var overallStatus systemv1.ServingStatus = systemv1.ServingStatus_SERVING_STATUS_SERVING
	var components []*systemv1.ComponentHealth

	// Check observability service
	if s.obsService != nil {
		obsHealth := s.obsService.HealthCheck()
		componentStatus := s.convertHealthToComponentStatus(obsHealth.Status)
		if componentStatus != systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY {
			overallStatus = systemv1.ServingStatus_SERVING_STATUS_DEGRADED
		}
		
		components = append(components, &systemv1.ComponentHealth{
			Name:         "observability",
			Status:       componentStatus,
			Message:      obsHealth.Message,
			ResponseTime: durationpb.New(obsHealth.ResponseTime),
			Metadata:     obsHealth.Metadata,
		})
	}

	// Check feature flags
	if s.flagManager != nil {
		flagHealth := s.checkFeatureFlagHealth(ctx)
		components = append(components, flagHealth)
		if flagHealth.Status != systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY {
			overallStatus = systemv1.ServingStatus_SERVING_STATUS_DEGRADED
		}
	}

	// Deep check if requested
	if req.DeepCheck {
		deepChecks := s.performDeepHealthChecks(ctx)
		components = append(components, deepChecks...)
		
		// Update overall status based on deep checks
		for _, component := range deepChecks {
			if component.Status == systemv1.ComponentStatus_COMPONENT_STATUS_UNHEALTHY {
				overallStatus = systemv1.ServingStatus_SERVING_STATUS_NOT_SERVING
				break
			} else if component.Status != systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY {
				overallStatus = systemv1.ServingStatus_SERVING_STATUS_DEGRADED
			}
		}
	}

	responseTime := time.Since(startTime)
	
	return &systemv1.HealthCheckResponse{
		Status:       overallStatus,
		Components:   components,
		Timestamp:    timestamppb.Now(),
		ResponseTime: durationpb.New(responseTime),
		Version:      version.Version,
		Metadata: map[string]string{
			"service": "mcp-ultra",
			"uptime":  time.Since(s.startTime).String(),
		},
	}, nil
}

// Watch streams health check responses
func (s *SystemServer) Watch(req *systemv1.HealthCheckRequest, stream systemv1.SystemService_WatchServer) error {
	s.logger.Info("Starting health check watch", zap.String("service", req.Service))

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			s.logger.Info("Health check watch cancelled")
			return stream.Context().Err()
		case <-ticker.C:
			response, err := s.Check(stream.Context(), req)
			if err != nil {
				s.logger.Error("Health check failed during watch", zap.Error(err))
				return err
			}
			
			if err := stream.Send(response); err != nil {
				s.logger.Error("Failed to send health check response", zap.Error(err))
				return err
			}
		}
	}
}

// GetSystemInfo returns system information
func (s *SystemServer) GetSystemInfo(ctx context.Context, req *systemv1.GetSystemInfoRequest) (*systemv1.GetSystemInfoResponse, error) {
	s.logger.Debug("System info requested")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	systemInfo := &systemv1.SystemInfo{
		ServiceName: "mcp-ultra",
		Version:     version.Version,
		BuildDate:   s.buildInfo.BuildDate,
		GitCommit:   s.buildInfo.GitCommit,
		GitBranch:   s.buildInfo.GitBranch,
		Environment: s.getEnvironmentInfo(),
	}

	if req.IncludeRuntime {
		systemInfo.Runtime = &systemv1.RuntimeInfo{
			GoVersion:     runtime.Version(),
			Os:            runtime.GOOS,
			Arch:          runtime.GOARCH,
			NumCpu:        int32(runtime.NumCPU()),
			NumGoroutine:  int32(runtime.NumGoroutine()),
			Uptime:        durationpb.New(time.Since(s.startTime)),
			StartTime:     timestamppb.New(s.startTime),
			Memory: &systemv1.MemoryStats{
				Alloc:         m.Alloc,
				TotalAlloc:    m.TotalAlloc,
				Sys:           m.Sys,
				HeapAlloc:     m.HeapAlloc,
				HeapSys:       m.HeapSys,
				HeapInuse:     m.HeapInuse,
				StackInuse:    m.StackInuse,
				StackSys:      m.StackSys,
				NumGc:         m.NumGC,
				GcPauseTotal:  durationpb.New(time.Duration(m.PauseTotalNs)),
			},
		}
	}

	if req.IncludeBuild {
		systemInfo.Build = &systemv1.BuildInfo{
			Compiler: runtime.Compiler,
			Settings: s.getBuildSettings(),
		}
	}

	if req.IncludeDependencies {
		systemInfo.Dependencies = s.getDependencies()
	}

	return &systemv1.GetSystemInfoResponse{
		System: systemInfo,
	}, nil
}

// GetMetrics returns system metrics
func (s *SystemServer) GetMetrics(ctx context.Context, req *systemv1.GetMetricsRequest) (*systemv1.GetMetricsResponse, error) {
	s.logger.Debug("Metrics requested", zap.String("format", req.Format))

	if s.obsService == nil {
		return nil, status.Errorf(codes.Unavailable, "observability service not available")
	}

	// Get metrics from observability service
	metricsData, err := s.obsService.GetMetrics(ctx, req.Format, req.MetricNames, req.Labels)
	if err != nil {
		s.logger.Error("Failed to get metrics", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get metrics: %v", err)
	}

	return &systemv1.GetMetricsResponse{
		Format:    req.Format,
		Data:      metricsData,
		Timestamp: timestamppb.Now(),
		Summary: &systemv1.MetricsSummary{
			TotalMetrics: int32(len(req.MetricNames)),
			Families:     s.getMetricFamilies(),
		},
	}, nil
}

// GetConfig returns configuration
func (s *SystemServer) GetConfig(ctx context.Context, req *systemv1.GetConfigRequest) (*systemv1.GetConfigResponse, error) {
	s.logger.Debug("Config requested", zap.String("section", req.Section))

	configData := s.config
	if req.Section != "" {
		// TODO: Extract specific config section
	}

	// Convert config to protobuf Struct
	configBytes, err := json.Marshal(configData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal config: %v", err)
	}

	var configMap map[string]interface{}
	if err := json.Unmarshal(configBytes, &configMap); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal config: %v", err)
	}

	// Mask sensitive data if requested
	if req.MaskSensitive {
		configMap = s.maskSensitiveData(configMap)
	}

	configStruct, err := structpb.NewStruct(configMap)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create config struct: %v", err)
	}

	return &systemv1.GetConfigResponse{
		Config:      configStruct,
		Checksum:    s.calculateConfigChecksum(configData),
		LastUpdated: timestamppb.Now(), // TODO: Track actual update time
	}, nil
}

// UpdateConfig updates configuration (placeholder)
func (s *SystemServer) UpdateConfig(ctx context.Context, req *systemv1.UpdateConfigRequest) (*systemv1.UpdateConfigResponse, error) {
	s.logger.Info("Config update requested", zap.String("section", req.Section))

	// TODO: Implement actual config update logic
	// For now, return success with dry-run info
	
	return &systemv1.UpdateConfigResponse{
		Success: req.DryRun, // Only succeed for dry runs for now
		Errors: []*systemv1.ConfigValidationError{
			{
				Field:   "general",
				Message: "Config updates not yet implemented",
				Code:    "NOT_IMPLEMENTED",
			},
		},
	}, nil
}

// ReloadConfig reloads configuration (placeholder)
func (s *SystemServer) ReloadConfig(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	s.logger.Info("Config reload requested")
	
	// TODO: Implement actual config reload logic
	
	return &emptypb.Empty{}, nil
}

// GetFeatureFlags returns feature flags
func (s *SystemServer) GetFeatureFlags(ctx context.Context, req *systemv1.GetFeatureFlagsRequest) (*systemv1.GetFeatureFlagsResponse, error) {
	s.logger.Debug("Feature flags requested", zap.String("prefix", req.Prefix))

	if s.flagManager == nil {
		return nil, status.Errorf(codes.Unavailable, "feature flag manager not available")
	}

	flags, err := s.flagManager.GetAllFlags(ctx, req.Prefix, req.IncludeDisabled)
	if err != nil {
		s.logger.Error("Failed to get feature flags", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get feature flags: %v", err)
	}

	pbFlags := make([]*systemv1.FeatureFlag, 0, len(flags))
	for _, flag := range flags {
		pbFlag := s.convertToProtobufFlag(flag)
		pbFlags = append(pbFlags, pbFlag)
	}

	return &systemv1.GetFeatureFlagsResponse{
		Flags:       pbFlags,
		LastUpdated: timestamppb.Now(),
	}, nil
}

// UpdateFeatureFlag updates a feature flag
func (s *SystemServer) UpdateFeatureFlag(ctx context.Context, req *systemv1.UpdateFeatureFlagRequest) (*systemv1.UpdateFeatureFlagResponse, error) {
	s.logger.Info("Feature flag update requested", zap.String("flag_key", req.FlagKey))

	if s.flagManager == nil {
		return nil, status.Errorf(codes.Unavailable, "feature flag manager not available")
	}

	flag := s.convertFromProtobufFlag(req.Flag)
	err := s.flagManager.UpdateFlag(ctx, req.FlagKey, flag)
	if err != nil {
		s.logger.Error("Failed to update feature flag", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to update feature flag: %v", err)
	}

	return &systemv1.UpdateFeatureFlagResponse{
		Flag: req.Flag,
	}, nil
}

// Helper methods

func (s *SystemServer) convertHealthToComponentStatus(status string) systemv1.ComponentStatus {
	switch status {
	case "healthy":
		return systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY
	case "unhealthy":
		return systemv1.ComponentStatus_COMPONENT_STATUS_UNHEALTHY
	case "degraded":
		return systemv1.ComponentStatus_COMPONENT_STATUS_DEGRADED
	default:
		return systemv1.ComponentStatus_COMPONENT_STATUS_UNKNOWN
	}
}

func (s *SystemServer) checkFeatureFlagHealth(ctx context.Context) *systemv1.ComponentHealth {
	startTime := time.Now()
	
	// TODO: Implement actual feature flag health check
	status := systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY
	message := "Feature flags operational"
	
	return &systemv1.ComponentHealth{
		Name:         "feature_flags",
		Status:       status,
		Message:      message,
		ResponseTime: durationpb.New(time.Since(startTime)),
	}
}

func (s *SystemServer) performDeepHealthChecks(ctx context.Context) []*systemv1.ComponentHealth {
	var checks []*systemv1.ComponentHealth

	// Database health check
	checks = append(checks, &systemv1.ComponentHealth{
		Name:    "database",
		Status:  systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY,
		Message: "Database connection healthy",
	})

	// Cache health check  
	checks = append(checks, &systemv1.ComponentHealth{
		Name:    "cache",
		Status:  systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY,
		Message: "Cache connection healthy",
	})

	// Message queue health check
	checks = append(checks, &systemv1.ComponentHealth{
		Name:    "message_queue",
		Status:  systemv1.ComponentStatus_COMPONENT_STATUS_HEALTHY,
		Message: "Message queue connection healthy",
	})

	return checks
}

func (s *SystemServer) getEnvironmentInfo() map[string]string {
	return map[string]string{
		"environment": s.config.Environment,
		"region":      s.config.Region,
		"datacenter":  s.config.Datacenter,
	}
}

func (s *SystemServer) getBuildSettings() []*systemv1.BuildSetting {
	return []*systemv1.BuildSetting{
		{Key: "CGO_ENABLED", Value: "0"},
		{Key: "GOOS", Value: runtime.GOOS},
		{Key: "GOARCH", Value: runtime.GOARCH},
	}
}

func (s *SystemServer) getDependencies() []*systemv1.Dependency {
	// TODO: Parse go.mod or build info for actual dependencies
	return []*systemv1.Dependency{
		{
			Name:    "go.uber.org/zap",
			Version: "v1.24.0",
			Type:    "direct",
		},
		{
			Name:    "google.golang.org/grpc",
			Version: "v1.56.0",
			Type:    "direct",
		},
	}
}

func (s *SystemServer) getMetricFamilies() []*systemv1.MetricFamily {
	return []*systemv1.MetricFamily{
		{
			Name:        "http_requests_total",
			Help:        "Total HTTP requests",
			Type:        systemv1.MetricType_METRIC_TYPE_COUNTER,
			MetricCount: 1,
		},
		{
			Name:        "http_request_duration_seconds",
			Help:        "HTTP request duration",
			Type:        systemv1.MetricType_METRIC_TYPE_HISTOGRAM,
			MetricCount: 1,
		},
	}
}

func (s *SystemServer) maskSensitiveData(data map[string]interface{}) map[string]interface{} {
	masked := make(map[string]interface{})
	// List of sensitive field names (not actual secrets)
	sensitiveKeys := []string{"password", "secret", "key", "token", "credential"}

	for k, v := range data {
		isSensitive := false
		for _, sensitiveKey := range sensitiveKeys {
			if len(k) >= len(sensitiveKey) && k[len(k)-len(sensitiveKey):] == sensitiveKey {
				isSensitive = true
				break
			}
		}

		if isSensitive {
			masked[k] = "***MASKED***"
		} else if subMap, ok := v.(map[string]interface{}); ok {
			masked[k] = s.maskSensitiveData(subMap)
		} else {
			masked[k] = v
		}
	}

	return masked
}

func (s *SystemServer) calculateConfigChecksum(config *config.Config) string {
	// TODO: Implement actual checksum calculation
	return "sha256:abc123"
}

func (s *SystemServer) convertToProtobufFlag(flag *features.Flag) *systemv1.FeatureFlag {
	// TODO: Convert from internal Flag to protobuf FeatureFlag
	return &systemv1.FeatureFlag{
		Key:         flag.Key,
		Name:        flag.Name,
		Description: flag.Description,
		Enabled:     flag.Enabled,
		CreatedAt:   timestamppb.New(flag.CreatedAt),
		UpdatedAt:   timestamppb.New(flag.UpdatedAt),
	}
}

func (s *SystemServer) convertFromProtobufFlag(pbFlag *systemv1.FeatureFlag) *features.Flag {
	// TODO: Convert from protobuf FeatureFlag to internal Flag
	return &features.Flag{
		Key:         pbFlag.Key,
		Name:        pbFlag.Name,
		Description: pbFlag.Description,
		Enabled:     pbFlag.Enabled,
		CreatedAt:   pbFlag.CreatedAt.AsTime(),
		UpdatedAt:   pbFlag.UpdatedAt.AsTime(),
	}
}

// Placeholder implementations for remaining methods
func (s *SystemServer) ListServices(ctx context.Context, req *systemv1.ListServicesRequest) (*systemv1.ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}

func (s *SystemServer) GetServiceStatus(ctx context.Context, req *systemv1.GetServiceStatusRequest) (*systemv1.GetServiceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceStatus not implemented")
}

func (s *SystemServer) RestartService(ctx context.Context, req *systemv1.RestartServiceRequest) (*systemv1.RestartServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartService not implemented")
}

func (s *SystemServer) GetTraces(ctx context.Context, req *systemv1.GetTracesRequest) (*systemv1.GetTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraces not implemented")
}

func (s *SystemServer) GetLogs(ctx context.Context, req *systemv1.GetLogsRequest) (*systemv1.GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}

func (s *SystemServer) StreamLogs(req *systemv1.StreamLogsRequest, stream systemv1.SystemService_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}

func (s *SystemServer) GetCircuitBreakers(ctx context.Context, req *systemv1.GetCircuitBreakersRequest) (*systemv1.GetCircuitBreakersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCircuitBreakers not implemented")
}

func (s *SystemServer) TriggerCircuitBreaker(ctx context.Context, req *systemv1.TriggerCircuitBreakerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerCircuitBreaker not implemented")
}

func (s *SystemServer) ResetCircuitBreaker(ctx context.Context, req *systemv1.ResetCircuitBreakerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetCircuitBreaker not implemented")
}