import os

base = r"E:ertikonusiness\SaaS	emplates\mcp-ultra	emplatesi\go"

ai_code = """package ai

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/vertikon/mcp-ultra/templates/ai/go/adapters"
	"github.com/vertikon/mcp-ultra/templates/ai/go/budgets"
	"github.com/vertikon/mcp-ultra/templates/ai/go/events"
	"github.com/vertikon/mcp-ultra/templates/ai/go/policies"
	"github.com/vertikon/mcp-ultra/templates/ai/go/router"
	"github.com/vertikon/mcp-ultra/templates/ai/go/telemetry"
	"github.com/vertikon/mcp-ultra/templates/ai/go/types"
)

type AIService struct {
	enabled       bool
	router        *router.Router
	policyEngine  *policies.PolicyEngine
	budgetTracker *budgets.BudgetTracker
	eventsPublisher *events.NATSPublisher
	adapters      map[types.Provider]adapters.AIAdapter
}

type Config struct {
	EnableAI          bool
	CanaryPercent     int
	RouterConfigPath  string
	PolicyConfigPath  string
	GuardrailsPath    string
	BudgetConfigPath  string
	NATSURL           string
	NATSStream        string
	NATSSubjectPrefix string
	ProviderPrimary   string
}

func NewAIService(cfg Config) (*AIService, error) {
	if !cfg.EnableAI {
		return &AIService{enabled: false}, nil
	}

	r, err := router.NewRouter(cfg.RouterConfigPath, cfg.CanaryPercent)
	if err != nil {
		return nil, fmt.Errorf("failed to create router: %w", err)
	}

	p, err := policies.NewPolicyEngine(cfg.PolicyConfigPath, cfg.GuardrailsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create policy engine: %w", err)
	}

	b, err := budgets.NewBudgetTracker(cfg.BudgetConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create budget tracker: %w", err)
	}

	np, err := events.NewNATSPublisher(cfg.NATSURL, cfg.NATSStream, cfg.NATSSubjectPrefix)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS publisher: %w", err)
	}

	adapterMap := make(map[types.Provider]adapters.AIAdapter)
	adapterMap[types.ProviderLocal] = adapters.NewLocalAdapter()

	return &AIService{
		enabled:         true,
		router:          r,
		policyEngine:    p,
		budgetTracker:   b,
		eventsPublisher: np,
		adapters:        adapterMap,
	}, nil
}

func (s *AIService) Infer(ctx context.Context, req types.InferenceRequest) (types.InferenceResponse, error) {
	if !s.enabled {
		return types.InferenceResponse{}, fmt.Errorf("AI service is disabled")
	}

	start := time.Now()

	allowed, block, err := s.policyEngine.CheckPre(ctx, req)
	if err != nil {
		return types.InferenceResponse{}, fmt.Errorf("policy pre-check failed: %w", err)
	}
	if !allowed {
		s.eventsPublisher.PublishPolicyBlock(ctx, *block)
		telemetry.AIPolicyBlocksTotal.WithLabelValues(
			req.Context.MCPID,
			req.Context.SDKName,
			block.Rule,
			block.Severity,
		).Inc()
		return types.InferenceResponse{}, fmt.Errorf("blocked by policy: %s", block.Rule)
	}

	decision, err := s.router.Route(ctx, req)
	if err != nil {
		return types.InferenceResponse{}, fmt.Errorf("routing failed: %w", err)
	}

	s.eventsPublisher.PublishRouterDecision(ctx, decision)
	telemetry.AIRouterDecisionsTotal.WithLabelValues(
		req.Context.MCPID,
		req.Context.SDKName,
		string(decision.Provider),
		decision.Model,
		decision.Reason,
	).Inc()

	if decision.Reason == "canary_skip" {
		return types.InferenceResponse{Content: "[CANARY SKIP]"}, nil
	}

	allowed, breach, err := s.budgetTracker.CheckBudget(ctx, req, 0.01)
	if err != nil {
		return types.InferenceResponse{}, fmt.Errorf("budget check failed: %w", err)
	}
	if !allowed {
		telemetry.AIBudgetBreachesTotal.WithLabelValues(breach.Scope).Inc()
		if breach.Action == "block" {
			return types.InferenceResponse{}, fmt.Errorf("budget exceeded: %s", breach.Scope)
		}
	}

	adapter, ok := s.adapters[decision.Provider]
	if !ok {
		s.eventsPublisher.PublishInferenceError(ctx, req, decision.Provider, decision.Model, "NO_ADAPTER", "Adapter not found")
		return types.InferenceResponse{}, fmt.Errorf("no adapter for provider: %s", decision.Provider)
	}

	resp, err := adapter.Infer(ctx, req)
	if err != nil {
		s.eventsPublisher.PublishInferenceError(ctx, req, decision.Provider, decision.Model, "INFERENCE_FAILED", err.Error())
		return types.InferenceResponse{}, fmt.Errorf("inference failed: %w", err)
	}

	allowed, block, err = s.policyEngine.CheckPost(ctx, resp)
	if err != nil {
		return types.InferenceResponse{}, fmt.Errorf("policy post-check failed: %w", err)
	}
	if !allowed {
		s.eventsPublisher.PublishPolicyBlock(ctx, *block)
		telemetry.AIPolicyBlocksTotal.WithLabelValues(
			req.Context.MCPID,
			req.Context.SDKName,
			block.Rule,
			block.Severity,
		).Inc()
		return types.InferenceResponse{}, fmt.Errorf("blocked by post-policy: %s", block.Rule)
	}

	resp.LatencyMS = time.Since(start).Milliseconds()
	s.budgetTracker.RecordUsage(ctx, req, resp.CostBRL)
	s.eventsPublisher.PublishInferenceSummary(ctx, req, resp)

	telemetry.AIInferenceRequestsTotal.WithLabelValues(
		req.Context.MCPID, req.Context.SDKName, req.Context.TenantID,
		string(resp.Provider), resp.Model,
	).Inc()

	telemetry.AIInferenceLatency.WithLabelValues(
		req.Context.MCPID, req.Context.SDKName, req.Context.TenantID,
		string(resp.Provider), resp.Model,
	).Observe(float64(resp.LatencyMS))

	telemetry.AITokensInTotal.WithLabelValues(
		req.Context.MCPID, req.Context.SDKName, req.Context.TenantID,
	).Add(float64(resp.TokensIn))

	telemetry.AITokensOutTotal.WithLabelValues(
		req.Context.MCPID, req.Context.SDKName, req.Context.TenantID,
	).Add(float64(resp.TokensOut))

	telemetry.AICostBRLTotal.WithLabelValues(
		req.Context.MCPID, req.Context.SDKName, req.Context.TenantID,
	).Add(resp.CostBRL)

	return resp, nil
}

func (s *AIService) Close() error {
	if s.enabled && s.eventsPublisher != nil {
		return s.eventsPublisher.Close()
	}
	return nil
}

func LoadConfigFromEnv() Config {
	return Config{
		EnableAI:          getEnvBool("ENABLE_AI", false),
		CanaryPercent:     getEnvInt("AI_CANARY_PERCENT", 0),
		RouterConfigPath:  getEnv("AI_ROUTER_CONFIG", "./config/ai-router.rules.json"),
		PolicyConfigPath:  getEnv("AI_POLICY_CONFIG", "./config/ai-policies.yaml"),
		GuardrailsPath:    getEnv("AI_GUARDRAILS_CONFIG", "./config/ai-guardrails.json"),
		BudgetConfigPath:  getEnv("AI_BUDGET_CONFIG", "./config/ai-budgets.json"),
		NATSURL:           getEnv("NATS_URL", "nats://localhost:4222"),
		NATSStream:        getEnv("NATS_STREAM", "ultra_ai_events"),
		NATSSubjectPrefix: getEnv("NATS_SUBJECT_PREFIX", "ultra.ai"),
		ProviderPrimary:   getEnv("PROVIDER_PRIMARY", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return fallback
}
"""

with open(os.path.join(base, "ai.go"), "w", encoding="utf-8") as f:
    f.write(ai_code)

print("OK: ai.go created")
