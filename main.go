package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vertikon/mcp-ultra-fix/pkg/httpx"
	"github.com/vertikon/mcp-ultra-fix/pkg/logger"
	"github.com/vertikon/mcp-ultra-fix/pkg/metrics"
	"github.com/vertikon/mcp-ultra-fix/pkg/version"
	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/handlers"
)

func main() {
	// Initialize logger
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func() {
		if syncErr := logger.Sync(); syncErr != nil {
			// Ignore sync errors on shutdown (common on Windows)
			log.Printf("Warning: failed to sync logger: %v", syncErr)
		}
	}()

	logger.Info("Starting MCP Ultra service",
		"version", version.Version,
		"build_date", version.BuildDate,
		"commit", version.GitCommit,
	)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", "error", err)
	}

	// Initialize HTTP router
	router := httpx.NewRouter()

	// Add middleware
	router.Use(httpx.RequestID)
	router.Use(httpx.RealIP)
	router.Use(httpx.Logger)
	router.Use(httpx.Recoverer)
	router.Use(httpx.Timeout(60 * time.Second))

	// CORS configuration
	router.Use(httpx.CORS(httpx.CORSOptions{
		AllowedOrigins:   []string{"*"}, // Configure appropriately for production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Initialize health handler
	healthHandler := handlers.NewHealthHandler()

	// Register routes
	router.Get("/livez", healthHandler.Livez)
	router.Get("/readyz", healthHandler.Readyz)
	router.Get("/health", healthHandler.Health)
	router.Get("/metrics", metrics.Handler().ServeHTTP)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info("Starting HTTP server",
			"address", server.Addr,
			"port", cfg.Server.Port,
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start HTTP server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}

	logger.Info("Server exited")
}
