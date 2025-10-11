# AI Turbo - Go Implementation

Implementacao Go completa do sistema de IA para MCP-Ultra.

## Estrutura

- `types/` - Tipos e interfaces core
- `router/` - Roteamento de requisicoes (canary, overrides, fallbacks)
- `policies/` - Filtros pre/pos-inferencia (PII, profanity, guardrails)
- `budgets/` - Controle de orcamento (global, por tenant, por MCP)
- `adapters/` - Implementacoes de providers (Local, OpenAI, Qwen)
- `telemetry/` - Metricas Prometheus
- `events/` - Publisher NATS para eventos
- `ai.go` - Orquestrador principal (AIService)

## Uso basico

```go
cfg := ai.LoadConfigFromEnv()
svc, err := ai.NewAIService(cfg)
if err != nil {
    log.Fatal(err)
}
defer svc.Close()

req := types.InferenceRequest{
    Context: types.Context{
        TenantID: "tenant-1",
        MCPID:    "mcp-wa-notifications",
        SDKName:  "sdk-wa-notifications",
    },
    UseCase: types.UseCaseGeneration,
    Prompt:  "Classifique este texto...",
}

resp, err := svc.Infer(context.Background(), req)
if err != nil {
    log.Printf("Error: %v", err)
}
```

## DRY-RUN (sem provider real)

```bash
export ENABLE_AI=true
export AI_CANARY_PERCENT=100
export PROVIDER_PRIMARY=local
go test -v ./...
```

## Build

```bash
go mod tidy
go build ./...
```

## Dependencias

- github.com/nats-io/nats.go v1.31.0
- github.com/prometheus/client_golang v1.17.0

## Observabilidade

Metricas expostas em `/metrics`:
- ai_inference_requests_total
- ai_inference_latency_ms
- ai_tokens_in_total
- ai_tokens_out_total
- ai_cost_brl_total
- ai_policy_blocks_total
- ai_router_decisions_total
- ai_budget_breaches_total

Eventos NATS:
- ultra.ai.router.decision
- ultra.ai.policy.block
- ultra.ai.inference.summary
- ultra.ai.inference.error
