package adapters

import (
	"context"
	"fmt"

	"github.com/vertikon/mcp-ultra/templates/ai/go/types"
)

// AIAdapter interface para providers
type AIAdapter interface {
	Infer(ctx context.Context, req types.InferenceRequest) (types.InferenceResponse, error)
	HealthCheck(ctx context.Context) error
}

// LocalAdapter implementacao local (no-op para DRY-RUN)
type LocalAdapter struct{}

func NewLocalAdapter() *LocalAdapter {
	return &LocalAdapter{}
}

func (a *LocalAdapter) Infer(ctx context.Context, req types.InferenceRequest) (types.InferenceResponse, error) {
	return types.InferenceResponse{
		Content:   "[LOCAL NO-OP] " + req.Prompt,
		TokensIn:  len(req.Prompt) / 4,
		TokensOut: 50,
		LatencyMS: 100,
		CostBRL:   0.001,
		Cached:    false,
		Provider:  types.ProviderLocal,
		Model:     "local-dry-run",
	}, nil
}

func (a *LocalAdapter) HealthCheck(ctx context.Context) error {
	return nil
}

// OpenAIAdapter placeholder
type OpenAIAdapter struct {
	apiKey   string
	endpoint string
}

func NewOpenAIAdapter(apiKey, endpoint string) *OpenAIAdapter {
	return &OpenAIAdapter{
		apiKey:   apiKey,
		endpoint: endpoint,
	}
}

func (a *OpenAIAdapter) Infer(ctx context.Context, req types.InferenceRequest) (types.InferenceResponse, error) {
	// TODO: Implementar chamada real para OpenAI API
	return types.InferenceResponse{}, fmt.Errorf("OpenAI adapter not implemented")
}

func (a *OpenAIAdapter) HealthCheck(ctx context.Context) error {
	return fmt.Errorf("OpenAI health check not implemented")
}

// QwenAdapter placeholder
type QwenAdapter struct {
	apiKey   string
	endpoint string
}

func NewQwenAdapter(apiKey, endpoint string) *QwenAdapter {
	return &QwenAdapter{
		apiKey:   apiKey,
		endpoint: endpoint,
	}
}

func (a *QwenAdapter) Infer(ctx context.Context, req types.InferenceRequest) (types.InferenceResponse, error) {
	// TODO: Implementar chamada real para Qwen API
	return types.InferenceResponse{}, fmt.Errorf("Qwen adapter not implemented")
}

func (a *QwenAdapter) HealthCheck(ctx context.Context) error {
	return fmt.Errorf("Qwen health check not implemented")
}
