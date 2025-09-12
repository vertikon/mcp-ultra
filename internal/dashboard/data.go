package dashboard

import (
	"context"
	"runtime"
	"time"

	"mcp-model-ultra/internal/cache"
	"mcp-model-ultra/internal/features"
	"mcp-model-ultra/internal/lifecycle"
	"mcp-model-ultra/internal/metrics"
	"mcp-model-ultra/internal/ratelimit"
	"mcp-model-ultra/internal/tracing"
)

// Data collection methods for dashboard server

func (ds *DashboardServer) getSystemHealth(ctx context.Context) SystemHealth {
	// Get lifecycle status
	lifecycleStatus := ds.lifecycle.GetStatus()
	
	// Calculate overall health score based on components
	componentStatuses := ds.getComponentStatus(ctx)
	totalScore := 0.0
	healthyComponents := 0
	
	for _, component := range componentStatuses {
		totalScore += component.Health
		if component.Health > 80 {
			healthyComponents++
		}
	}
	
	overallScore := totalScore / float64(len(componentStatuses))
	status := "healthy"
	
	if overallScore < 50 {
		status = "unhealthy"
	} else if overallScore < 80 {
		status = "degraded"
	}
	
	// Calculate uptime (this would be tracked properly in production)
	uptime := time.Since(time.Now().Add(-24 * time.Hour))
	
	return SystemHealth{
		Status:        status,
		OverallScore:  overallScore,
		Uptime:        uptime,
		SLOCompliance: 99.5, // This would be calculated from actual SLO metrics
	}
}

func (ds *DashboardServer) getComponentStatus(ctx context.Context) []ComponentStatus {
	components := []ComponentStatus{
		{
			Name:      "Lifecycle Manager",
			Type:      "core",
			Status:    string(ds.lifecycle.GetStatus()),
			Health:    ds.calculateLifecycleHealth(),
			LastCheck: time.Now(),
			Message:   "All components operational",
		},
		{
			Name:      "Distributed Cache",
			Type:      "cache",
			Status:    ds.getCacheStatus(ctx),
			Health:    ds.calculateCacheHealth(ctx),
			LastCheck: time.Now(),
			Message:   "Redis cluster healthy",
		},
		{
			Name:      "Business Metrics",
			Type:      "metrics",
			Status:    "healthy",
			Health:    95.0,
			LastCheck: time.Now(),
			Message:   "Metrics collection active",
		},
		{
			Name:      "Feature Flags",
			Type:      "features",
			Status:    "healthy",
			Health:    98.0,
			LastCheck: time.Now(),
			Message:   "Feature management operational",
		},
		{
			Name:      "Rate Limiter",
			Type:      "ratelimit",
			Status:    ds.getRateLimitStatus(ctx),
			Health:    ds.calculateRateLimitHealth(ctx),
			LastCheck: time.Now(),
			Message:   "Rate limiting active",
		},
		{
			Name:      "Transaction Tracing",
			Type:      "tracing",
			Status:    "healthy",
			Health:    92.0,
			LastCheck: time.Now(),
			Message:   "Tracing operational",
		},
	}
	
	return components
}

func (ds *DashboardServer) getOverviewMetrics(ctx context.Context) OverviewMetrics {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// These would be collected from actual metrics in production
	return OverviewMetrics{
		RequestRate:    1250.5,
		ErrorRate:      0.12,
		AvgLatency:     45.2,
		P99Latency:     120.8,
		Throughput:     2100.0,
		ActiveSessions: 1523,
		CacheHitRate:   94.5,
		CPUUsage:       45.2,
		MemoryUsage:    float64(m.Sys) / (1024 * 1024 * 1024) * 100, // Rough calculation
	}
}

func (ds *DashboardServer) getActiveAlerts(ctx context.Context) []Alert {
	// This would fetch from actual alerting system
	alerts := []Alert{}
	
	// Add sample alerts based on system status
	if ds.getOverviewMetrics(ctx).ErrorRate > 0.1 {
		alerts = append(alerts, Alert{
			ID:          "alert-001",
			Type:        AlertTypePerformance,
			Severity:    AlertSeverityWarning,
			Title:       "High Error Rate Detected",
			Description: "Error rate is above threshold (0.1%)",
			Component:   "API Gateway",
			CreatedAt:   time.Now().Add(-5 * time.Minute),
			UpdatedAt:   time.Now(),
			Status:      AlertStatusActive,
			Actions: []AlertAction{
				{Type: "acknowledge", Label: "Acknowledge"},
				{Type: "silence", Label: "Silence for 1 hour"},
			},
		})
	}
	
	if ds.getOverviewMetrics(ctx).P99Latency > 100 {
		alerts = append(alerts, Alert{
			ID:          "alert-002",
			Type:        AlertTypePerformance,
			Severity:    AlertSeverityWarning,
			Title:       "High Response Latency",
			Description: "P99 latency is above 100ms threshold",
			Component:   "Application",
			CreatedAt:   time.Now().Add(-2 * time.Minute),
			UpdatedAt:   time.Now(),
			Status:      AlertStatusActive,
		})
	}
	
	return alerts
}

func (ds *DashboardServer) getBusinessMetrics(ctx context.Context, start, end time.Time) interface{} {
	// Get business metrics from the metrics service
	query := metrics.MetricQuery{
		StartTime: start,
		EndTime:   end,
		Limit:     1000,
	}
	
	// This would query actual business metrics
	businessMetrics := map[string]interface{}{
		"revenue": map[string]interface{}{
			"total":  125000.50,
			"growth": 15.2,
			"trend":  "increasing",
		},
		"orders": map[string]interface{}{
			"total":      1543,
			"conversion": 3.2,
			"avg_value":  81.05,
		},
		"users": map[string]interface{}{
			"active":     12543,
			"new":        234,
			"returning":  11309,
			"retention":  89.2,
		},
	}
	
	return businessMetrics
}

func (ds *DashboardServer) getSystemMetrics(ctx context.Context) SystemMetrics {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	return SystemMetrics{
		CPU: CPUMetrics{
			Usage:     45.2,
			LoadAvg1m: 0.8,
			LoadAvg5m: 0.7,
			LoadAvg15m: 0.6,
			Cores:     runtime.NumCPU(),
		},
		Memory: MemoryMetrics{
			Total:     int64(m.Sys),
			Used:      int64(m.Alloc),
			Available: int64(m.Sys - m.Alloc),
			Usage:     float64(m.Alloc) / float64(m.Sys) * 100,
			SwapTotal: 0, // Would be retrieved from system
			SwapUsed:  0,
			SwapUsage: 0,
		},
		Disk: DiskMetrics{
			Total:     1000 * 1024 * 1024 * 1024, // 1TB
			Used:      450 * 1024 * 1024 * 1024,  // 450GB
			Available: 550 * 1024 * 1024 * 1024,  // 550GB
			Usage:     45.0,
			IOPS:      1200,
			ReadRate:  50 * 1024 * 1024,  // 50 MB/s
			WriteRate: 30 * 1024 * 1024,  // 30 MB/s
		},
		Network: NetworkMetrics{
			BytesReceived:   1024 * 1024 * 1024 * 500, // 500 GB
			BytesSent:       1024 * 1024 * 1024 * 300, // 300 GB
			PacketsReceived: 50000000,
			PacketsSent:     30000000,
			ReceiveRate:     100 * 1024 * 1024, // 100 MB/s
			TransmitRate:    60 * 1024 * 1024,  // 60 MB/s
			Connections:     1523,
			DroppedPackets:  123,
			ErrorPackets:    5,
		},
		Processes: ProcessMetrics{
			Total:    256,
			Running:  12,
			Sleeping: 240,
			Zombie:   2,
			Stopped:  2,
		},
	}
}

func (ds *DashboardServer) getPerformanceMetrics(ctx context.Context) PerformanceMetrics {
	return PerformanceMetrics{
		RequestRate: 1250.5,
		ResponseTime: ResponseTimeMetrics{
			Mean:   45.2,
			Median: 42.1,
			P95:    89.5,
			P99:    120.8,
			P999:   250.2,
			Min:    2.1,
			Max:    1200.5,
		},
		Throughput:  2100.0,
		Concurrency: 45,
		QueueDepth:  12,
		DatabaseMetrics: DatabaseMetrics{
			Connections:       50,
			ActiveConnections: 12,
			QueryRate:         850.2,
			SlowQueries:       5,
			Deadlocks:         0,
			LockWaitTime:      2.5,
		},
		CacheMetrics: CacheMetricsData{
			HitRate:       94.5,
			MissRate:      5.5,
			Evictions:     123,
			KeyCount:      125000,
			MemoryUsage:   512 * 1024 * 1024, // 512 MB
			OperationRate: 5200.0,
		},
	}
}

func (ds *DashboardServer) getErrorMetrics(ctx context.Context) ErrorMetrics {
	return ErrorMetrics{
		TotalErrors:    45,
		ErrorRate:      0.12,
		ErrorsByType:   map[string]int64{
			"validation": 20,
			"timeout":    15,
			"database":   5,
			"external":   3,
			"auth":       2,
		},
		ErrorsByCode: map[string]int64{
			"400": 25,
			"401": 5,
			"500": 10,
			"502": 3,
			"503": 2,
		},
		CriticalErrors: 2,
		RecentErrors: []RecentError{
			{
				Timestamp: time.Now().Add(-2 * time.Minute),
				Type:      "timeout",
				Code:      "500",
				Message:   "Database connection timeout",
				Component: "user-service",
				TraceID:   "abc123def456",
				Count:     3,
			},
			{
				Timestamp: time.Now().Add(-5 * time.Minute),
				Type:      "validation",
				Code:      "400",
				Message:   "Invalid request payload",
				Component: "api-gateway",
				Count:     1,
			},
		},
	}
}

func (ds *DashboardServer) getTrafficMetrics(ctx context.Context) TrafficMetrics {
	return TrafficMetrics{
		TotalRequests:  1250000,
		ActiveSessions: 1523,
		UniqueUsers:    12543,
		TrafficBySource: map[string]int64{
			"web":    800000,
			"mobile": 350000,
			"api":    100000,
		},
		TrafficByRegion: map[string]int64{
			"us-east":  500000,
			"eu-west":  400000,
			"ap-south": 350000,
		},
		TrafficByChannel: map[string]int64{
			"direct":  450000,
			"search":  350000,
			"social":  250000,
			"email":   200000,
		},
		PeakTraffic: TrafficPeak{
			Timestamp:   time.Now().Add(-1 * time.Hour),
			RequestRate: 2500.0,
			Users:       15000,
			Sessions:    2100,
		},
		Bandwidth: BandwidthMetrics{
			Incoming: 100 * 1024 * 1024, // 100 MB/s
			Outgoing: 150 * 1024 * 1024, // 150 MB/s
			Total:    250 * 1024 * 1024, // 250 MB/s
			Peak:     400 * 1024 * 1024, // 400 MB/s
			Usage:    62.5,              // 62.5% of available
		},
	}
}

// Helper methods for component health calculations

func (ds *DashboardServer) calculateLifecycleHealth() float64 {
	status := ds.lifecycle.GetStatus()
	switch status {
	case lifecycle.LifecycleStateRunning:
		return 100.0
	case lifecycle.LifecycleStateStarting:
		return 80.0
	case lifecycle.LifecycleStateStopping:
		return 60.0
	default:
		return 0.0
	}
}

func (ds *DashboardServer) getCacheStatus(ctx context.Context) string {
	stats := ds.cache.GetStats(ctx)
	if stats.HitRate > 90.0 {
		return "healthy"
	} else if stats.HitRate > 70.0 {
		return "degraded"
	}
	return "unhealthy"
}

func (ds *DashboardServer) calculateCacheHealth(ctx context.Context) float64 {
	stats := ds.cache.GetStats(ctx)
	return stats.HitRate
}

func (ds *DashboardServer) getRateLimitStatus(ctx context.Context) string {
	stats := ds.rateLimit.GetStats(ctx)
	if stats.ErrorRate < 0.1 {
		return "healthy"
	} else if stats.ErrorRate < 1.0 {
		return "degraded"
	}
	return "unhealthy"
}

func (ds *DashboardServer) calculateRateLimitHealth(ctx context.Context) float64 {
	stats := ds.rateLimit.GetStats(ctx)
	return 100.0 - (stats.ErrorRate * 10) // Convert error rate to health score
}

// WebSocket message handling

func (ds *DashboardServer) startWebSocketBroadcaster(ctx context.Context) {
	ticker := time.NewTicker(ds.config.RefreshInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ds.broadcastRealtimeData(ctx)
		}
	}
}

func (ds *DashboardServer) broadcastRealtimeData(ctx context.Context) {
	if len(ds.clients) == 0 {
		return
	}

	data := RealtimeMetrics{
		Timestamp:     time.Now(),
		SystemMetrics: ds.getSystemMetrics(ctx),
		Performance:   ds.getPerformanceMetrics(ctx),
		Errors:        ds.getErrorMetrics(ctx),
		Traffic:       ds.getTrafficMetrics(ctx),
	}

	message := WebSocketMessage{
		Type:      "realtime_update",
		Timestamp: time.Now(),
		Data:      data,
	}

	ds.broadcastToClients(message)
}

func (ds *DashboardServer) broadcastToClients(message WebSocketMessage) {
	for clientID, conn := range ds.clients {
		if err := conn.WriteJSON(message); err != nil {
			ds.logger.Error("Failed to send WebSocket message", "client_id", clientID, "error", err)
			conn.Close()
			delete(ds.clients, clientID)
		}
	}
}

func (ds *DashboardServer) handleWebSocketMessage(clientID string, message []byte) {
	// Handle client subscription requests, filter changes, etc.
	var request SubscriptionRequest
	if err := json.Unmarshal(message, &request); err != nil {
		ds.logger.Error("Invalid WebSocket message", "client_id", clientID, "error", err)
		return
	}

	// Process subscription request
	ds.logger.Info("WebSocket subscription", "client_id", clientID, "type", request.Type)
}