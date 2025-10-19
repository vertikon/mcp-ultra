package wiring

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/vertikon/mcp-ultra/internal/ai/router"
	"github.com/vertikon/mcp-ultra/internal/ai/telemetry"
	"github.com/vertikon/mcp-ultra/pkg/metrics"
)

type Config struct {
	BasePathAI string // path to templates/ai
	Registry   metrics.Registerer
}

// Service holds minimal IA singletons (router + telemetry enabled flag).
type Service struct {
	Router  *router.Router
	Enabled bool
}

func Init(ctx context.Context, cfg Config) (*Service, error) {
	// feature flags are inside feature_flags.json at BasePathAI
	base := cfg.BasePathAI
	if base == "" {
		cwd, _ := os.Getwd()
		base = filepath.Join(cwd, "templates", "ai")
	}

	r, _ := router.Load(base)
	telemetry.Init(cfg.Registry)

	svc := &Service{Router: r, Enabled: r != nil && r.Enabled()}
	_ = ctx // reserved for future async init
	time.AfterFunc(0, func() { /* noop */ })
	return svc, nil
}
