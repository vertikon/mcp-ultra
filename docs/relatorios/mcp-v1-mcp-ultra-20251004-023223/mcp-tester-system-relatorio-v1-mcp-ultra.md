# MCP Tester System — Relatório v1-mcp-ultra (20251004-023223)

## ✅ Resumo
- **Coverage (total):** total:	(statements)	0.0%
- **Validador (sem servidor):** 
- **Validador (com servidor):** 

## 🔍 Detalhes
### Endpoints de Saúde — resultados
\\\

[OK] /health -> 200
[OK] /healthz -> 200
[OK] /health/live -> 200
[OK] /health/ready -> 200
[OK] /livez -> 200
[OK] /readyz -> 200
[OK] /ping -> 200
[OK] /metrics -> 200

\\\

### Validador v4 — sem servidor
\\\
GetFileAttributesEx enhanced_validator_v4.go: The system cannot find the file specified.

\\\

### Validador v4 — com servidor
\\\
GetFileAttributesEx enhanced_validator_v4.go: The system cannot find the file specified.

\\\

### Coverage (go tool cover -func)
\\\
total:	(statements)	0.0%

\\\

### Build & Test (resumo)
\\\
go: downloading github.com/go-chi/chi/v5 v5.1.0
go: downloading github.com/go-chi/cors v1.2.1
go: downloading github.com/go-redis/redis/v8 v8.11.5
go: downloading github.com/prometheus/client_golang v1.23.0
go: downloading go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.63.0
go: downloading go.opentelemetry.io/contrib/instrumentation/runtime v0.63.0
go: downloading go.opentelemetry.io/otel/exporters/jaeger v1.17.0
go: downloading go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.38.0
go: downloading go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.38.0
go: downloading go.opentelemetry.io/otel/exporters/prometheus v0.60.0
go: downloading go.opentelemetry.io/otel/sdk v1.38.0
go: downloading go.opentelemetry.io/otel/sdk/metric v1.38.0
go: downloading github.com/nats-io/nats.go v1.37.0
go: downloading github.com/gorilla/mux v1.8.1
go: downloading github.com/golang-jwt/jwt/v5 v5.2.1
go: downloading github.com/hashicorp/vault/api v1.21.0
go: downloading github.com/felixge/httpsnoop v1.0.4
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/prometheus/client_model v0.6.2
go: downloading github.com/prometheus/common v0.65.0
go: downloading golang.org/x/sys v0.36.0
go: downloading google.golang.org/protobuf v1.36.8
go: downloading go.opentelemetry.io/proto/otlp v1.7.1
go: downloading github.com/lib/pq v1.10.9
go: downloading github.com/redis/go-redis/v9 v9.7.3
go: downloading go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.38.0
go: downloading github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
go: downloading google.golang.org/grpc v1.75.1
go: downloading github.com/cenkalti/backoff/v5 v5.0.3
go: downloading github.com/prometheus/otlptranslator v0.0.2
go: downloading github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
go: downloading github.com/cenkalti/backoff/v4 v4.3.0
go: downloading github.com/go-jose/go-jose/v4 v4.1.1
go: downloading github.com/hashicorp/errwrap v1.1.0
go: downloading github.com/hashicorp/go-cleanhttp v0.5.2
go: downloading github.com/hashicorp/go-retryablehttp v0.7.8
go: downloading github.com/hashicorp/go-multierror v1.1.1
go: downloading github.com/hashicorp/go-rootcerts v1.0.2
go: downloading github.com/hashicorp/go-secure-stdlib/parseutil v0.2.0
go: downloading github.com/hashicorp/go-secure-stdlib/strutil v0.1.2
go: downloading github.com/hashicorp/hcl v1.0.1-vault-7
go: downloading github.com/mitchellh/mapstructure v1.5.0
go: downloading golang.org/x/net v0.43.0
go: downloading golang.org/x/time v0.12.0
go: downloading github.com/nats-io/nkeys v0.4.7
go: downloading github.com/grafana/regexp v0.0.0-20240518133315-a468a5bfb3bc
go: downloading google.golang.org/genproto/googleapis/api v0.0.0-20250825161204-c5933d9347a5
go: downloading google.golang.org/genproto/googleapis/rpc v0.0.0-20250825161204-c5933d9347a5
go: downloading github.com/ryanuber/go-glob v1.0.0
go: downloading github.com/hashicorp/go-sockaddr v1.0.7
go: downloading golang.org/x/text v0.28.0
# ./tests/integration
stat E:\vertikon\business\SaaS\templates\mcp-ultra\tests\integration: directory not found
# ./tests/smoke
stat E:\vertikon\business\SaaS\templates\mcp-ultra\tests\smoke: directory not found
FAIL	./tests/integration [setup failed]
FAIL	./tests/smoke [setup failed]
# github.com/vertikon/mcp-ultra/internal/handlers
# internal/coverage/rtcov
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goexperiment
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goarch
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/byteorder
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goos
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/unsafeheader
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/asan
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/godebugs
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/profilerecord
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/msan
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/trace/tracev2
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# math/bits
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/cpu
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/itoa
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# structs
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/syscall/windows/sysdll
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode/utf8
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# sync/atomic
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# cmp
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode/utf16
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/nettrace
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# log/internal
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage/calloc
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage/uleb128
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# encoding
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# github.com/vertikon/mcp-ultra/internal/observability
internal\observability\integration.go:48:3: unknown field MetricsInterval in struct literal of type TelemetryConfig
internal\observability\integration.go:49:3: unknown field HistogramBuckets in struct literal of type TelemetryConfig
internal\observability\integration.go:52:3: unknown field JaegerEnabled in struct literal of type TelemetryConfig
internal\observability\integration.go:54:3: unknown field JaegerUser in struct literal of type TelemetryConfig
internal\observability\integration.go:55:3: unknown field JaegerPassword in struct literal of type TelemetryConfig
internal\observability\integration.go:57:3: unknown field OTLPEnabled in struct literal of type TelemetryConfig
internal\observability\integration.go:59:3: unknown field OTLPInsecure in struct literal of type TelemetryConfig
internal\observability\integration.go:60:3: unknown field OTLPHeaders in struct literal of type TelemetryConfig
internal\observability\integration.go:62:3: unknown field ConsoleEnabled in struct literal of type TelemetryConfig
internal\observability\integration.go:85:24: s.telemetry.Start undefined (type *TelemetryService has no field or method Start)
internal\observability\integration.go:85:24: too many errors
FAIL	github.com/vertikon/mcp-ultra [build failed]

\\\

Logs do servidor:
- STDOUT: E:\vertikon\business\SaaS\templates\mcp-ultra\docs\relatorios\mcp-v1-mcp-ultra-20251004-023223\server.log
- STDERR: E:\vertikon\business\SaaS\templates\mcp-ultra\docs\relatorios\mcp-v1-mcp-ultra-20251004-023223\server.err
