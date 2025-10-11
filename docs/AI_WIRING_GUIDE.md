# AI Wiring Mínimo — Guia de Integração

## Visão Geral

Este guia documenta como integrar a camada de IA opt-in no MCP-Ultra com telemetria Prometheus, roteamento de providers e eventos NATS.

## Componentes

### 1. Telemetria (`internal/ai/telemetry/metrics.go`)

Métricas Prometheus para observabilidade de IA:

- `ai_inference_requests_total` - Total de requisições de inferência
- `ai_inference_latency_ms` - Latência de inferência (histograma)
- `ai_tokens_in_total` - Total de tokens de entrada
- `ai_tokens_out_total` - Total de tokens de saída
- `ai_cost_brl_total` - Custo acumulado em BRL
- `ai_policy_blocks_total` - Total de bloqueios por política
- `ai_router_decisions_total` - Total de decisões de roteamento
- `ai_budget_breaches_total` - Violações de orçamento

### 2. Roteador (`internal/ai/router/router.go`)

Carrega configurações de:
- `templates/ai/feature_flags.json` - flags de ativação
- `templates/ai/config/ai-router.rules.json` - regras de roteamento

### 3. Eventos NATS (`internal/ai/events/handlers.go`)

Publicadores para eventos IA:
- `ultra.ai.router.decision` - Decisões de roteamento
- `ultra.ai.policy.block` - Bloqueios de política
- `ultra.ai.inference.error` - Erros de inferência
- `ultra.ai.inference.summary` - Resumo de inferência

### 4. Wiring (`internal/ai/wiring/wiring.go`)

Inicialização centralizada dos componentes de IA.

## Integração no Main

### 1. Inicialização

```go
import (
	"context"
	"time"

	"github.com/vertikon/mcp-ultra/internal/ai/wiring"
	"github.com/vertikon/mcp-ultra/internal/ai/telemetry"
	"github.com/vertikon/mcp-ultra/internal/ai/events"
	natsx "github.com/vertikon/mcp-ultra/internal/nats"
)

func initAI(ctx context.Context, natsPublisher *natsx.Publisher) (*wiring.Service, error) {
	svc, err := wiring.Init(ctx, wiring.Config{
		BasePathAI: "templates/ai",
		Registry:   nil, // usa prometheus.DefaultRegisterer
	})
	if err != nil {
		return nil, err
	}

	if !svc.Enabled {
		log.Info("AI layer is disabled (ENABLE_AI=false)")
		return svc, nil
	}

	log.Info("AI layer initialized successfully")
	return svc, nil
}
```

### 2. Uso do Roteador

```go
func handleInferenceRequest(ctx context.Context, aiSvc *wiring.Service, useCase string) error {
	if !aiSvc.Enabled {
		return errors.New("AI disabled")
	}

	// Obter decisão de roteamento
	dec, err := aiSvc.Router.Decide(useCase)
	if err != nil {
		return fmt.Errorf("routing decision failed: %w", err)
	}

	// Registrar decisão nas métricas
	telemetry.IncRouterDecision(telemetry.Labels{
		TenantID: "default",
		MCPID:    "mcp-ultra",
		SDKName:  "sdk-xyz",
		Provider: dec.Provider,
		Model:    dec.Model,
		Reason:   dec.Reason,
	})

	// Publicar evento de decisão
	_ = events.PublishRouterDecision(ctx, natsPublisher, "ultra.ai.router.decision", events.RouterDecision{
		Base:     events.Base{TenantID: "default", MCPID: "mcp-ultra", SDKName: "sdk-xyz"},
		UseCase:  useCase,
		Provider: dec.Provider,
		Model:    dec.Model,
		Reason:   dec.Reason,
	})

	return nil
}
```

### 3. Telemetria de Inferência

```go
func performInference(ctx context.Context, provider, model, prompt string) error {
	start := telemetry.ObserveStart()

	// Executar inferência no provider
	result, err := callProvider(provider, model, prompt)
	if err != nil {
		// Publicar erro
		_ = events.PublishInferenceError(ctx, natsPublisher, "ultra.ai.inference.error", events.InferenceError{
			Base:     events.Base{TenantID: "default", MCPID: "mcp-ultra"},
			Provider: provider,
			Model:    model,
			Code:     "PROVIDER_ERROR",
			Message:  err.Error(),
		})
		return err
	}

	// Registrar telemetria
	telemetry.ObserveInference(telemetry.InferenceMeta{
		Labels: telemetry.Labels{
			TenantID: "default",
			MCPID:    "mcp-ultra",
			SDKName:  "sdk-xyz",
			Provider: provider,
			Model:    model,
			UseCase:  "generation",
		},
		TokensIn:  result.TokensIn,
		TokensOut: result.TokensOut,
		CostBRL:   result.Cost,
		Start:     start,
		End:       time.Now(),
	})

	// Publicar resumo
	_ = events.PublishInferenceSummary(ctx, natsPublisher, "ultra.ai.inference.summary", events.InferenceSummary{
		Base:      events.Base{TenantID: "default", MCPID: "mcp-ultra"},
		UseCase:   "generation",
		TokensIn:  result.TokensIn,
		TokensOut: result.TokensOut,
		LatencyMs: int(time.Since(start).Milliseconds()),
		CostBRL:   result.Cost,
		Cached:    false,
	})

	return nil
}
```

### 4. Bloqueios de Política

```go
func checkPolicy(ctx context.Context, content string) error {
	if containsPII(content) {
		// Incrementar métrica
		telemetry.IncPolicyBlock(telemetry.Labels{
			TenantID: "default",
			MCPID:    "mcp-ultra",
			SDKName:  "sdk-xyz",
			Rule:     "pii_check",
			Severity: "high",
		})

		// Publicar evento
		_ = events.PublishPolicyBlock(ctx, natsPublisher, "ultra.ai.policy.block", events.PolicyBlock{
			Base:     events.Base{TenantID: "default", MCPID: "mcp-ultra"},
			Rule:     "pii_check",
			Severity: "high",
			Sample:   content[:50], // primeiros 50 chars
		})

		return errors.New("content blocked by PII policy")
	}
	return nil
}
```

## Configuração

### feature_flags.json

```json
{
  "ai": {
    "enabled": false,
    "mode": "balanced",
    "canary_percent": 0,
    "router": "rules",
    "budgets": {
      "enforce": true,
      "hard_stop_on_breach": true
    },
    "guardrails": {
      "pre": true,
      "post": true
    },
    "telemetry": {
      "prometheus": true,
      "otel": true
    }
  }
}
```

### .env

```env
ENABLE_AI=false
AI_MODE=balanced
AI_CANARY_PERCENT=0
PROVIDER_PRIMARY=local
```

## DRY-RUN (Teste sem Custo)

Para testar a integração sem custos:

1. Mantenha `ENABLE_AI=false` ou use `PROVIDER_PRIMARY=local`
2. Configure `AI_CANARY_PERCENT=0`
3. Verifique métricas em `/metrics`
4. Monitore eventos NATS em `ultra.ai.*`

## Boas Práticas

1. **Inicialização Lazy**: Só inicialize IA quando `ENABLE_AI=true`
2. **Fail-Safe**: Se o roteador falhar, use fallback local
3. **Observabilidade First**: Sempre registre métricas antes de publicar eventos
4. **Secrets Management**: Nunca comite API keys; use Secret Manager
5. **Budget Enforcement**: Configure limites e monitore `ai_budget_breaches_total`

## Troubleshooting

### Métricas não aparecem em /metrics

- Verifique se `telemetry.Init()` foi chamado
- Confirme que `ENABLE_AI=true` em feature_flags.json

### Eventos NATS não são publicados

- Verifique se o publisher está conectado ao NATS
- Confirme que os subjects estão corretos (`ultra.ai.*`)

### Decisão de roteamento falha

- Verifique se `ai-router.rules.json` existe e é válido
- Confirme que o use-case existe nas regras default

## Próximos Passos

1. Implementar adaptadores de providers (OpenAI, Qwen, Local)
2. Adicionar políticas de guardrails (PII, profanity)
3. Implementar controle de budget
4. Criar dashboard Grafana para métricas AI
