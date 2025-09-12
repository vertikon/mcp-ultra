package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"go.uber.org/zap"

	complivancev1 "github.com/vertikon/mcp-ultra/api/grpc/gen/compliance/v1"
	systemv1 "github.com/vertikon/mcp-ultra/api/grpc/gen/system/v1"
	taskv1 "github.com/vertikon/mcp-ultra/api/grpc/gen/task/v1"
	"github.com/vertikon/mcp-ultra/internal/compliance"
	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/features"
	"github.com/vertikon/mcp-ultra/internal/observability"
	"github.com/vertikon/mcp-ultra/internal/security"
	"github.com/vertikon/mcp-ultra/internal/services"
)

// GRPCServer represents the main gRPC server
type GRPCServer struct {
	config              *config.Config
	server              *grpc.Server
	listener            net.Listener
	logger              *zap.Logger
	healthServer        *health.Server
	
	// Service dependencies
	taskService         *services.TaskService
	complianceFramework *compliance.ComplianceFramework
	obsService          *observability.Service
	flagManager         *features.FlagManager
	tlsManager          *security.TLSManager
}

// NewGRPCServer creates a new gRPC server with all services registered
func NewGRPCServer(
	config *config.Config,
	taskService *services.TaskService,
	complianceFramework *compliance.ComplianceFramework,
	obsService *observability.Service,
	flagManager *features.FlagManager,
	tlsManager *security.TLSManager,
	logger *zap.Logger,
) (*GRPCServer, error) {
	// Create listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to create listener: %w", err)
	}

	// Configure server options
	serverOpts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(config.GRPC.MaxRecvMessageSize),
		grpc.MaxSendMsgSize(config.GRPC.MaxSendMessageSize),
		grpc.ConnectionTimeout(config.GRPC.ConnectionTimeout),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     config.GRPC.Keepalive.MaxConnectionIdle,
			MaxConnectionAge:      config.GRPC.Keepalive.MaxConnectionAge,
			MaxConnectionAgeGrace: config.GRPC.Keepalive.MaxConnectionAgeGrace,
			Time:                  config.GRPC.Keepalive.Time,
			Timeout:               config.GRPC.Keepalive.Timeout,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             config.GRPC.Keepalive.MinTime,
			PermitWithoutStream: config.GRPC.Keepalive.PermitWithoutStream,
		}),
	}

	// Add TLS credentials if configured
	if config.Security.TLS.CertFile != "" && config.Security.TLS.KeyFile != "" {
		tlsConfig, err := tlsManager.GetServerTLSConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get TLS config: %w", err)
		}
		creds := credentials.NewTLS(tlsConfig)
		serverOpts = append(serverOpts, grpc.Creds(creds))
		logger.Info("gRPC server configured with TLS")
	}

	// Add interceptors
	serverOpts = append(serverOpts,
		grpc.ChainUnaryInterceptor(
			// Add observability interceptor first
			observabilityUnaryInterceptor(obsService, logger),
			// Add authentication interceptor
			authenticationUnaryInterceptor(config, logger),
			// Add rate limiting interceptor
			rateLimitingUnaryInterceptor(config, logger),
			// Add validation interceptor
			validationUnaryInterceptor(logger),
		),
		grpc.ChainStreamInterceptor(
			// Add observability interceptor first
			observabilityStreamInterceptor(obsService, logger),
			// Add authentication interceptor
			authenticationStreamInterceptor(config, logger),
			// Add rate limiting interceptor
			rateLimitingStreamInterceptor(config, logger),
		),
	)

	// Create gRPC server
	server := grpc.NewServer(serverOpts...)

	// Create health server
	healthServer := health.NewServer()

	grpcServer := &GRPCServer{
		config:              config,
		server:              server,
		listener:            listener,
		logger:              logger,
		healthServer:        healthServer,
		taskService:         taskService,
		complianceFramework: complianceFramework,
		obsService:          obsService,
		flagManager:         flagManager,
		tlsManager:          tlsManager,
	}

	// Register services
	if err := grpcServer.registerServices(); err != nil {
		return nil, fmt.Errorf("failed to register services: %w", err)
	}

	return grpcServer, nil
}

// registerServices registers all gRPC services
func (s *GRPCServer) registerServices() error {
	// Register health service
	grpc_health_v1.RegisterHealthServer(s.server, s.healthServer)
	
	// Set initial health status for all services
	s.healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	s.healthServer.SetServingStatus("task.v1.TaskService", grpc_health_v1.HealthCheckResponse_SERVING)
	s.healthServer.SetServingStatus("compliance.v1.ComplianceService", grpc_health_v1.HealthCheckResponse_SERVING)
	s.healthServer.SetServingStatus("system.v1.SystemService", grpc_health_v1.HealthCheckResponse_SERVING)

	// Register TaskService
	taskServer := NewTaskServer(s.taskService, s.logger)
	taskv1.RegisterTaskServiceServer(s.server, taskServer)
	s.logger.Info("Registered TaskService")

	// Register ComplianceService
	complianceServer := NewComplianceServer(s.complianceFramework, s.logger)
	complivancev1.RegisterComplianceServiceServer(s.server, complianceServer)
	s.logger.Info("Registered ComplianceService")

	// Register SystemService
	systemServer := NewSystemServer(s.config, s.obsService, s.flagManager, s.logger)
	systemv1.RegisterSystemServiceServer(s.server, systemServer)
	s.logger.Info("Registered SystemService")

	// Enable reflection for development
	if s.config.Environment == "development" {
		reflection.Register(s.server)
		s.logger.Info("Enabled gRPC reflection for development")
	}

	s.logger.Info("All gRPC services registered successfully")
	return nil
}

// Start starts the gRPC server
func (s *GRPCServer) Start(ctx context.Context) error {
	s.logger.Info("Starting gRPC server",
		zap.String("address", s.listener.Addr().String()),
		zap.Bool("tls_enabled", s.config.Security.TLS.CertFile != ""),
		zap.Bool("reflection_enabled", s.config.Environment == "development"),
	)

	// Start server in a goroutine
	errChan := make(chan error, 1)
	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			errChan <- fmt.Errorf("gRPC server failed: %w", err)
		}
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.Done():
		s.logger.Info("Shutting down gRPC server due to context cancellation")
		return s.Stop(ctx)
	case err := <-errChan:
		return err
	}
}

// Stop gracefully stops the gRPC server
func (s *GRPCServer) Stop(ctx context.Context) error {
	s.logger.Info("Stopping gRPC server")

	// Set all services to not serving
	s.healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	s.healthServer.SetServingStatus("task.v1.TaskService", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	s.healthServer.SetServingStatus("compliance.v1.ComplianceService", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	s.healthServer.SetServingStatus("system.v1.SystemService", grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	// Create a channel to signal when shutdown is complete
	done := make(chan struct{})

	// Start graceful shutdown in a goroutine
	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	// Wait for graceful shutdown or timeout
	shutdownTimeout := 30 * time.Second
	if s.config.GRPC.ShutdownTimeout > 0 {
		shutdownTimeout = s.config.GRPC.ShutdownTimeout
	}

	select {
	case <-done:
		s.logger.Info("gRPC server stopped gracefully")
		return nil
	case <-time.After(shutdownTimeout):
		s.logger.Warn("gRPC server graceful shutdown timed out, forcing stop")
		s.server.Stop()
		return nil
	case <-ctx.Done():
		s.logger.Warn("Context cancelled during gRPC server shutdown, forcing stop")
		s.server.Stop()
		return ctx.Err()
	}
}

// GetListener returns the server's listener (useful for testing)
func (s *GRPCServer) GetListener() net.Listener {
	return s.listener
}

// UpdateHealthStatus updates the health status of a service
func (s *GRPCServer) UpdateHealthStatus(service string, status grpc_health_v1.HealthCheckResponse_ServingStatus) {
	s.healthServer.SetServingStatus(service, status)
	s.logger.Debug("Updated service health status",
		zap.String("service", service),
		zap.String("status", status.String()),
	)
}

// GetAddress returns the server's listening address
func (s *GRPCServer) GetAddress() string {
	return s.listener.Addr().String()
}

// IsHealthy checks if all registered services are healthy
func (s *GRPCServer) IsHealthy(ctx context.Context) bool {
	// TODO: Implement actual health checks for all services
	// For now, return true if server is running
	return s.server != nil
}

// GetMetrics returns gRPC server metrics
func (s *GRPCServer) GetMetrics(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"address":           s.GetAddress(),
		"tls_enabled":       s.config.Security.TLS.CertFile != "",
		"reflection_enabled": s.config.Environment == "development",
		"services": []string{
			"grpc.health.v1.Health",
			"task.v1.TaskService", 
			"compliance.v1.ComplianceService",
			"system.v1.SystemService",
		},
	}
}