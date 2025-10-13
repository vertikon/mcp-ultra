package router

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"{{MODULE_PATH}}/templates/ai/go/types"
)

type RouterConfig struct {
	Version   string                     	Default   map[types.UseCase]ModelDef 	Overrides []Override                 	Fallbacks []Fallback                 }

type ModelDef struct {
	Provider types.Provider 	Model    string         }

type Override struct {
	When map[string]string              	Use  map[types.UseCase]ModelDef     	Mode *types.AIMode                  }

type Fallback struct {
	From ModelDef 	To   ModelDef }

type Router struct {
	config        RouterConfig
	canaryPercent int
}

func NewRouter(configPath string, canaryPercent int) (*Router, error) {
	data, err := os.ReadFile(configPath)
	if err \!= nil {
		return nil, fmt.Errorf("failed to read router config: %w", err)
	}

	var cfg RouterConfig
	if err := json.Unmarshal(data, &cfg); err \!= nil {
		return nil, fmt.Errorf("failed to parse router config: %w", err)
	}

	return &Router{
		config:        cfg,
		canaryPercent: canaryPercent,
	}, nil
}

func (r *Router) Route(ctx context.Context, req types.InferenceRequest) (types.RouterDecision, error) {
	// Canary check
	if r.canaryPercent == 0 || rand.Intn(100) >= r.canaryPercent {
		return types.RouterDecision{
			Context:  req.Context,
			UseCase:  req.UseCase,
			Provider: "none",
			Model:    "none",
			Reason:   "canary_skip",
		}, nil
	}

	// Check overrides
	for _, override := range r.config.Overrides {
		if matches(override.When, req.Context) {
			if model, ok := override.Use[req.UseCase]; ok {
				return types.RouterDecision{
					Context:  req.Context,
					UseCase:  req.UseCase,
					Provider: model.Provider,
					Model:    model.Model,
					Reason:   "override",
				}, nil
			}
		}
	}

	// Default routing
	if model, ok := r.config.Default[req.UseCase]; ok {
		return types.RouterDecision{
			Context:  req.Context,
			UseCase:  req.UseCase,
			Provider: model.Provider,
			Model:    model.Model,
			Reason:   "default",
		}, nil
	}

	return types.RouterDecision{}, fmt.Errorf("no route found for use case: %s", req.UseCase)
}

func matches(when map[string]string, ctx types.Context) bool {
	if mcpID, ok := when["mcp_id"]; ok && mcpID \!= ctx.MCPID {
		return false
	}
	if tenantID, ok := when["tenant_id"]; ok && tenantID \!= ctx.TenantID {
		return false
	}
	return true
}
