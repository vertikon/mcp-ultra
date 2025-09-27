package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/lifecycle"
	"github.com/vertikon/mcp-ultra/internal/services"
	"github.com/vertikon/mcp-ultra/internal/repository/postgres"
	"github.com/vertikon/mcp-ultra/internal/repository/redis"
)

// Handlers holds all HTTP handlers
type Handlers struct {
	taskHandlers        *TaskHandlers
	featureFlagHandlers *FeatureFlagHandlers
	logger              *zap.Logger
	config              *config.Config
	lifecycleManager    *lifecycle.Manager
}

// NewHandlers creates new HTTP handlers
func NewHandlers(logger *zap.Logger, cfg *config.Config, lifecycleManager *lifecycle.Manager) (*Handlers, error) {
	// Initialize database connections
	pgRepo, err := postgres.NewConnection(cfg.Database.PostgreSQL)
	if err != nil {
		return nil, err
	}

	redisRepo, err := redis.NewConnection(cfg.Database.Redis)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	taskRepo := postgres.NewTaskRepository(pgRepo, logger)
	cacheRepo := redis.NewCacheRepository(redisRepo, logger)

	// Initialize services
	taskService := services.NewTaskService(taskRepo, cacheRepo, logger)

	// Initialize handlers
	taskHandlers := NewTaskHandlers(taskService, logger)
	featureFlagHandlers := NewFeatureFlagHandlers(logger)

	return &Handlers{
		taskHandlers:        taskHandlers,
		featureFlagHandlers: featureFlagHandlers,
		logger:              logger,
		config:              cfg,
		lifecycleManager:    lifecycleManager,
	}, nil
}

// RegisterRoutes registers all HTTP routes
func (h *Handlers) RegisterRoutes(r chi.Router) {
	// Task routes
	r.Route("/api/v1/tasks", func(r chi.Router) {
		r.Post("/", h.taskHandlers.CreateTask)
		r.Get("/", h.taskHandlers.ListTasks)
		r.Get("/{id}", h.taskHandlers.GetTask)
		r.Put("/{id}", h.taskHandlers.UpdateTask)
		r.Delete("/{id}", h.taskHandlers.DeleteTask)
		r.Post("/{id}/complete", h.taskHandlers.CompleteTask)
		r.Get("/status/{status}", h.taskHandlers.GetTasksByStatus)
		r.Get("/assignee/{assigneeId}", h.taskHandlers.GetTasksByAssignee)
	})

	// Feature flag routes
	r.Route("/api/v1/features", func(r chi.Router) {
		r.Get("/", h.featureFlagHandlers.ListFlags)
		r.Get("/{key}", h.featureFlagHandlers.GetFlag)
		r.Post("/{key}/toggle", h.featureFlagHandlers.ToggleFlag)
	})

	// System routes
	r.Route("/api/v1/system", func(r chi.Router) {
		r.Get("/info", h.handleSystemInfo)
		r.Get("/metrics", h.handleMetrics)
	})

	// Swagger documentation
	r.Get("/swagger/*", h.handleSwagger)
}

// handleSystemInfo returns system information
func (h *Handlers) handleSystemInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"service":     h.config.Telemetry.ServiceName,
		"version":     h.config.Telemetry.ServiceVersion,
		"environment": h.config.Telemetry.Environment,
		"ready":       h.lifecycleManager.IsReady(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// JSON encoding would be handled by a utility function
}

// handleMetrics handles metrics endpoint
func (h *Handlers) handleMetrics(w http.ResponseWriter, r *http.Request) {
	// Metrics would be handled by Prometheus middleware
	w.WriteHeader(http.StatusOK)
}

// handleSwagger serves Swagger documentation
func (h *Handlers) handleSwagger(w http.ResponseWriter, r *http.Request) {
	// Swagger UI would be served here
	w.WriteHeader(http.StatusOK)
}

// Shutdown gracefully shuts down handlers
func (h *Handlers) Shutdown(ctx context.Context) error {
	h.logger.Info("Shutting down HTTP handlers")
	return nil
}