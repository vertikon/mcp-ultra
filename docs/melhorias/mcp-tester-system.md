# MCP Tester System — Relatório v1-mcp-ultra (20251004-040617)

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
# github.com/vertikon/mcp-ultra/internal/features
internal\features\analytics.go:69:15: undefined: UserContext
internal\features\advanced.go:467:95: undefined: UserContext
internal\features\advanced.go:663:83: undefined: UserContext
internal\features\advanced.go:729:90: undefined: UserContext
internal\features\advanced.go:752:83: undefined: UserContext
internal\features\advanced.go:797:91: undefined: UserContext
internal\features\advanced.go:811:86: undefined: UserContext
internal\features\advanced.go:825:87: undefined: UserContext
internal\features\advanced.go:839:78: undefined: UserContext
internal\features\advanced.go:844:83: undefined: UserContext
internal\features\advanced.go:844:83: too many errors
# github.com/vertikon/mcp-ultra/internal/lifecycle
internal\lifecycle\components.go:93:13: undefined: events.EventBus
internal\lifecycle\components.go:160:13: status.Healthy undefined (type map[string]interface{} has no field or method Healthy)
internal\lifecycle\components.go:161:72: status.Message undefined (type map[string]interface{} has no field or method Message)
internal\lifecycle\components.go:167:33: o.Service.HealthCheck().Healthy undefined (type map[string]interface{} has no field or method Healthy)
internal\lifecycle\components.go:171:33: o.Service.HealthCheck().Healthy undefined (type map[string]interface{} has no field or method Healthy)
internal\lifecycle\components.go:203:13: status.Healthy undefined (type map[string]interface{} has no field or method Healthy)
internal\lifecycle\components.go:205:15: status.Violations undefined (type map[string]interface{} has no field or method Violations)
internal\lifecycle\components.go:205:39: status.Errors undefined (type map[string]interface{} has no field or method Errors)
internal\lifecycle\components.go:220:16: status.Healthy undefined (type map[string]interface{} has no field or method Healthy)
internal\lifecycle\health.go:519:2: declared and not used: data
internal\lifecycle\health.go:519:2: too many errors
# github.com/vertikon/mcp-ultra/internal/handlers
# internal/godebugs
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/unsafeheader
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/byteorder
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goexperiment
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage/rtcov
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goarch
ok  	github.com/vertikon/mcp-ultra/tests/integration	0.525s	coverage: [no statements]
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/cpu
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/goos
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/profilerecord
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/asan
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/trace/tracev2
ok  	github.com/vertikon/mcp-ultra/tests/smoke	0.530s	coverage: [no statements]
FAIL	github.com/vertikon/mcp-ultra [build failed]
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/msan
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode/utf8
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# math/bits
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/itoa
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/syscall/windows/sysdll
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# cmp
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# structs
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# unicode/utf16
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# log/internal
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage/uleb128
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# encoding
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# sync/atomic
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/nettrace
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage/calloc
compile: version "go1.25.0" does not match go tool version "go1.25.1"
# internal/coverage
compile: version "go1.25.0" does not match go tool version "go1.25.1"

\\\

Logs do servidor:
- STDOUT: E:\vertikon\business\SaaS\templates\mcp-ultra\docs\relatorios\mcp-v1-mcp-ultra-20251004-040617\server.log
- STDERR: E:\vertikon\business\SaaS\templates\mcp-ultra\docs\relatorios\mcp-v1-mcp-ultra-20251004-040617\server.err
