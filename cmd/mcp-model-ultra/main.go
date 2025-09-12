package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/vertikon/mcp-ultra/internal/compliance"
	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/events"
	"github.com/vertikon/mcp-ultra/internal/features"
	grpcserver "github.com/vertikon/mcp-ultra/internal/grpc/server"
	httphandlers "github.com/vertikon/mcp-ultra/internal/handlers/http"
	"github.com/vertikon/mcp-ultra/internal/lifecycle"
	"github.com/vertikon/mcp-ultra/internal/observability"
	"github.com/vertikon/mcp-ultra/internal/repository/postgres"
	"github.com/vertikon/mcp-ultra/internal/repository/redis"
	"github.com/vertikon/mcp-ultra/internal/security"
	"github.com/vertikon/mcp-ultra/internal/services"
	"github.com/vertikon/mcp-ultra/internal/slo"
	"github.com/vertikon/mcp-ultra/pkg/logger"
	"github.com/vertikon/mcp-ultra/pkg/version"
)

func main() {
	// Initialize logger
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	// Create logger adapter
	log := logger.NewZapAdapter(zapLogger)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		zapLogger.Fatal("Failed to load config", zap.Error(err))
	}

	// Initialize observability service with telemetry
	telemetryService, err := observability.NewTelemetryService(cfg.Telemetry, log)
	if err != nil {
		zapLogger.Fatal("Failed to initialize telemetry service", zap.Error(err))
	}

	// Initialize lifecycle manager
	lifecycleConfig := lifecycle.DefaultConfig()
	lifecycleManager := lifecycle.NewLifecycleManager(lifecycleConfig, log, telemetryService)
	
	// Initialize health monitor
	healthConfig := lifecycle.DefaultHealthConfig()
	healthMonitor := lifecycle.NewHealthMonitor(healthConfig, version.Version, log)
	
	// Initialize SLO monitor
	sloMonitor := slo.NewMonitor(slo.DefaultConfig(), log)
	
	// Initialize operations manager
	operationsConfig := lifecycle.DefaultOperationsConfig()
	operationsManager := lifecycle.NewOperationsManager(operationsConfig, log)
	
	// Initialize deployment automation
	deploymentConfig := lifecycle.DeploymentConfig{
		Strategy:         lifecycle.DeploymentRolling,
		Environment:      cfg.Environment,
		Namespace:        "mcp-ultra",
		Image:            "vertikon/mcp-ultra",
		Tag:              version.Version,
		ValidateConfig:   true,
		ValidateImage:    true,
		ValidateResources: true,
		ProgressTimeout:  10 * time.Minute,
		ReadinessTimeout: 5 * time.Minute,
		LivenessTimeout:  5 * time.Minute,
		AutoRollback:     true,
		EnableMetrics:    true,
		EnableAlerting:   true,
		ManifestPath:     "./deploy/k8s",
		RollbackThresholds: lifecycle.RollbackThresholds{
			ErrorRate:        5.0,
			ResponseTime:     2 * time.Second,
			HealthCheckFails: 3,
			TimeWindow:       5 * time.Minute,
		},
	}
	deploymentAutomation := lifecycle.NewDeploymentAutomation(deploymentConfig, log)
	
	// Initialize observability service
	obsService, err := observability.NewService(cfg.Telemetry, zapLogger)
	if err != nil {
		zapLogger.Fatal("Failed to initialize observability service", zap.Error(err))
	}
	
	// Start observability service
	ctx := context.Background()
	if err := obsService.Start(ctx); err != nil {
		zapLogger.Fatal("Failed to start observability service", zap.Error(err))
	}
	defer obsService.Stop(ctx)

	// Initialize compliance framework
	complianceFramework, err := compliance.NewComplianceFramework(convertToComplianceConfig(cfg.Compliance), zapLogger)
	if err != nil {
		zapLogger.Fatal("Failed to initialize compliance framework", zap.Error(err))
	}

	// Database connections
	db, err := postgres.Connect(cfg.Database.PostgreSQL)
	if err != nil {
		zapLogger.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
	}
	defer db.Close()

	// Redis connection
	cacheClient := redis.NewClient(cfg.Database.Redis)
	defer cacheClient.Close()

	// Test Redis connection
	if err := redis.Ping(cacheClient); err != nil {
		zapLogger.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	// Initialize NATS event bus
	eventBus, err := events.NewNATSEventBus(cfg.NATS.URL, zapLogger)
	if err != nil {
		zapLogger.Fatal("Failed to initialize NATS event bus", zap.Error(err))
	}
	defer eventBus.Close()

	// Initialize security services
	opaService := security.NewOPAService(cfg.Security.OPA, zapLogger)
	vaultService := security.NewVaultService(cfg.Security.Vault, zapLogger)
	authService := security.NewAuthService(cfg.Security.Auth, zapLogger, opaService)
	tlsManager := security.NewTLSManager(cfg.Security.TLS, zapLogger)

	// Initialize repositories
	taskRepo := postgres.NewTaskRepository(db)
	cacheRepo := redis.NewCacheRepository(cacheClient)
	
	// Initialize feature flags
	flagManager := features.NewFlagManager(nil, cacheRepo, zapLogger) // TODO: Implement flag repo
	defer flagManager.Stop()
	
	// Initialize services (compliance integration will be added to TaskService later)
	taskService := services.NewTaskService(taskRepo, nil, nil, cacheRepo, zapLogger, eventBus)
	
	// Register lifecycle components
	lifecycleManager.RegisterComponent(&DatabaseComponent{DB: db})
	lifecycleManager.RegisterComponent(&RedisComponent{Client: cacheClient})
	lifecycleManager.RegisterComponent(&EventBusComponent{Bus: eventBus})
	lifecycleManager.RegisterComponent(&ObservabilityComponent{Service: obsService})
	lifecycleManager.RegisterComponent(&ComplianceComponent{Framework: complianceFramework})
	
	// Register health checkers
	healthMonitor.RegisterChecker(lifecycle.NewDatabaseHealthChecker("postgresql"))
	healthMonitor.RegisterChecker(lifecycle.NewRedisHealthChecker("redis"))
	
	// Register SLO default configuration
	slos := slo.DefaultSLOs()
	for _, sloConfig := range slos {
		if err := sloMonitor.AddSLO(sloConfig); err != nil {
			zapLogger.Error("Failed to add SLO", zap.String("slo", sloConfig.Name), zap.Error(err))
		}
	}
	
	// Register operation executors
	operationsManager.RegisterExecutor(lifecycle.OperationMaintenance, lifecycle.NewMaintenanceExecutor(log))
	
	// Start lifecycle systems
	if err := lifecycleManager.Start(ctx); err != nil {
		zapLogger.Fatal("Failed to start lifecycle manager", zap.Error(err))
	}
	defer lifecycleManager.Stop(ctx)
	
	if err := healthMonitor.Start(); err != nil {
		zapLogger.Fatal("Failed to start health monitor", zap.Error(err))
	}
	defer healthMonitor.Stop()
	
	if err := sloMonitor.Start(ctx); err != nil {
		zapLogger.Fatal("Failed to start SLO monitor", zap.Error(err))
	}
	defer sloMonitor.Stop(ctx)
	
	if err := operationsManager.Start(); err != nil {
		zapLogger.Fatal("Failed to start operations manager", zap.Error(err))
	}
	defer operationsManager.Stop()
	
	// Initialize gRPC server
	grpcSrv, err := grpcserver.NewGRPCServer(
		cfg,
		taskService,
		complianceFramework,
		obsService,
		flagManager,
		tlsManager,
		zapLogger,
	)
	if err != nil {
		zapLogger.Fatal("Failed to initialize gRPC server", zap.Error(err))
	}

	// Setup HTTP router with security and observability middleware
	router := httphandlers.NewRouter(taskService, flagManager, zapLogger)
	
	// Add lifecycle endpoints
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		report := healthMonitor.GetHealth(r.Context())
		w.Header().Set("Content-Type", "application/json")
		if report.Status == lifecycle.HealthStatusHealthy {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		// TODO: Add JSON encoder for the report
	})
	
	router.HandleFunc("/lifecycle/status", func(w http.ResponseWriter, r *http.Request) {
		metrics := lifecycleManager.GetMetrics()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// TODO: Add JSON encoder for the metrics
	})
	
	router.HandleFunc("/slo/status", func(w http.ResponseWriter, r *http.Request) {
		status := sloMonitor.GetStatus()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// TODO: Add JSON encoder for the status
	})
	
	// Apply observability middleware first (outermost)
	instrumentedRouter := obsService.HTTPMiddleware()(router)
	
	// Apply security middleware
	secureRouter := authService.JWTMiddleware(instrumentedRouter)
	if cfg.Security.TLS.ClientAuth {
		secureRouter = tlsManager.MTLSMiddleware(secureRouter)
	}
	
	// Add metrics endpoint
	router.Handle("/metrics", promhttp.Handler())

	// HTTP Server configuration with TLS
	httpSrv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      secureRouter,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Configure TLS if enabled
	if cfg.Security.TLS.CertFile != "" && cfg.Security.TLS.KeyFile != "" {
		tlsConfig, err := tlsManager.GetServerTLSConfig()
		if err != nil {
			logger.Fatal("Failed to configure TLS", zap.Error(err))
		}
		httpSrv.TLSConfig = tlsConfig
	}

	// Setup event handlers
	taskEventHandler := events.NewTaskEventHandler(zapLogger)
	if _, err := eventBus.SubscribeQueue("task.*", "task-handlers", taskEventHandler); err != nil {
		zapLogger.Error("Failed to subscribe to task events", zap.Error(err))
	}

	// Start both servers concurrently
	var wg sync.WaitGroup
	
	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcSrv.Start(ctx); err != nil {
			zapLogger.Error("gRPC server failed", zap.Error(err))
		}
	}()
	
	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		zapLogger.Info("Shutting down gracefully...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Shutdown gRPC server first
		if err := grpcSrv.Stop(shutdownCtx); err != nil {
			zapLogger.Error("gRPC server shutdown error", zap.Error(err))
		}

		// Then shutdown HTTP server
		if err := httpSrv.Shutdown(shutdownCtx); err != nil {
			zapLogger.Error("HTTP server shutdown error", zap.Error(err))
		}
	}()

	// Health checks for external services
	zapLogger.Info("Performing health checks...")
	if err := opaService.HealthCheck(context.Background()); err != nil {
		zapLogger.Warn("OPA health check failed", zap.Error(err))
	}
	if err := vaultService.HealthCheck(context.Background()); err != nil {
		zapLogger.Warn("Vault health check failed", zap.Error(err))
	}

	// Log observability health status
	healthStatus := obsService.HealthCheck()
	zapLogger.Info("Observability status", zap.Any("health", healthStatus))

	// Log compliance status
	complianceStatus, _ := complianceFramework.GetComplianceStatus(ctx)
	zapLogger.Info("Compliance status", zap.Any("compliance", complianceStatus))
	
	// Log lifecycle status
	lifecycleMetrics := lifecycleManager.GetMetrics()
	zapLogger.Info("Lifecycle status", 
		zap.String("state", lifecycleMetrics.State),
		zap.Int("components", lifecycleMetrics.ComponentCount),
		zap.Int("healthy", lifecycleMetrics.HealthyComponents),
		zap.Duration("uptime", lifecycleMetrics.Uptime))

	// Start HTTP server  
	zapLogger.Info("Starting MCP Ultra", 
		zap.Int("http_port", cfg.Server.Port),
		zap.Int("grpc_port", cfg.GRPC.Port),
		zap.String("version", version.Version),
		zap.Bool("tls_enabled", httpSrv.TLSConfig != nil),
		zap.Bool("mtls_enabled", cfg.Security.TLS.ClientAuth),
		zap.Bool("observability_enabled", cfg.Telemetry.Enabled),
		zap.Bool("tracing_enabled", cfg.Telemetry.Tracing.Enabled),
		zap.Bool("metrics_enabled", cfg.Telemetry.Metrics.Enabled),
		zap.Bool("compliance_enabled", cfg.Compliance.Enabled),
		zap.Bool("lgpd_enabled", cfg.Compliance.LGPD.Enabled),
		zap.Bool("gdpr_enabled", cfg.Compliance.GDPR.Enabled),
		zap.Bool("grpc_enabled", true),
		zap.Bool("lifecycle_enabled", true),
		zap.Bool("health_monitoring_enabled", true),
		zap.Bool("slo_monitoring_enabled", true),
		zap.Bool("operations_management_enabled", true))
	
	if httpSrv.TLSConfig != nil {
		if err := httpSrv.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			zapLogger.Fatal("HTTPS server failed", zap.Error(err))
		}
	} else {
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zapLogger.Fatal("HTTP server failed", zap.Error(err))
		}
	}

	// Wait for gRPC server to complete
	wg.Wait()

	zapLogger.Info("All servers stopped")
}

// convertToComplianceConfig converts config.ComplianceConfig to compliance.ComplianceConfig
func convertToComplianceConfig(cfg config.ComplianceConfig) compliance.ComplianceConfig {
	return compliance.ComplianceConfig{
		Enabled:       cfg.Enabled,
		DefaultRegion: cfg.DefaultRegion,
		PIIDetection: compliance.PIIDetectionConfig{
			Enabled:           cfg.PIIDetection.Enabled,
			ScanFields:        cfg.PIIDetection.ScanFields,
			ClassificationAPI: cfg.PIIDetection.ClassificationAPI,
			Confidence:        cfg.PIIDetection.Confidence,
			AutoMask:          cfg.PIIDetection.AutoMask,
		},
		Consent: compliance.ConsentConfig{
			Enabled:         cfg.Consent.Enabled,
			DefaultPurposes: cfg.Consent.DefaultPurposes,
			TTL:             cfg.Consent.TTL,
			GranularLevel:   cfg.Consent.GranularLevel,
		},
		DataRetention: compliance.DataRetentionConfig{
			Enabled:         cfg.DataRetention.Enabled,
			DefaultPeriod:   cfg.DataRetention.DefaultPeriod,
			CategoryPeriods: cfg.DataRetention.CategoryPeriods,
			AutoDelete:      cfg.DataRetention.AutoDelete,
			BackupRetention: cfg.DataRetention.BackupRetention,
		},
		AuditLogging: compliance.AuditLoggingConfig{
			Enabled:           cfg.AuditLogging.Enabled,
			DetailLevel:       cfg.AuditLogging.DetailLevel,
			RetentionPeriod:   cfg.AuditLogging.RetentionPeriod,
			EncryptionEnabled: cfg.AuditLogging.EncryptionEnabled,
			ExternalLogging:   cfg.AuditLogging.ExternalLogging,
			ExternalEndpoint:  cfg.AuditLogging.ExternalEndpoint,
		},
		LGPD: compliance.LGPDConfig{
			Enabled:         cfg.LGPD.Enabled,
			DPOContact:      cfg.LGPD.DPOContact,
			LegalBasis:      cfg.LGPD.LegalBasis,
			DataCategories:  cfg.LGPD.DataCategories,
			SharedThirdParty: cfg.LGPD.SharedThirdParty,
		},
		GDPR: compliance.GDPRConfig{
			Enabled:           cfg.GDPR.Enabled,
			DPOContact:        cfg.GDPR.DPOContact,
			LegalBasis:        cfg.GDPR.LegalBasis,
			DataCategories:    cfg.GDPR.DataCategories,
			CrossBorderTransfer: cfg.GDPR.CrossBorderTransfer,
			AdequacyDecisions: cfg.GDPR.AdequacyDecisions,
		},
		Anonymization: compliance.AnonymizationConfig{
			Enabled:    cfg.Anonymization.Enabled,
			Methods:    cfg.Anonymization.Methods,
			HashSalt:   cfg.Anonymization.HashSalt,
			Reversible: cfg.Anonymization.Reversible,
			KAnonymity: cfg.Anonymization.KAnonymity,
			Algorithms: cfg.Anonymization.Algorithms,
		},
		DataRights: compliance.DataRightsConfig{
			Enabled:              cfg.DataRights.Enabled,
			ResponseTime:         cfg.DataRights.ResponseTime,
			AutoFulfillment:      cfg.DataRights.AutoFulfillment,
			VerificationRequired: cfg.DataRights.VerificationRequired,
			NotificationChannels: cfg.DataRights.NotificationChannels,
		},
	}
}
