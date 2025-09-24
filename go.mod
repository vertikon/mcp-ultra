module github.com/vertikon/mcp-ultra

go 1.22

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/cors v1.2.1
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/lib/pq v1.10.9
	github.com/redis/go-redis/v9 v9.5.1
	github.com/nats-io/nats.go v1.31.0
	github.com/prometheus/client_golang v1.19.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/google/uuid v1.6.0
	go.uber.org/zap v1.27.0
	
	// OpenTelemetry Core
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/trace v1.28.0
	go.opentelemetry.io/otel/metric v1.28.0
	go.opentelemetry.io/otel/sdk v1.28.0
	go.opentelemetry.io/otel/sdk/metric v1.28.0
	
	// OpenTelemetry Exporters
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.28.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.28.0
	go.opentelemetry.io/otel/exporters/jaeger v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.28.0
	go.opentelemetry.io/otel/exporters/prometheus v0.50.0
	
	// OpenTelemetry Instrumentation
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0
	go.opentelemetry.io/contrib/instrumentation/github.com/lib/pq/otelpq v0.53.0
	go.opentelemetry.io/contrib/instrumentation/github.com/go-redis/redis/v8/otelredis v0.53.0
	
	// Other dependencies
	github.com/cespare/xxhash/v2 v2.2.0
	github.com/golang-migrate/migrate/v4 v4.17.0
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1
)

