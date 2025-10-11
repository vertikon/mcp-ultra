package telemetry

import (
\t"github.com/prometheus/client_golang/prometheus"
\t"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
\tAIInferenceRequestsTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_inference_requests_total",
\t\t\tHelp: "Total number of AI inference requests",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "tenant_id", "provider", "model"},
\t)

\tAIInferenceLatency = promauto.NewHistogramVec(
\t\tprometheus.HistogramOpts{
\t\t\tName:    "ai_inference_latency_ms",
\t\t\tHelp:    "AI inference latency in milliseconds",
\t\t\tBuckets: prometheus.ExponentialBuckets(10, 2, 10),
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "tenant_id", "provider", "model"},
\t)

\tAITokensInTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_tokens_in_total",
\t\t\tHelp: "Total input tokens",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "tenant_id"},
\t)

\tAITokensOutTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_tokens_out_total",
\t\t\tHelp: "Total output tokens",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "tenant_id"},
\t)

\tAICostBRLTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_cost_brl_total",
\t\t\tHelp: "Total AI cost in BRL",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "tenant_id"},
\t)

\tAIPolicyBlocksTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_policy_blocks_total",
\t\t\tHelp: "Total policy blocks",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "rule", "severity"},
\t)

\tAIRouterDecisionsTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_router_decisions_total",
\t\t\tHelp: "Total router decisions",
\t\t},
\t\t[]string{"mcp_id", "sdk_name", "provider", "model", "reason"},
\t)

\tAIBudgetBreachesTotal = promauto.NewCounterVec(
\t\tprometheus.CounterOpts{
\t\t\tName: "ai_budget_breaches_total",
\t\t\tHelp: "Total budget breaches",
\t\t},
\t\t[]string{"scope"},
\t)

\tAIBudgetRemainingBRL = promauto.NewGaugeVec(
\t\tprometheus.GaugeOpts{
\t\t\tName: "ai_budget_remaining_brl",
\t\t\tHelp: "Remaining budget in BRL",
\t\t},
\t\t[]string{"scope", "tenant_id", "mcp_id"},
\t)
)
