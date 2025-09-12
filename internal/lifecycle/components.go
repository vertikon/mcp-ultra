package lifecycle

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	
	"github.com/vertikon/mcp-ultra/internal/compliance"
	"github.com/vertikon/mcp-ultra/internal/events"
	"github.com/vertikon/mcp-ultra/internal/observability"
)

// DatabaseComponent wraps database connection for lifecycle management
type DatabaseComponent struct {
	DB *sql.DB
}

func (d *DatabaseComponent) Name() string {
	return "database"
}

func (d *DatabaseComponent) Priority() int {
	return 1 // High priority - infrastructure component
}

func (d *DatabaseComponent) Start(ctx context.Context) error {
	// Database is already connected, just verify
	return d.DB.PingContext(ctx)
}

func (d *DatabaseComponent) Stop(ctx context.Context) error {
	return d.DB.Close()
}

func (d *DatabaseComponent) HealthCheck(ctx context.Context) error {
	return d.DB.PingContext(ctx)
}

func (d *DatabaseComponent) IsReady() bool {
	return d.DB.Ping() == nil
}

func (d *DatabaseComponent) IsHealthy() bool {
	return d.DB.Ping() == nil
}

// RedisComponent wraps Redis client for lifecycle management
type RedisComponent struct {
	Client redis.Cmdable
}

func (r *RedisComponent) Name() string {
	return "redis"
}

func (r *RedisComponent) Priority() int {
	return 2 // High priority - cache infrastructure
}

func (r *RedisComponent) Start(ctx context.Context) error {
	// Redis is already connected, just verify
	return r.Client.Ping(ctx).Err()
}

func (r *RedisComponent) Stop(ctx context.Context) error {
	if client, ok := r.Client.(*redis.Client); ok {
		return client.Close()
	}
	return nil
}

func (r *RedisComponent) HealthCheck(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func (r *RedisComponent) IsReady() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.Client.Ping(ctx).Err() == nil
}

func (r *RedisComponent) IsHealthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.Client.Ping(ctx).Err() == nil
}

// EventBusComponent wraps event bus for lifecycle management
type EventBusComponent struct {
	Bus events.EventBus
}

func (e *EventBusComponent) Name() string {
	return "eventbus"
}

func (e *EventBusComponent) Priority() int {
	return 3 // Medium priority - messaging infrastructure
}

func (e *EventBusComponent) Start(ctx context.Context) error {
	// Event bus is already initialized
	return nil
}

func (e *EventBusComponent) Stop(ctx context.Context) error {
	e.Bus.Close()
	return nil
}

func (e *EventBusComponent) HealthCheck(ctx context.Context) error {
	// Simple health check - try to publish a test event
	return e.Bus.Publish("health.check", map[string]interface{}{
		"timestamp": time.Now(),
		"component": "eventbus",
	})
}

func (e *EventBusComponent) IsReady() bool {
	// Event bus is ready if we can publish events
	err := e.Bus.Publish("readiness.check", map[string]interface{}{
		"timestamp": time.Now(),
		"component": "eventbus",
	})
	return err == nil
}

func (e *EventBusComponent) IsHealthy() bool {
	// Similar to ready check
	return e.IsReady()
}

// ObservabilityComponent wraps observability service for lifecycle management
type ObservabilityComponent struct {
	Service *observability.Service
}

func (o *ObservabilityComponent) Name() string {
	return "observability"
}

func (o *ObservabilityComponent) Priority() int {
	return 4 // Medium priority - monitoring
}

func (o *ObservabilityComponent) Start(ctx context.Context) error {
	// Observability service is already started
	return nil
}

func (o *ObservabilityComponent) Stop(ctx context.Context) error {
	return o.Service.Stop(ctx)
}

func (o *ObservabilityComponent) HealthCheck(ctx context.Context) error {
	status := o.Service.HealthCheck()
	if !status.Healthy {
		return fmt.Errorf("observability service is not healthy: %s", status.Message)
	}
	return nil
}

func (o *ObservabilityComponent) IsReady() bool {
	return o.Service.HealthCheck().Healthy
}

func (o *ObservabilityComponent) IsHealthy() bool {
	return o.Service.HealthCheck().Healthy
}

// ComplianceComponent wraps compliance framework for lifecycle management
type ComplianceComponent struct {
	Framework *compliance.ComplianceFramework
}

func (c *ComplianceComponent) Name() string {
	return "compliance"
}

func (c *ComplianceComponent) Priority() int {
	return 5 // Lower priority - business logic
}

func (c *ComplianceComponent) Start(ctx context.Context) error {
	// Compliance framework is already initialized
	return nil
}

func (c *ComplianceComponent) Stop(ctx context.Context) error {
	// Compliance framework doesn't need explicit shutdown
	return nil
}

func (c *ComplianceComponent) HealthCheck(ctx context.Context) error {
	status, err := c.Framework.GetComplianceStatus(ctx)
	if err != nil {
		return fmt.Errorf("failed to get compliance status: %w", err)
	}
	
	if !status.Healthy {
		return fmt.Errorf("compliance framework is not healthy: violations=%d, errors=%d", 
			len(status.Violations), len(status.Errors))
	}
	
	return nil
}

func (c *ComplianceComponent) IsReady() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	status, err := c.Framework.GetComplianceStatus(ctx)
	if err != nil {
		return false
	}
	
	return status.Healthy
}

func (c *ComplianceComponent) IsHealthy() bool {
	return c.IsReady()
}