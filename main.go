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

	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/handlers"
	"github.com/vertikon/mcp-ultra/pkg/httpx"
	"github.com/vertikon/mcp-ultra/pkg/logger"
	"github.com/vertikon/mcp-ultra/pkg/metrics"
	"github.com/vertikon/mcp-ultra/pkg/version"
)

func main() {
	// Initialize logger
	zapLog, err := logger.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func() {
		if syncErr := zapLog.Sync(); syncErr != nil {
			// Ignore sync errors on shutdown (common on Windows)
			log.Printf("Warning: failed to sync logger: %v", syncErr)
		}
	}()

	zapLog.Info("Starting MCP Ultra service",
		logger.String("version", version.Version),
		logger.String("build_date", version.BuildDate),
		logger.String("commit", version.GitCommit),
	)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		zapLog.Fatal("Failed to load configuration", logger.Error(err))
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
	router.Use(httpx.DefaultCORS())

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
		zapLog.Info("Starting HTTP server",
			logger.String("address", server.Addr),
			logger.Int("port", cfg.Server.Port),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zapLog.Fatal("Failed to start HTTP server", logger.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zapLog.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := server.Shutdown(ctx); err != nil {
		zapLog.Error("Server forced to shutdown", logger.Error(err))
	}

	zapLog.Info("Server exited")
}
