package events

import (
\t"context"
\t"encoding/json"
\t"fmt"
\t"time"

\t"github.com/nats-io/nats.go"
\t"github.com/vertikon/mcp-ultra/templates/ai/go/types"
)

type NATSPublisher struct {
\tnc            *nats.Conn
\tjs            nats.JetStreamContext
\tstream        string
\tsubjectPrefix string
}

func NewNATSPublisher(url, stream, subjectPrefix string) (*NATSPublisher, error) {
\tnc, err := nats.Connect(url)
\tif err != nil {
\t\treturn nil, fmt.Errorf("failed to connect to NATS: %w", err)
\t}

\tjs, err := nc.JetStream()
\tif err != nil {
\t\treturn nil, fmt.Errorf("failed to create JetStream context: %w", err)
\t}

\treturn &NATSPublisher{
\t\tnc:            nc,
\t\tjs:            js,
\t\tstream:        stream,
\t\tsubjectPrefix: subjectPrefix,
\t}, nil
}

func (p *NATSPublisher) PublishRouterDecision(ctx context.Context, decision types.RouterDecision) error {
\tdecision.Timestamp = time.Now()
\tdata, err := json.Marshal(decision)
\tif err != nil {
\t\treturn fmt.Errorf("failed to marshal router decision: %w", err)
\t}

\tsubject := fmt.Sprintf("%s.router.decision", p.subjectPrefix)
\t_, err = p.js.Publish(subject, data)
\treturn err
}

func (p *NATSPublisher) PublishPolicyBlock(ctx context.Context, block types.PolicyBlock) error {
\tblock.Timestamp = time.Now()
\tdata, err := json.Marshal(block)
\tif err != nil {
\t\treturn fmt.Errorf("failed to marshal policy block: %w", err)
\t}

\tsubject := fmt.Sprintf("%s.policy.block", p.subjectPrefix)
\t_, err = p.js.Publish(subject, data)
\treturn err
}

func (p *NATSPublisher) PublishInferenceSummary(ctx context.Context, req types.InferenceRequest, resp types.InferenceResponse) error {
\tsummary := map[string]interface{}{
\t\t"timestamp":  time.Now().Format(time.RFC3339),
\t\t"tenant_id":  req.Context.TenantID,
\t\t"mcp_id":     req.Context.MCPID,
\t\t"sdk_name":   req.Context.SDKName,
\t\t"use_case":   req.UseCase,
\t\t"tokens_in":  resp.TokensIn,
\t\t"tokens_out": resp.TokensOut,
\t\t"latency_ms": resp.LatencyMS,
\t\t"cost_brl":   resp.CostBRL,
\t\t"cached":     resp.Cached,
\t}

\tdata, err := json.Marshal(summary)
\tif err != nil {
\t\treturn fmt.Errorf("failed to marshal inference summary: %w", err)
\t}

\tsubject := fmt.Sprintf("%s.inference.summary", p.subjectPrefix)
\t_, err = p.js.Publish(subject, data)
\treturn err
}

func (p *NATSPublisher) PublishInferenceError(ctx context.Context, req types.InferenceRequest, provider types.Provider, model string, errCode string, errMsg string) error {
\terrorEvent := map[string]interface{}{
\t\t"timestamp":  time.Now().Format(time.RFC3339),
\t\t"tenant_id":  req.Context.TenantID,
\t\t"mcp_id":     req.Context.MCPID,
\t\t"sdk_name":   req.Context.SDKName,
\t\t"provider":   provider,
\t\t"model":      model,
\t\t"code":       errCode,
\t\t"message":    errMsg,
\t}

\tdata, err := json.Marshal(errorEvent)
\tif err != nil {
\t\treturn fmt.Errorf("failed to marshal inference error: %w", err)
\t}

\tsubject := fmt.Sprintf("%s.inference.error", p.subjectPrefix)
\t_, err = p.js.Publish(subject, data)
\treturn err
}

func (p *NATSPublisher) Close() error {
\tp.nc.Close()
\treturn nil
}
