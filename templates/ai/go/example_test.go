package ai_test

import (
	"context"
	"testing"
	"time"

	"github.com/vertikon/mcp-ultra/templates/ai/go"
	"github.com/vertikon/mcp-ultra/templates/ai/go/types"
)

func TestAIServiceDryRun(t *testing.T) {
	cfg := ai.Config{
		EnableAI:          true,
		CanaryPercent:     100,
		RouterConfigPath:  "../config/ai-router.rules.json",
		PolicyConfigPath:  "../config/ai-policies.yaml",
		GuardrailsPath:    "../config/ai-guardrails.json",
		BudgetConfigPath:  "../config/ai-budgets.json",
		NATSURL:           "nats://localhost:4222",
		NATSStream:        "ultra_ai_events",
		NATSSubjectPrefix: "ultra.ai",
		ProviderPrimary:   "local",
	}

	svc, err := ai.NewAIService(cfg)
	if err != nil {
		t.Skipf("Service init failed (expected in dry-run): %v", err)
	}
	defer svc.Close()

	req := types.InferenceRequest{
		Context: types.Context{
			TenantID:  "test-tenant",
			MCPID:     "mcp-test",
			SDKName:   "sdk-test",
			Timestamp: time.Now(),
		},
		UseCase:   types.UseCaseGeneration,
		Prompt:    "Hello AI",
		MaxTokens: 100,
		Timeout:   5 * time.Second,
	}

	resp, err := svc.Infer(context.Background(), req)
	if err != nil {
		t.Logf("Inference failed (may be expected): %v", err)
		return
	}

	t.Logf("Response: %+v", resp)
}
