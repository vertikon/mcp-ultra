package dashboard

import (
	"context"
	"embed"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"mcp-model-ultra/internal/cache"
	"mcp-model-ultra/internal/features"
	"mcp-model-ultra/internal/lifecycle"
	"mcp-model-ultra/internal/logger"
	"mcp-model-ultra/internal/metrics"
	"mcp-model-ultra/internal/observability"
	"mcp-model-ultra/internal/ratelimit"
	"mcp-model-ultra/internal/tracing"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var static embed.FS

// DashboardServer provides custom observability dashboard
type DashboardServer struct {
	config        Config
	logger        logger.Logger
	server        *http.Server
	router        *mux.Router
	websocket     *websocket.Upgrader
	clients       map[string]*websocket.Conn
	
	// Service integrations
	lifecycle     *lifecycle.LifecycleManager
	cache         *cache.DistributedCache
	metrics       *metrics.BusinessMetrics
	features      *features.AdvancedFeatureFlags
	rateLimit     *ratelimit.DistributedRateLimiter
	tracing       *tracing.BusinessTransactionTracer
	observability *observability.TelemetryService
	
	// Templates
	tmpl *template.Template
}

// Config holds dashboard configuration
type Config struct {
	Address         string        `yaml:"address" json:"address"`
	Port            int           `yaml:"port" json:"port"`
	EnableAuth      bool          `yaml:"enable_auth" json:"enable_auth"`
	AuthSecret      string        `yaml:"auth_secret" json:"auth_secret"`
	RefreshInterval time.Duration `yaml:"refresh_interval" json:"refresh_interval"`
	MetricsHistory  time.Duration `yaml:"metrics_history" json:"metrics_history"`
	TLSEnabled      bool          `yaml:"tls_enabled" json:"tls_enabled"`
	TLSCertFile     string        `yaml:"tls_cert_file" json:"tls_cert_file"`
	TLSKeyFile      string        `yaml:"tls_key_file" json:"tls_key_file"`
}

// DefaultConfig returns default dashboard configuration
func DefaultConfig() Config {
	return Config{
		Address:         "0.0.0.0",
		Port:            8090,
		EnableAuth:      true,
		RefreshInterval: 30 * time.Second,
		MetricsHistory:  24 * time.Hour,
		TLSEnabled:      false,
	}
}

// NewDashboardServer creates a new dashboard server
func NewDashboardServer(
	config Config,
	logger logger.Logger,
	lifecycle *lifecycle.LifecycleManager,
	cache *cache.DistributedCache,
	metrics *metrics.BusinessMetrics,
	features *features.AdvancedFeatureFlags,
	rateLimit *ratelimit.DistributedRateLimiter,
	tracing *tracing.BusinessTransactionTracer,
	observability *observability.TelemetryService,
) (*DashboardServer, error) {
	// Parse templates
	tmpl, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	// Create websocket upgrader
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // In production, implement proper CORS
		},
	}

	dashboard := &DashboardServer{
		config:        config,
		logger:        logger,
		websocket:     upgrader,
		clients:       make(map[string]*websocket.Conn),
		lifecycle:     lifecycle,
		cache:         cache,
		metrics:       metrics,
		features:      features,
		rateLimit:     rateLimit,
		tracing:       tracing,
		observability: observability,
		tmpl:          tmpl,
	}

	// Setup router
	dashboard.setupRouter()

	// Create HTTP server
	dashboard.server = &http.Server{
		Addr:         dashboard.config.Address + ":" + strconv.Itoa(dashboard.config.Port),
		Handler:      dashboard.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return dashboard, nil
}

// Start starts the dashboard server
func (ds *DashboardServer) Start(ctx context.Context) error {
	ds.logger.Info("Starting dashboard server", "address", ds.server.Addr)

	// Start websocket broadcaster
	go ds.startWebSocketBroadcaster(ctx)

	// Start server
	go func() {
		var err error
		if ds.config.TLSEnabled {
			err = ds.server.ListenAndServeTLS(ds.config.TLSCertFile, ds.config.TLSKeyFile)
		} else {
			err = ds.server.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			ds.logger.Error("Dashboard server error", "error", err)
		}
	}()

	// Wait for shutdown
	<-ctx.Done()
	return ds.shutdown()
}

// Stop stops the dashboard server
func (ds *DashboardServer) Stop() error {
	return ds.shutdown()
}

func (ds *DashboardServer) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Close websocket connections
	for clientID, conn := range ds.clients {
		conn.Close()
		delete(ds.clients, clientID)
	}

	return ds.server.Shutdown(ctx)
}

func (ds *DashboardServer) setupRouter() {
	ds.router = mux.NewRouter()

	// Static files
	ds.router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.FS(static))),
	)

	// Main dashboard
	ds.router.HandleFunc("/", ds.handleDashboard).Methods("GET")
	ds.router.HandleFunc("/health", ds.handleHealth).Methods("GET")

	// API endpoints
	api := ds.router.PathPrefix("/api/v1").Subrouter()
	if ds.config.EnableAuth {
		api.Use(ds.authMiddleware)
	}

	// System overview
	api.HandleFunc("/overview", ds.handleOverview).Methods("GET")
	api.HandleFunc("/lifecycle", ds.handleLifecycle).Methods("GET")
	
	// Component-specific endpoints
	api.HandleFunc("/cache/stats", ds.handleCacheStats).Methods("GET")
	api.HandleFunc("/metrics/business", ds.handleBusinessMetrics).Methods("GET")
	api.HandleFunc("/features/flags", ds.handleFeatureFlags).Methods("GET")
	api.HandleFunc("/ratelimit/stats", ds.handleRateLimitStats).Methods("GET")
	api.HandleFunc("/tracing/transactions", ds.handleTracingTransactions).Methods("GET")
	
	// Real-time data
	api.HandleFunc("/realtime/metrics", ds.handleRealtimeMetrics).Methods("GET")
	api.HandleFunc("/alerts", ds.handleAlerts).Methods("GET")
	
	// WebSocket endpoint
	ds.router.HandleFunc("/ws", ds.handleWebSocket)
}

// Dashboard page handler
func (ds *DashboardServer) handleDashboard(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Config  Config
		Version string
	}{
		Title:   "MCP Ultra - Observability Dashboard",
		Config:  ds.config,
		Version: "v21.0.0",
	}

	if err := ds.tmpl.ExecuteTemplate(w, "dashboard.html", data); err != nil {
		ds.logger.Error("Template execution error", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// Health check handler
func (ds *DashboardServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"uptime":    time.Since(time.Now()), // This would be calculated properly
		"version":   "v21.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// System overview handler
func (ds *DashboardServer) handleOverview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	overview := SystemOverview{
		Timestamp:    time.Now(),
		SystemHealth: ds.getSystemHealth(ctx),
		Components:   ds.getComponentStatus(ctx),
		Metrics:      ds.getOverviewMetrics(ctx),
		Alerts:       ds.getActiveAlerts(ctx),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(overview)
}

// Lifecycle status handler
func (ds *DashboardServer) handleLifecycle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	status := ds.lifecycle.GetDetailedStatus(ctx)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// Cache statistics handler
func (ds *DashboardServer) handleCacheStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	stats := ds.cache.GetStats(ctx)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// Business metrics handler
func (ds *DashboardServer) handleBusinessMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	// Get time range from query parameters
	startTime, _ := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	endTime, _ := time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	
	if startTime.IsZero() {
		startTime = time.Now().Add(-1 * time.Hour)
	}
	if endTime.IsZero() {
		endTime = time.Now()
	}
	
	metrics := ds.getBusinessMetrics(ctx, startTime, endTime)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

// Feature flags handler
func (ds *DashboardServer) handleFeatureFlags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	flags := ds.features.GetAllFlags(ctx)
	experiments := ds.features.GetActiveExperiments(ctx)
	
	response := map[string]interface{}{
		"flags":       flags,
		"experiments": experiments,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Rate limit statistics handler
func (ds *DashboardServer) handleRateLimitStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	stats := ds.rateLimit.GetStats(ctx)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// Tracing transactions handler
func (ds *DashboardServer) handleTracingTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	// Get query parameters
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 50
	}
	
	status := r.URL.Query().Get("status")
	transactionType := r.URL.Query().Get("type")
	
	transactions := ds.tracing.GetRecentTransactions(ctx, limit, status, transactionType)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// Real-time metrics handler
func (ds *DashboardServer) handleRealtimeMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	metrics := RealtimeMetrics{
		Timestamp:     time.Now(),
		SystemMetrics: ds.getSystemMetrics(ctx),
		Performance:   ds.getPerformanceMetrics(ctx),
		Errors:        ds.getErrorMetrics(ctx),
		Traffic:       ds.getTrafficMetrics(ctx),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

// Alerts handler
func (ds *DashboardServer) handleAlerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	alerts := ds.getActiveAlerts(ctx)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerts)
}

// WebSocket handler for real-time updates
func (ds *DashboardServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ds.websocket.Upgrade(w, r, nil)
	if err != nil {
		ds.logger.Error("WebSocket upgrade error", "error", err)
		return
	}
	defer conn.Close()

	clientID := generateClientID()
	ds.clients[clientID] = conn

	ds.logger.Info("WebSocket client connected", "client_id", clientID)

	// Handle client disconnection
	defer func() {
		delete(ds.clients, clientID)
		ds.logger.Info("WebSocket client disconnected", "client_id", clientID)
	}()

	// Keep connection alive and handle incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				ds.logger.Error("WebSocket error", "error", err)
			}
			break
		}

		// Handle client messages (e.g., subscribe to specific metrics)
		ds.handleWebSocketMessage(clientID, message)
	}
}

// Authentication middleware
func (ds *DashboardServer) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simple token-based authentication
		token := r.Header.Get("Authorization")
		if token != "Bearer "+ds.config.AuthSecret {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func generateClientID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}