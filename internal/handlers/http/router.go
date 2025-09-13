package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"github.com/vertikon/mcp-ultra/internal/features"
	"github.com/vertikon/mcp-ultra/internal/services"
	"github.com/vertikon/mcp-ultra/internal/telemetry"
)

// Router creates and configures the HTTP router
func NewRouter(
	taskService *services.TaskService,
	flagManager *features.FlagManager,
	healthService *HealthService,
	logger *zap.Logger,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(telemetry.HTTPMetrics)

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Configure for production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health endpoints using comprehensive health service
	if healthService != nil {
		healthService.RegisterRoutes(r)
	} else {
		// Fallback to basic health checks
		r.Get("/healthz", healthCheck)
		r.Get("/readyz", readinessCheck)
	}

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Task routes
		r.Mount("/tasks", TaskRoutes(taskService, logger))
		
		// Feature flag routes
		r.Mount("/flags", FeatureFlagRoutes(flagManager, logger))
	})

	return r
}

// TaskRoutes creates task-related routes
func TaskRoutes(taskService *services.TaskService, logger *zap.Logger) chi.Router {
	r := chi.NewRouter()
	handlers := NewTaskHandlers(taskService, logger)

	r.Post("/", handlers.CreateTask)
	r.Get("/", handlers.ListTasks)
	r.Get("/{id}", handlers.GetTask)
	r.Put("/{id}", handlers.UpdateTask)
	r.Delete("/{id}", handlers.DeleteTask)
	r.Post("/{id}/complete", handlers.CompleteTask)
	r.Get("/status/{status}", handlers.GetTasksByStatus)
	r.Get("/assignee/{assigneeId}", handlers.GetTasksByAssignee)

	return r
}

// FeatureFlagRoutes creates feature flag routes
func FeatureFlagRoutes(flagManager *features.FlagManager, logger *zap.Logger) chi.Router {
	r := chi.NewRouter()
	handlers := NewFeatureFlagHandlers(flagManager, logger)

	r.Get("/", handlers.ListFlags)
	r.Get("/{key}", handlers.GetFlag)
	r.Post("/", handlers.CreateFlag)
	r.Put("/{key}", handlers.UpdateFlag)
	r.Delete("/{key}", handlers.DeleteFlag)
	r.Post("/{key}/evaluate", handlers.EvaluateFlag)

	return r
}

// Health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
}

// Readiness check endpoint
func readinessCheck(w http.ResponseWriter, r *http.Request) {
	// Add checks for dependencies (database, cache, etc.)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ready", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
}