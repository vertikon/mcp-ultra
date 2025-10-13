package budgets

import (
\t"context"
\t"encoding/json"
\t"fmt"
\t"os"
\t"sync"
\t"time"

\t"{{MODULE_PATH}}/templates/ai/go/types"
)

type BudgetConfig struct {
\tVersion    string          \tGlobal     BudgetLimit     \tPerTenant  []TenantBudget  \tPerMCP     []MCPBudget     }

type BudgetLimit struct {
\tDailyBRLCap float64 \tOnBreach    string  \ // degrade|block|alert_only
}

type TenantBudget struct {
\tTenantID    string  \tDailyBRLCap float64 \tOnBreach    string  }

type MCPBudget struct {
\tMCPID       string  \tDailyBRLCap float64 \tOnBreach    string  }

type BudgetTracker struct {
\tconfig BudgetConfig
\tmu     sync.RWMutex
\tusage  map[string]float64 // key: scope:id, value: daily usage
}

func NewBudgetTracker(configPath string) (*BudgetTracker, error) {
\tdata, err := os.ReadFile(configPath)
\tif err \!= nil {
\t\treturn nil, fmt.Errorf("failed to read budget config: %w", err)
\t}

\tvar cfg BudgetConfig
\tif err := json.Unmarshal(data, &cfg); err \!= nil {
\t\treturn nil, fmt.Errorf("failed to parse budget config: %w", err)
\t}

\treturn &BudgetTracker{
\t\tconfig: cfg,
\t\tusage:  make(map[string]float64),
\t}, nil
}

func (b *BudgetTracker) CheckBudget(ctx context.Context, req types.InferenceRequest, estimatedCost float64) (bool, *types.BudgetBreach, error) {
\tb.mu.RLock()
\tdefer b.mu.RUnlock()

\t// Check global budget
\tglobalKey := "global:all"
\tglobalUsage := b.usage[globalKey]
\tif globalUsage+estimatedCost > b.config.Global.DailyBRLCap {
\t\treturn false, &types.BudgetBreach{
\t\t\tScope:    "global",
\t\t\tCapBRL:   b.config.Global.DailyBRLCap,
\t\t\tUsedBRL:  globalUsage,
\t\t\tAction:   b.config.Global.OnBreach,
\t\t}, nil
\t}

\t// Check tenant budget
\ttenantKey := fmt.Sprintf("tenant:%s", req.Context.TenantID)
\tfor _, tb := range b.config.PerTenant {
\t\tif tb.TenantID == req.Context.TenantID {
\t\t\ttenantUsage := b.usage[tenantKey]
\t\t\tif tenantUsage+estimatedCost > tb.DailyBRLCap {
\t\t\t\treturn false, &types.BudgetBreach{
\t\t\t\t\tScope:    "tenant",
\t\t\t\t\tTenantID: req.Context.TenantID,
\t\t\t\t\tCapBRL:   tb.DailyBRLCap,
\t\t\t\t\tUsedBRL:  tenantUsage,
\t\t\t\t\tAction:   tb.OnBreach,
\t\t\t\t}, nil
\t\t\t}
\t\t}
\t}

\t// Check MCP budget
\tmcpKey := fmt.Sprintf("mcp:%s", req.Context.MCPID)
\tfor _, mb := range b.config.PerMCP {
\t\tif mb.MCPID == req.Context.MCPID {
\t\t\tmcpUsage := b.usage[mcpKey]
\t\t\tif mcpUsage+estimatedCost > mb.DailyBRLCap {
\t\t\t\treturn false, &types.BudgetBreach{
\t\t\t\t\tScope:  "mcp",
\t\t\t\t\tMCPID:  req.Context.MCPID,
\t\t\t\t\tCapBRL: mb.DailyBRLCap,
\t\t\t\t\tUsedBRL: mcpUsage,
\t\t\t\t\tAction: mb.OnBreach,
\t\t\t\t}, nil
\t\t\t}
\t\t}
\t}

\treturn true, nil, nil
}

func (b *BudgetTracker) RecordUsage(ctx context.Context, req types.InferenceRequest, actualCost float64) error {
\tb.mu.Lock()
\tdefer b.mu.Unlock()

\tglobalKey := "global:all"
\tb.usage[globalKey] += actualCost

\ttenantKey := fmt.Sprintf("tenant:%s", req.Context.TenantID)
\tb.usage[tenantKey] += actualCost

\tmcpKey := fmt.Sprintf("mcp:%s", req.Context.MCPID)
\tb.usage[mcpKey] += actualCost

\treturn nil
}

func (b *BudgetTracker) ResetDaily() {
\tb.mu.Lock()
\tdefer b.mu.Unlock()
\tb.usage = make(map[string]float64)
}
