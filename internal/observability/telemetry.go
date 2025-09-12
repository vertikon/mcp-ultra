package observability

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TelemetryConfig holds telemetry configuration
type TelemetryConfig struct {
	ServiceName     string        `yaml:"service_name"`
	ServiceVersion  string        `yaml:"service_version"`
	Environment     string        `yaml:"environment"`
	JaegerEndpoint  string        `yaml:"jaeger_endpoint"`
	OTLPEndpoint    string        `yaml:"otlp_endpoint"`
	MetricsPort     int           `yaml:"metrics_port"`
	SamplingRate    float64       `yaml:"sampling_rate"`
	BatchTimeout    time.Duration `yaml:"batch_timeout"`
	BatchSize       int           `yaml:"batch_size"`
	Enabled         bool          `yaml:"enabled"`
	Debug           bool          `yaml:"debug"`
}

// TelemetryService manages OpenTelemetry instrumentation
type TelemetryService struct {
	config         TelemetryConfig
	logger         *zap.Logger
	tracerProvider trace.TracerProvider
	meterProvider  metric.MeterProvider
	tracer         trace.Tracer
	meter          metric.Meter
	
	// Business metrics
	requestCounter    metric.Int64Counter
	requestDuration   metric.Float64Histogram
	activeConnections metric.Int64UpDownCounter
	errorCounter      metric.Int64Counter
	taskMetrics       *TaskMetrics
	
	// System metrics
	cpuUsage     metric.Float64ObservableGauge
	memoryUsage  metric.Float64ObservableGauge
	goroutines   metric.Int64ObservableGauge
}

// TaskMetrics holds task-specific metrics
type TaskMetrics struct {
	taskCreated   metric.Int64Counter
	taskCompleted metric.Int64Counter
	taskFailed    metric.Int64Counter
	taskDuration  metric.Float64Histogram
	
	tasksByStatus metric.Int64ObservableGauge
	tasksByPriority metric.Int64ObservableGauge
}

// NewTelemetryService creates a new telemetry service
func NewTelemetryService(config TelemetryConfig, logger *zap.Logger) (*TelemetryService, error) {
	if !config.Enabled {
		logger.Info("Telemetry disabled")
		return &TelemetryService{
			config: config,
			logger: logger,
		}, nil
	}

	ts := &TelemetryService{
		config: config,
		logger: logger,
	}

	// Initialize resource
	res, err := ts.initResource()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resource: %w", err)
	}

	// Initialize tracing
	if err := ts.initTracing(res); err != nil {
		return nil, fmt.Errorf("failed to initialize tracing: %w", err)
	}

	// Initialize metrics
	if err := ts.initMetrics(res); err != nil {
		return nil, fmt.Errorf("failed to initialize metrics: %w", err)
	}

	// Initialize business metrics
	if err := ts.initBusinessMetrics(); err != nil {
		return nil, fmt.Errorf("failed to initialize business metrics: %w", err)
	}

	// Initialize system metrics
	if err := ts.initSystemMetrics(); err != nil {
		return nil, fmt.Errorf("failed to initialize system metrics: %w", err)
	}

	logger.Info("Telemetry initialized successfully",
		zap.String("service", config.ServiceName),
		zap.String("version", config.ServiceVersion),
		zap.String("environment", config.Environment))

	return ts, nil
}

// initResource creates the OpenTelemetry resource
func (ts *TelemetryService) initResource() (*resource.Resource, error) {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(ts.config.ServiceName),
		semconv.ServiceVersion(ts.config.ServiceVersion),
		semconv.DeploymentEnvironment(ts.config.Environment),
		attribute.String("service.instance.id", generateInstanceID()),
		attribute.String("telemetry.sdk.name", "opentelemetry"),
		attribute.String("telemetry.sdk.language", "go"),
		attribute.String("telemetry.sdk.version", otel.Version()),
	), nil
}

// initTracing sets up distributed tracing
func (ts *TelemetryService) initTracing(res *resource.Resource) error {
	var exporter sdktrace.SpanExporter
	var err error

	// Choose exporter based on configuration
	if ts.config.JaegerEndpoint != "" {
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint(ts.config.JaegerEndpoint),
		))
		if err != nil {
			return fmt.Errorf("failed to create Jaeger exporter: %w", err)
		}
		ts.logger.Info("Using Jaeger exporter", zap.String("endpoint", ts.config.JaegerEndpoint))
	} else if ts.config.OTLPEndpoint != "" {
		exporter, err = otlptracehttp.New(context.Background(),
			otlptracehttp.WithEndpoint(ts.config.OTLPEndpoint),
			otlptracehttp.WithInsecure(), // Use HTTPS in production
		)
		if err != nil {
			return fmt.Errorf("failed to create OTLP exporter: %w", err)
		}
		ts.logger.Info("Using OTLP exporter", zap.String("endpoint", ts.config.OTLPEndpoint))
	} else {
		return fmt.Errorf("no tracing endpoint configured")
	}

	// Create trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(ts.config.SamplingRate)),
		sdktrace.WithBatcher(exporter,
			sdktrace.WithBatchTimeout(ts.config.BatchTimeout),
			sdktrace.WithMaxExportBatchSize(ts.config.BatchSize),
		),
	)

	// Set global tracer provider
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	ts.tracerProvider = tp
	ts.tracer = tp.Tracer(
		ts.config.ServiceName,
		trace.WithInstrumentationVersion(ts.config.ServiceVersion),
		trace.WithSchemaURL(semconv.SchemaURL),
	)

	return nil
}

// initMetrics sets up metrics collection
func (ts *TelemetryService) initMetrics(res *resource.Resource) error {
	// Create Prometheus exporter
	promExporter, err := prometheus.New()
	if err != nil {
		return fmt.Errorf("failed to create Prometheus exporter: %w", err)
	}

	// Create meter provider
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(promExporter),
	)

	otel.SetMeterProvider(mp)

	ts.meterProvider = mp
	ts.meter = mp.Meter(
		ts.config.ServiceName,
		metric.WithInstrumentationVersion(ts.config.ServiceVersion),
		metric.WithSchemaURL(semconv.SchemaURL),
	)

	return nil
}

// initBusinessMetrics creates business-specific metrics
func (ts *TelemetryService) initBusinessMetrics() error {
	var err error

	// HTTP request metrics
	ts.requestCounter, err = ts.meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return fmt.Errorf("failed to create request counter: %w", err)
	}

	ts.requestDuration, err = ts.meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10),
	)
	if err != nil {
		return fmt.Errorf("failed to create request duration histogram: %w", err)
	}

	ts.activeConnections, err = ts.meter.Int64UpDownCounter(
		"http_active_connections",
		metric.WithDescription("Number of active HTTP connections"),
		metric.WithUnit("{connection}"),
	)
	if err != nil {
		return fmt.Errorf("failed to create active connections counter: %w", err)
	}

	ts.errorCounter, err = ts.meter.Int64Counter(
		"application_errors_total",
		metric.WithDescription("Total number of application errors"),
		metric.WithUnit("{error}"),
	)
	if err != nil {
		return fmt.Errorf("failed to create error counter: %w", err)
	}

	// Initialize task metrics
	ts.taskMetrics, err = ts.initTaskMetrics()
	if err != nil {
		return fmt.Errorf("failed to initialize task metrics: %w", err)
	}

	return nil
}

// initTaskMetrics creates task-specific metrics
func (ts *TelemetryService) initTaskMetrics() (*TaskMetrics, error) {
	taskMetrics := &TaskMetrics{}
	var err error

	taskMetrics.taskCreated, err = ts.meter.Int64Counter(
		"tasks_created_total",
		metric.WithDescription("Total number of tasks created"),
		metric.WithUnit("{task}"),
	)
	if err != nil {
		return nil, err
	}

	taskMetrics.taskCompleted, err = ts.meter.Int64Counter(
		"tasks_completed_total",
		metric.WithDescription("Total number of tasks completed"),
		metric.WithUnit("{task}"),
	)
	if err != nil {
		return nil, err
	}

	taskMetrics.taskFailed, err = ts.meter.Int64Counter(
		"tasks_failed_total",
		metric.WithDescription("Total number of tasks failed"),
		metric.WithUnit("{task}"),
	)
	if err != nil {
		return nil, err
	}

	taskMetrics.taskDuration, err = ts.meter.Float64Histogram(
		"task_duration_seconds",
		metric.WithDescription("Task processing duration in seconds"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.1, 0.5, 1, 2, 5, 10, 30, 60, 300),
	)
	if err != nil {
		return nil, err
	}

	// Observable gauges for current state
	taskMetrics.tasksByStatus, err = ts.meter.Int64ObservableGauge(
		"tasks_by_status",
		metric.WithDescription("Number of tasks by status"),
		metric.WithUnit("{task}"),
	)
	if err != nil {
		return nil, err
	}

	taskMetrics.tasksByPriority, err = ts.meter.Int64ObservableGauge(
		"tasks_by_priority",
		metric.WithDescription("Number of tasks by priority"),
		metric.WithUnit("{task}"),
	)
	if err != nil {
		return nil, err
	}

	return taskMetrics, nil
}

// initSystemMetrics creates system-level metrics
func (ts *TelemetryService) initSystemMetrics() error {
	var err error

	ts.cpuUsage, err = ts.meter.Float64ObservableGauge(
		"system_cpu_usage_percent",
		metric.WithDescription("CPU usage percentage"),
		metric.WithUnit("%"),
	)
	if err != nil {
		return err
	}

	ts.memoryUsage, err = ts.meter.Float64ObservableGauge(
		"system_memory_usage_bytes",
		metric.WithDescription("Memory usage in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return err
	}

	ts.goroutines, err = ts.meter.Int64ObservableGauge(
		"go_goroutines",
		metric.WithDescription("Number of goroutines"),
		metric.WithUnit("{goroutine}"),
	)
	if err != nil {
		return err
	}

	// Register callbacks for system metrics
	_, err = ts.meter.RegisterCallback(
		ts.collectSystemMetrics,
		ts.cpuUsage,
		ts.memoryUsage,
		ts.goroutines,
	)
	if err != nil {
		return fmt.Errorf("failed to register system metrics callback: %w", err)
	}

	return nil
}

// Tracer returns the configured tracer
func (ts *TelemetryService) Tracer() trace.Tracer {
	if ts.tracer == nil {
		return otel.Tracer("noop")
	}
	return ts.tracer
}

// Meter returns the configured meter
func (ts *TelemetryService) Meter() metric.Meter {
	if ts.meter == nil {
		return otel.Meter("noop")
	}
	return ts.meter
}

// StartSpan starts a new trace span
func (ts *TelemetryService) StartSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if !ts.config.Enabled {
		return ctx, trace.SpanFromContext(ctx)
	}
	return ts.tracer.Start(ctx, spanName, opts...)
}

// RecordHTTPRequest records HTTP request metrics
func (ts *TelemetryService) RecordHTTPRequest(method, path, status string, duration time.Duration) {
	if !ts.config.Enabled {
		return
	}

	attrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.path", path),
		attribute.String("http.status", status),
	}

	ts.requestCounter.Add(context.Background(), 1, metric.WithAttributes(attrs...))
	ts.requestDuration.Record(context.Background(), duration.Seconds(), metric.WithAttributes(attrs...))
}

// RecordError records application errors
func (ts *TelemetryService) RecordError(errorType, component string) {
	if !ts.config.Enabled {
		return
	}

	attrs := []attribute.KeyValue{
		attribute.String("error.type", errorType),
		attribute.String("component", component),
	}

	ts.errorCounter.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// RecordTaskCreated records task creation metrics
func (ts *TelemetryService) RecordTaskCreated(priority string) {
	if !ts.config.Enabled || ts.taskMetrics == nil {
		return
	}

	attrs := []attribute.KeyValue{
		attribute.String("task.priority", priority),
	}

	ts.taskMetrics.taskCreated.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// RecordTaskCompleted records task completion metrics
func (ts *TelemetryService) RecordTaskCompleted(priority string, duration time.Duration) {
	if !ts.config.Enabled || ts.taskMetrics == nil {
		return
	}

	attrs := []attribute.KeyValue{
		attribute.String("task.priority", priority),
	}

	ts.taskMetrics.taskCompleted.Add(context.Background(), 1, metric.WithAttributes(attrs...))
	ts.taskMetrics.taskDuration.Record(context.Background(), duration.Seconds(), metric.WithAttributes(attrs...))
}

// RecordTaskFailed records task failure metrics
func (ts *TelemetryService) RecordTaskFailed(priority string, reason string) {
	if !ts.config.Enabled || ts.taskMetrics == nil {
		return
	}

	attrs := []attribute.KeyValue{
		attribute.String("task.priority", priority),
		attribute.String("failure.reason", reason),
	}

	ts.taskMetrics.taskFailed.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// IncrementActiveConnections increments active connections counter
func (ts *TelemetryService) IncrementActiveConnections() {
	if !ts.config.Enabled {
		return
	}
	ts.activeConnections.Add(context.Background(), 1)
}

// DecrementActiveConnections decrements active connections counter
func (ts *TelemetryService) DecrementActiveConnections() {
	if !ts.config.Enabled {
		return
	}
	ts.activeConnections.Add(context.Background(), -1)
}

// collectSystemMetrics collects system-level metrics
func (ts *TelemetryService) collectSystemMetrics(ctx context.Context, observer metric.Observer) error {
	// Collect system metrics (simplified implementation)
	// In production, use proper system metric collection libraries
	
	// CPU usage (mock implementation)
	observer.ObserveFloat64(ts.cpuUsage, 0.0) // Would collect actual CPU usage
	
	// Memory usage (mock implementation)  
	observer.ObserveFloat64(ts.memoryUsage, 0.0) // Would collect actual memory usage
	
	// Goroutines
	observer.ObserveInt64(ts.goroutines, int64(runtime.NumGoroutine()))
	
	return nil
}

// Shutdown gracefully shuts down the telemetry service
func (ts *TelemetryService) Shutdown(ctx context.Context) error {
	if !ts.config.Enabled {
		return nil
	}

	var err error

	// Shutdown tracer provider
	if tp, ok := ts.tracerProvider.(*sdktrace.TracerProvider); ok {
		if shutdownErr := tp.Shutdown(ctx); shutdownErr != nil {
			err = fmt.Errorf("failed to shutdown tracer provider: %w", shutdownErr)
		}
	}

	// Shutdown meter provider
	if mp, ok := ts.meterProvider.(*sdkmetric.MeterProvider); ok {
		if shutdownErr := mp.Shutdown(ctx); shutdownErr != nil {
			if err != nil {
				err = fmt.Errorf("%w; failed to shutdown meter provider: %w", err, shutdownErr)
			} else {
				err = fmt.Errorf("failed to shutdown meter provider: %w", shutdownErr)
			}
		}
	}

	ts.logger.Info("Telemetry service shutdown complete")
	return err
}

// generateInstanceID generates a unique instance identifier
func generateInstanceID() string {
	// In production, this could be based on hostname, pod name, etc.
	return fmt.Sprintf("instance-%d", time.Now().UnixNano())
}

// Helper function to import runtime for goroutines metric
import "runtime"