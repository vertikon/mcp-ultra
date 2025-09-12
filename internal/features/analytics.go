package features

import (
	"context"
	"sync"
	"time"

	"github.com/vertikon/mcp-ultra/internal/observability"
	"github.com/vertikon/mcp-ultra/pkg/logger"
)

// FeatureAnalytics provides analytics for feature flags and experiments
type FeatureAnalytics struct {
	config    AdvancedConfig
	logger    logger.Logger
	telemetry *observability.TelemetryService
	
	// State
	mu                sync.RWMutex
	exposures         map[string][]ExposureEvent
	conversions       map[string][]ConversionEvent
	flagEvaluations   map[string][]EvaluationEvent
	
	// Metrics
	metrics           AnalyticsMetrics
	
	// Background processing
	ctx               context.Context
	cancel            context.CancelFunc
	wg                sync.WaitGroup
}

// ExposureEvent represents a user exposure to a feature or experiment
type ExposureEvent struct {
	UserID        string                 `json:"user_id"`
	ExperimentID  string                 `json:"experiment_id"`
	VariantID     string                 `json:"variant_id"`
	FlagKey       string                 `json:"flag_key,omitempty"`
	Value         interface{}            `json:"value,omitempty"`
	Timestamp     time.Time              `json:"timestamp"`
	Properties    map[string]interface{} `json:"properties"`
	SessionID     string                 `json:"session_id,omitempty"`
	UserAgent     string                 `json:"user_agent,omitempty"`
	IPAddress     string                 `json:"ip_address,omitempty"`
	Country       string                 `json:"country,omitempty"`
}

// ConversionEvent represents a user conversion event
type ConversionEvent struct {
	UserID        string                 `json:"user_id"`
	ExperimentID  string                 `json:"experiment_id"`
	VariantID     string                 `json:"variant_id"`
	EventType     string                 `json:"event_type"`
	EventValue    float64                `json:"event_value"`
	Timestamp     time.Time              `json:"timestamp"`
	Properties    map[string]interface{} `json:"properties"`
	SessionID     string                 `json:"session_id,omitempty"`
}

// EvaluationEvent represents a flag evaluation
type EvaluationEvent struct {
	UserID        string                 `json:"user_id"`
	FlagKey       string                 `json:"flag_key"`
	Value         interface{}            `json:"value"`
	DefaultValue  interface{}            `json:"default_value"`
	Reason        string                 `json:"reason"`
	Timestamp     time.Time              `json:"timestamp"`
	Duration      time.Duration          `json:"duration"`
	UserContext   UserContext            `json:"user_context"`
}

// AnalyticsMetrics contains analytics metrics
type AnalyticsMetrics struct {
	TotalExposures      int64                           `json:"total_exposures"`
	TotalConversions    int64                           `json:"total_conversions"`
	TotalEvaluations    int64                           `json:"total_evaluations"`
	ExposuresByVariant  map[string]map[string]int64     `json:"exposures_by_variant"`
	ConversionsByVariant map[string]map[string]int64    `json:"conversions_by_variant"`
	FlagEvaluations     map[string]int64                `json:"flag_evaluations"`
	ConversionRates     map[string]map[string]float64   `json:"conversion_rates"`
	LastUpdated         time.Time                       `json:"last_updated"`
}

// ABTestingEngine provides A/B testing statistical analysis
type ABTestingEngine struct {
	config    AdvancedConfig
	logger    logger.Logger
	telemetry *observability.TelemetryService
	
	// Statistical analysis
	mu          sync.RWMutex
	testResults map[string]*StatisticalTestResult
	
	// Background processing
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

// StatisticalTestResult contains statistical test results
type StatisticalTestResult struct {
	ExperimentID        string                        `json:"experiment_id"`
	TestType            string                        `json:"test_type"`
	Confidence          float64                       `json:"confidence"`
	PValue              float64                       `json:"p_value"`
	EffectSize          float64                       `json:"effect_size"`
	PowerAnalysis       PowerAnalysis                 `json:"power_analysis"`
	VariantComparisons  []VariantComparison           `json:"variant_comparisons"`
	Recommendation      TestRecommendation            `json:"recommendation"`
	CalculatedAt        time.Time                     `json:"calculated_at"`
}

// VariantComparison compares two variants
type VariantComparison struct {
	ControlVariant    string  `json:"control_variant"`
	TreatmentVariant  string  `json:"treatment_variant"`
	ControlSamples    int64   `json:"control_samples"`
	TreatmentSamples  int64   `json:"treatment_samples"`
	ControlRate       float64 `json:"control_rate"`
	TreatmentRate     float64 `json:"treatment_rate"`
	Improvement       float64 `json:"improvement"`
	PValue            float64 `json:"p_value"`
	Significant       bool    `json:"significant"`
	ConfidenceInterval ConfidenceInterval `json:"confidence_interval"`
}

// TestRecommendation provides recommendations based on test results
type TestRecommendation struct {
	Action          string  `json:"action"` // "continue", "stop_winner", "stop_loser", "inconclusive"
	WinningVariant  string  `json:"winning_variant,omitempty"`
	Confidence      float64 `json:"confidence"`
	Reasoning       string  `json:"reasoning"`
}

// NewFeatureAnalytics creates a new feature analytics instance
func NewFeatureAnalytics(config AdvancedConfig, logger logger.Logger, telemetry *observability.TelemetryService) *FeatureAnalytics {
	ctx, cancel := context.WithCancel(context.Background())
	
	fa := &FeatureAnalytics{
		config:          config,
		logger:          logger,
		telemetry:       telemetry,
		exposures:       make(map[string][]ExposureEvent),
		conversions:     make(map[string][]ConversionEvent),
		flagEvaluations: make(map[string][]EvaluationEvent),
		metrics: AnalyticsMetrics{
			ExposuresByVariant:   make(map[string]map[string]int64),
			ConversionsByVariant: make(map[string]map[string]int64),
			FlagEvaluations:      make(map[string]int64),
			ConversionRates:      make(map[string]map[string]float64),
		},
		ctx:    ctx,
		cancel: cancel,
	}
	
	// Start background processing
	fa.startBackgroundTasks()
	
	return fa
}

// RecordExposure records a user exposure to a feature or experiment
func (fa *FeatureAnalytics) RecordExposure(experimentID, variantID, userID string) {
	fa.mu.Lock()
	defer fa.mu.Unlock()
	
	exposure := ExposureEvent{
		UserID:       userID,
		ExperimentID: experimentID,
		VariantID:    variantID,
		Timestamp:    time.Now(),
		Properties:   make(map[string]interface{}),
	}
	
	fa.exposures[experimentID] = append(fa.exposures[experimentID], exposure)
	fa.metrics.TotalExposures++
	
	// Update variant exposure counts
	if fa.metrics.ExposuresByVariant[experimentID] == nil {
		fa.metrics.ExposuresByVariant[experimentID] = make(map[string]int64)
	}
	fa.metrics.ExposuresByVariant[experimentID][variantID]++
	
	// Record in telemetry
	if fa.telemetry != nil {
		fa.telemetry.RecordCounter("experiment_exposures_total", 1, map[string]string{
			"experiment_id": experimentID,
			"variant_id":    variantID,
		})
	}
}

// RecordConversion records a conversion event
func (fa *FeatureAnalytics) RecordConversion(experimentID, variantID, userID, eventType string, value float64) {
	fa.mu.Lock()
	defer fa.mu.Unlock()
	
	conversion := ConversionEvent{
		UserID:       userID,
		ExperimentID: experimentID,
		VariantID:    variantID,
		EventType:    eventType,
		EventValue:   value,
		Timestamp:    time.Now(),
		Properties:   make(map[string]interface{}),
	}
	
	fa.conversions[experimentID] = append(fa.conversions[experimentID], conversion)
	fa.metrics.TotalConversions++
	
	// Update variant conversion counts
	if fa.metrics.ConversionsByVariant[experimentID] == nil {
		fa.metrics.ConversionsByVariant[experimentID] = make(map[string]int64)
	}
	fa.metrics.ConversionsByVariant[experimentID][variantID]++
	
	// Record in telemetry
	if fa.telemetry != nil {
		fa.telemetry.RecordCounter("experiment_conversions_total", 1, map[string]string{
			"experiment_id": experimentID,
			"variant_id":    variantID,
			"event_type":    eventType,
		})
		
		if value > 0 {
			fa.telemetry.RecordHistogram("experiment_conversion_value", value, map[string]string{
				"experiment_id": experimentID,
				"variant_id":    variantID,
				"event_type":    eventType,
			})
		}
	}
}

// RecordEvaluation records a flag evaluation
func (fa *FeatureAnalytics) RecordEvaluation(userID, flagKey string, value, defaultValue interface{}, reason string, duration time.Duration, userContext UserContext) {
	fa.mu.Lock()
	defer fa.mu.Unlock()
	
	evaluation := EvaluationEvent{
		UserID:       userID,
		FlagKey:      flagKey,
		Value:        value,
		DefaultValue: defaultValue,
		Reason:       reason,
		Timestamp:    time.Now(),
		Duration:     duration,
		UserContext:  userContext,
	}
	
	fa.flagEvaluations[flagKey] = append(fa.flagEvaluations[flagKey], evaluation)
	fa.metrics.TotalEvaluations++
	fa.metrics.FlagEvaluations[flagKey]++
	
	// Record in telemetry
	if fa.telemetry != nil {
		fa.telemetry.RecordCounter("flag_evaluations_total", 1, map[string]string{
			"flag_key": flagKey,
			"reason":   reason,
		})
		
		fa.telemetry.RecordHistogram("flag_evaluation_duration", float64(duration.Milliseconds()), map[string]string{
			"flag_key": flagKey,
		})
	}
}

// GetExperimentAnalytics returns analytics for a specific experiment
func (fa *FeatureAnalytics) GetExperimentAnalytics(experimentID string) ExperimentAnalytics {
	fa.mu.RLock()
	defer fa.mu.RUnlock()
	
	exposures := fa.exposures[experimentID]
	conversions := fa.conversions[experimentID]
	
	analytics := ExperimentAnalytics{
		ExperimentID:     experimentID,
		TotalExposures:   int64(len(exposures)),
		TotalConversions: int64(len(conversions)),
		VariantAnalytics: make(map[string]VariantAnalytics),
		Timeline:         fa.calculateTimeline(exposures, conversions),
	}
	
	// Group by variant
	variantExposures := make(map[string][]ExposureEvent)
	variantConversions := make(map[string][]ConversionEvent)
	
	for _, exposure := range exposures {
		variantExposures[exposure.VariantID] = append(variantExposures[exposure.VariantID], exposure)
	}
	
	for _, conversion := range conversions {
		variantConversions[conversion.VariantID] = append(variantConversions[conversion.VariantID], conversion)
	}
	
	// Calculate variant analytics
	for variantID, vExposures := range variantExposures {
		vConversions := variantConversions[variantID]
		
		variantAnalytics := VariantAnalytics{
			VariantID:        variantID,
			Exposures:        int64(len(vExposures)),
			Conversions:      int64(len(vConversions)),
			ConversionRate:   0,
			AverageValue:     0,
			UniqueUsers:      fa.countUniqueUsers(vExposures),
		}
		
		if variantAnalytics.Exposures > 0 {
			variantAnalytics.ConversionRate = float64(variantAnalytics.Conversions) / float64(variantAnalytics.Exposures)
		}
		
		if len(vConversions) > 0 {
			totalValue := 0.0
			for _, conv := range vConversions {
				totalValue += conv.EventValue
			}
			variantAnalytics.AverageValue = totalValue / float64(len(vConversions))
		}
		
		analytics.VariantAnalytics[variantID] = variantAnalytics
	}
	
	// Calculate overall conversion rate
	if analytics.TotalExposures > 0 {
		analytics.OverallConversionRate = float64(analytics.TotalConversions) / float64(analytics.TotalExposures)
	}
	
	return analytics
}

// GetFlagAnalytics returns analytics for a specific flag
func (fa *FeatureAnalytics) GetFlagAnalytics(flagKey string) FlagAnalytics {
	fa.mu.RLock()
	defer fa.mu.RUnlock()
	
	evaluations := fa.flagEvaluations[flagKey]
	
	analytics := FlagAnalytics{
		FlagKey:         flagKey,
		TotalEvaluations: int64(len(evaluations)),
		UniqueUsers:     fa.countUniqueUsersFromEvaluations(evaluations),
		ValueDistribution: fa.calculateValueDistribution(evaluations),
	}
	
	return analytics
}

// GetOverallMetrics returns overall analytics metrics
func (fa *FeatureAnalytics) GetOverallMetrics() AnalyticsMetrics {
	fa.mu.RLock()
	defer fa.mu.RUnlock()
	
	// Update conversion rates
	fa.calculateConversionRates()
	
	metrics := fa.metrics
	metrics.LastUpdated = time.Now()
	
	return metrics
}

// ProcessMetrics processes and updates metrics (called by background task)
func (fa *FeatureAnalytics) ProcessMetrics() {
	fa.mu.Lock()
	defer fa.mu.Unlock()
	
	// Calculate conversion rates
	fa.calculateConversionRates()
	
	// Clean up old data if needed
	fa.cleanupOldData()
	
	fa.metrics.LastUpdated = time.Now()
}

// Close gracefully shuts down analytics
func (fa *FeatureAnalytics) Close() error {
	fa.cancel()
	fa.wg.Wait()
	return nil
}

// ExperimentAnalytics contains analytics for a specific experiment
type ExperimentAnalytics struct {
	ExperimentID           string                      `json:"experiment_id"`
	TotalExposures         int64                       `json:"total_exposures"`
	TotalConversions       int64                       `json:"total_conversions"`
	OverallConversionRate  float64                     `json:"overall_conversion_rate"`
	VariantAnalytics       map[string]VariantAnalytics `json:"variant_analytics"`
	Timeline               []TimelinePoint             `json:"timeline"`
}

// VariantAnalytics contains analytics for a specific variant
type VariantAnalytics struct {
	VariantID      string  `json:"variant_id"`
	Exposures      int64   `json:"exposures"`
	Conversions    int64   `json:"conversions"`
	ConversionRate float64 `json:"conversion_rate"`
	AverageValue   float64 `json:"average_value"`
	UniqueUsers    int64   `json:"unique_users"`
}

// FlagAnalytics contains analytics for a specific flag
type FlagAnalytics struct {
	FlagKey           string                    `json:"flag_key"`
	TotalEvaluations  int64                     `json:"total_evaluations"`
	UniqueUsers       int64                     `json:"unique_users"`
	ValueDistribution map[string]int64          `json:"value_distribution"`
}

// TimelinePoint represents a point in the experiment timeline
type TimelinePoint struct {
	Timestamp        time.Time `json:"timestamp"`
	Exposures        int64     `json:"exposures"`
	Conversions      int64     `json:"conversions"`
	ConversionRate   float64   `json:"conversion_rate"`
}

// NewABTestingEngine creates a new A/B testing engine
func NewABTestingEngine(config AdvancedConfig, logger logger.Logger, telemetry *observability.TelemetryService) *ABTestingEngine {
	ctx, cancel := context.WithCancel(context.Background())
	
	engine := &ABTestingEngine{
		config:      config,
		logger:      logger,
		telemetry:   telemetry,
		testResults: make(map[string]*StatisticalTestResult),
		ctx:         ctx,
		cancel:      cancel,
	}
	
	// Start background processing
	engine.startBackgroundTasks()
	
	return engine
}

// CalculateResults calculates statistical results for an experiment
func (ate *ABTestingEngine) CalculateResults(experimentID string) *ExperimentResults {
	ate.mu.Lock()
	defer ate.mu.Unlock()
	
	testResult, exists := ate.testResults[experimentID]
	if !exists {
		// Perform statistical analysis
		testResult = ate.performStatisticalAnalysis(experimentID)
		ate.testResults[experimentID] = testResult
	}
	
	if testResult == nil {
		return nil
	}
	
	// Convert to experiment results format
	results := &ExperimentResults{
		Status:            ate.determineResultStatus(testResult),
		Confidence:        testResult.Confidence,
		SignificantResult: testResult.PValue < (1.0 - ate.config.StatisticalSignificance),
		WinningVariant:    testResult.Recommendation.WinningVariant,
		VariantResults:    make(map[string]VariantResult),
		PValue:            testResult.PValue,
		EffectSize:        testResult.EffectSize,
		PowerAnalysis:     testResult.PowerAnalysis,
	}
	
	// Convert variant comparisons to variant results
	for _, comparison := range testResult.VariantComparisons {
		// Control variant
		results.VariantResults[comparison.ControlVariant] = VariantResult{
			VariantID:      comparison.ControlVariant,
			Samples:        comparison.ControlSamples,
			ConversionRate: comparison.ControlRate,
			Confidence:     comparison.ConfidenceInterval,
			Metrics:        make(map[string]MetricResult),
		}
		
		// Treatment variant
		results.VariantResults[comparison.TreatmentVariant] = VariantResult{
			VariantID:      comparison.TreatmentVariant,
			Samples:        comparison.TreatmentSamples,
			ConversionRate: comparison.TreatmentRate,
			Confidence:     comparison.ConfidenceInterval,
			Metrics:        make(map[string]MetricResult),
		}
	}
	
	// Calculate total samples
	for _, variant := range results.VariantResults {
		results.TotalSamples += variant.Samples
	}
	
	return results
}

// GetTestResult returns statistical test results for an experiment
func (ate *ABTestingEngine) GetTestResult(experimentID string) *StatisticalTestResult {
	ate.mu.RLock()
	defer ate.mu.RUnlock()
	
	return ate.testResults[experimentID]
}

// Close gracefully shuts down the A/B testing engine
func (ate *ABTestingEngine) Close() error {
	ate.cancel()
	ate.wg.Wait()
	return nil
}

// Private methods

func (fa *FeatureAnalytics) startBackgroundTasks() {
	// Metrics processing task
	fa.wg.Add(1)
	go fa.metricsProcessingTask()
	
	// Data cleanup task
	fa.wg.Add(1)
	go fa.dataCleanupTask()
}

func (fa *FeatureAnalytics) metricsProcessingTask() {
	defer fa.wg.Done()
	
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-fa.ctx.Done():
			return
		case <-ticker.C:
			fa.ProcessMetrics()
		}
	}
}

func (fa *FeatureAnalytics) dataCleanupTask() {
	defer fa.wg.Done()
	
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	
	for {
		select {
		case <-fa.ctx.Done():
			return
		case <-ticker.C:
			fa.cleanupOldData()
		}
	}
}

func (fa *FeatureAnalytics) calculateConversionRates() {
	for experimentID := range fa.exposures {
		if fa.metrics.ConversionRates[experimentID] == nil {
			fa.metrics.ConversionRates[experimentID] = make(map[string]float64)
		}
		
		for variantID, exposures := range fa.metrics.ExposuresByVariant[experimentID] {
			conversions := fa.metrics.ConversionsByVariant[experimentID][variantID]
			
			if exposures > 0 {
				fa.metrics.ConversionRates[experimentID][variantID] = float64(conversions) / float64(exposures)
			}
		}
	}
}

func (fa *FeatureAnalytics) cleanupOldData() {
	cutoff := time.Now().Add(-fa.config.AnalyticsRetention)
	
	// Clean exposures
	for experimentID, exposures := range fa.exposures {
		filtered := make([]ExposureEvent, 0)
		for _, exposure := range exposures {
			if exposure.Timestamp.After(cutoff) {
				filtered = append(filtered, exposure)
			}
		}
		fa.exposures[experimentID] = filtered
	}
	
	// Clean conversions
	for experimentID, conversions := range fa.conversions {
		filtered := make([]ConversionEvent, 0)
		for _, conversion := range conversions {
			if conversion.Timestamp.After(cutoff) {
				filtered = append(filtered, conversion)
			}
		}
		fa.conversions[experimentID] = filtered
	}
	
	// Clean evaluations
	for flagKey, evaluations := range fa.flagEvaluations {
		filtered := make([]EvaluationEvent, 0)
		for _, evaluation := range evaluations {
			if evaluation.Timestamp.After(cutoff) {
				filtered = append(filtered, evaluation)
			}
		}
		fa.flagEvaluations[flagKey] = filtered
	}
}

func (fa *FeatureAnalytics) countUniqueUsers(exposures []ExposureEvent) int64 {
	users := make(map[string]bool)
	for _, exposure := range exposures {
		users[exposure.UserID] = true
	}
	return int64(len(users))
}

func (fa *FeatureAnalytics) countUniqueUsersFromEvaluations(evaluations []EvaluationEvent) int64 {
	users := make(map[string]bool)
	for _, evaluation := range evaluations {
		users[evaluation.UserID] = true
	}
	return int64(len(users))
}

func (fa *FeatureAnalytics) calculateValueDistribution(evaluations []EvaluationEvent) map[string]int64 {
	distribution := make(map[string]int64)
	for _, evaluation := range evaluations {
		valueStr := fa.valueToString(evaluation.Value)
		distribution[valueStr]++
	}
	return distribution
}

func (fa *FeatureAnalytics) valueToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	default:
		return "other"
	}
}

func (fa *FeatureAnalytics) calculateTimeline(exposures []ExposureEvent, conversions []ConversionEvent) []TimelinePoint {
	// This is simplified - in reality you'd create hourly/daily buckets
	timeline := make([]TimelinePoint, 0)
	
	if len(exposures) == 0 {
		return timeline
	}
	
	// Create a single timeline point for demonstration
	point := TimelinePoint{
		Timestamp:      time.Now(),
		Exposures:      int64(len(exposures)),
		Conversions:    int64(len(conversions)),
		ConversionRate: 0,
	}
	
	if point.Exposures > 0 {
		point.ConversionRate = float64(point.Conversions) / float64(point.Exposures)
	}
	
	timeline = append(timeline, point)
	
	return timeline
}

func (ate *ABTestingEngine) startBackgroundTasks() {
	// Statistical analysis task
	ate.wg.Add(1)
	go ate.statisticalAnalysisTask()
}

func (ate *ABTestingEngine) statisticalAnalysisTask() {
	defer ate.wg.Done()
	
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-ate.ctx.Done():
			return
		case <-ticker.C:
			ate.performBatchAnalysis()
		}
	}
}

func (ate *ABTestingEngine) performBatchAnalysis() {
	// This would analyze all active experiments
	ate.logger.Debug("Performing batch statistical analysis")
}

func (ate *ABTestingEngine) performStatisticalAnalysis(experimentID string) *StatisticalTestResult {
	// This is a simplified implementation
	// In reality, you'd perform comprehensive statistical tests like:
	// - Chi-square test for categorical data
	// - T-test for continuous data
	// - Welch's t-test for unequal variances
	// - Bayesian analysis
	
	result := &StatisticalTestResult{
		ExperimentID: experimentID,
		TestType:     "chi_square",
		Confidence:   0.95,
		PValue:       0.05, // Placeholder
		EffectSize:   0.1,  // Placeholder
		PowerAnalysis: PowerAnalysis{
			Power:              0.8,
			MinDetectableEffect: 0.05,
			RecommendedSamples: 1000,
		},
		Recommendation: TestRecommendation{
			Action:     "continue",
			Confidence: 0.95,
			Reasoning:  "Insufficient data for conclusive results",
		},
		CalculatedAt: time.Now(),
	}
	
	return result
}

func (ate *ABTestingEngine) determineResultStatus(result *StatisticalTestResult) ResultStatus {
	if result.PValue < 0.05 && result.Confidence >= ate.config.StatisticalSignificance {
		return ResultStatusSignificant
	} else if result.PValue >= 0.05 {
		return ResultStatusInsignificant
	}
	return ResultStatusInconclusive
}