package features

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/vertikon/mcp-ultra/internal/observability"
	"github.com/vertikon/mcp-ultra/pkg/logger"
)

// AdvancedFlagManager provides sophisticated feature flag management
type AdvancedFlagManager struct {
	flagManager *FlagManager // Embed the basic flag manager
	
	// Advanced configuration
	config           AdvancedConfig
	logger           logger.Logger
	telemetry        *observability.TelemetryService
	
	// State management
	mu               sync.RWMutex
	experiments      map[string]*Experiment
	segments         map[string]*UserSegment
	rolloutStrategies map[string]*RolloutStrategy
	flagDependencies map[string][]string
	
	// Analytics and monitoring
	analytics        *FeatureAnalytics
	abTesting        *ABTestingEngine
	
	// Background processing
	ctx              context.Context
	cancel           context.CancelFunc
	wg               sync.WaitGroup
}

// AdvancedConfig configures advanced feature management
type AdvancedConfig struct {
	// Experimentation
	ExperimentationEnabled bool                 `yaml:"experimentation_enabled"`
	ExperimentDefaults     ExperimentDefaults   `yaml:"experiment_defaults"`
	StatisticalSignificance float64             `yaml:"statistical_significance"`
	MinSampleSize          int                  `yaml:"min_sample_size"`
	
	// Rollout strategies
	DefaultRolloutStrategy string               `yaml:"default_rollout_strategy"`
	RolloutStrategies      []RolloutStrategy    `yaml:"rollout_strategies"`
	
	// User segmentation
	SegmentationEnabled    bool                 `yaml:"segmentation_enabled"`
	DefaultSegments        []UserSegment        `yaml:"default_segments"`
	
	// Analytics
	AnalyticsEnabled       bool                 `yaml:"analytics_enabled"`
	AnalyticsRetention     time.Duration        `yaml:"analytics_retention"`
	
	// Dependencies
	DependencyTracking     bool                 `yaml:"dependency_tracking"`
	
	// Performance
	CacheEnabled           bool                 `yaml:"cache_enabled"`
	CacheTTL              time.Duration        `yaml:"cache_ttl"`
	EvaluationTimeout      time.Duration        `yaml:"evaluation_timeout"`
	
	// Safety
	KillSwitchEnabled      bool                 `yaml:"kill_switch_enabled"`
	CircuitBreakerEnabled  bool                 `yaml:"circuit_breaker_enabled"`
	MaxEvaluationsPerSecond int                 `yaml:"max_evaluations_per_second"`
}

// Experiment represents an A/B test or multivariate experiment
type Experiment struct {
	ID                string                 `json:"id" yaml:"id"`
	Name              string                 `json:"name" yaml:"name"`
	Description       string                 `json:"description" yaml:"description"`
	Status            ExperimentStatus       `json:"status" yaml:"status"`
	Type              ExperimentType         `json:"type" yaml:"type"`
	
	// Timing
	StartTime         time.Time              `json:"start_time" yaml:"start_time"`
	EndTime           time.Time              `json:"end_time" yaml:"end_time"`
	Duration          time.Duration          `json:"duration" yaml:"duration"`
	
	// Traffic allocation
	TrafficPercent    float64                `json:"traffic_percent" yaml:"traffic_percent"`
	MinTrafficPercent float64                `json:"min_traffic_percent" yaml:"min_traffic_percent"`
	MaxTrafficPercent float64                `json:"max_traffic_percent" yaml:"max_traffic_percent"`
	
	// Variants
	Variants          []ExperimentVariant    `json:"variants" yaml:"variants"`
	
	// Targeting
	TargetingRules    []TargetingRule        `json:"targeting_rules" yaml:"targeting_rules"`
	UserSegments      []string               `json:"user_segments" yaml:"user_segments"`
	
	// Metrics
	PrimaryMetric     string                 `json:"primary_metric" yaml:"primary_metric"`
	SecondaryMetrics  []string               `json:"secondary_metrics" yaml:"secondary_metrics"`
	GuardrailMetrics  []string               `json:"guardrail_metrics" yaml:"guardrail_metrics"`
	
	// Configuration
	Config            map[string]interface{} `json:"config" yaml:"config"`
	Tags              []string               `json:"tags" yaml:"tags"`
	Owner             string                 `json:"owner" yaml:"owner"`
	
	// Results
	Results           *ExperimentResults     `json:"results,omitempty"`
}

// ExperimentStatus represents the status of an experiment
type ExperimentStatus string

const (
	ExperimentStatusDraft      ExperimentStatus = "draft"
	ExperimentStatusScheduled  ExperimentStatus = "scheduled"  
	ExperimentStatusRunning    ExperimentStatus = "running"
	ExperimentStatusPaused     ExperimentStatus = "paused"
	ExperimentStatusCompleted  ExperimentStatus = "completed"
	ExperimentStatusArchived   ExperimentStatus = "archived"
)

// ExperimentType represents the type of experiment
type ExperimentType string

const (
	ExperimentTypeAB          ExperimentType = "ab"
	ExperimentTypeMultivariate ExperimentType = "multivariate"
	ExperimentTypeBandit      ExperimentType = "bandit"
)

// ExperimentVariant represents a variant in an experiment
type ExperimentVariant struct {
	ID          string                 `json:"id" yaml:"id"`
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Weight      float64                `json:"weight" yaml:"weight"`
	Config      map[string]interface{} `json:"config" yaml:"config"`
	IsControl   bool                   `json:"is_control" yaml:"is_control"`
}

// ExperimentResults contains experiment results
type ExperimentResults struct {
	Status            ResultStatus                `json:"status"`
	Confidence        float64                     `json:"confidence"`
	SignificantResult bool                        `json:"significant_result"`
	WinningVariant    string                      `json:"winning_variant"`
	VariantResults    map[string]VariantResult    `json:"variant_results"`
	StartTime         time.Time                   `json:"start_time"`
	EndTime           time.Time                   `json:"end_time"`
	Duration          time.Duration               `json:"duration"`
	TotalSamples      int64                       `json:"total_samples"`
	
	// Statistical measures
	PValue            float64                     `json:"p_value"`
	EffectSize        float64                     `json:"effect_size"`
	PowerAnalysis     PowerAnalysis               `json:"power_analysis"`
}

// ResultStatus represents experiment result status
type ResultStatus string

const (
	ResultStatusInProgress    ResultStatus = "in_progress"
	ResultStatusSignificant   ResultStatus = "significant"
	ResultStatusInsignificant ResultStatus = "insignificant"
	ResultStatusInconclusive  ResultStatus = "inconclusive"
)

// VariantResult contains results for a specific variant
type VariantResult struct {
	VariantID       string                 `json:"variant_id"`
	Samples         int64                  `json:"samples"`
	Conversions     int64                  `json:"conversions"`
	ConversionRate  float64                `json:"conversion_rate"`
	Confidence      ConfidenceInterval     `json:"confidence"`
	Metrics         map[string]MetricResult `json:"metrics"`
}

// ConfidenceInterval represents a confidence interval
type ConfidenceInterval struct {
	Lower      float64 `json:"lower"`
	Upper      float64 `json:"upper"`
	Level      float64 `json:"level"`
}

// MetricResult contains metric-specific results
type MetricResult struct {
	MetricName string  `json:"metric_name"`
	Value      float64 `json:"value"`
	Change     float64 `json:"change"`
	Relative   float64 `json:"relative"`
	PValue     float64 `json:"p_value"`
}

// PowerAnalysis contains statistical power analysis
type PowerAnalysis struct {
	Power              float64 `json:"power"`
	MinDetectableEffect float64 `json:"min_detectable_effect"`
	RecommendedSamples int64   `json:"recommended_samples"`
}

// UserSegment represents a user segment for targeting
type UserSegment struct {
	ID          string                 `json:"id" yaml:"id"`
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Rules       []SegmentRule          `json:"rules" yaml:"rules"`
	Size        int64                  `json:"size" yaml:"size"`
	Tags        []string               `json:"tags" yaml:"tags"`
	CreatedAt   time.Time              `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" yaml:"updated_at"`
}

// SegmentRule represents a rule for user segmentation
type SegmentRule struct {
	Attribute  string      `json:"attribute" yaml:"attribute"`
	Operator   string      `json:"operator" yaml:"operator"`
	Value      interface{} `json:"value" yaml:"value"`
	Type       string      `json:"type" yaml:"type"`
}

// RolloutStrategy defines how features are rolled out
type RolloutStrategy struct {
	ID               string                 `json:"id" yaml:"id"`
	Name             string                 `json:"name" yaml:"name"`
	Type             RolloutType            `json:"type" yaml:"type"`
	Config           map[string]interface{} `json:"config" yaml:"config"`
	Stages           []RolloutStage         `json:"stages" yaml:"stages"`
	SafetyChecks     []SafetyCheck          `json:"safety_checks" yaml:"safety_checks"`
	RollbackStrategy RollbackStrategy       `json:"rollback_strategy" yaml:"rollback_strategy"`
}

// RolloutType represents the type of rollout
type RolloutType string

const (
	RolloutTypePercentage    RolloutType = "percentage"
	RolloutTypeCanary        RolloutType = "canary"
	RolloutTypeBlueGreen     RolloutType = "blue_green"
	RolloutTypeRing          RolloutType = "ring"
	RolloutTypeGeographic    RolloutType = "geographic"
	RolloutTypeScheduled     RolloutType = "scheduled"
)

// RolloutStage represents a stage in a rollout strategy
type RolloutStage struct {
	ID             string        `json:"id" yaml:"id"`
	Name           string        `json:"name" yaml:"name"`
	Percentage     float64       `json:"percentage" yaml:"percentage"`
	Duration       time.Duration `json:"duration" yaml:"duration"`
	SuccessCriteria []string     `json:"success_criteria" yaml:"success_criteria"`
	FailureCriteria []string     `json:"failure_criteria" yaml:"failure_criteria"`
	AutoAdvance    bool          `json:"auto_advance" yaml:"auto_advance"`
}

// SafetyCheck represents a safety check during rollout
type SafetyCheck struct {
	Type        string                 `json:"type" yaml:"type"`
	Metric      string                 `json:"metric" yaml:"metric"`
	Threshold   float64                `json:"threshold" yaml:"threshold"`
	Operator    string                 `json:"operator" yaml:"operator"`
	Window      time.Duration          `json:"window" yaml:"window"`
	Action      string                 `json:"action" yaml:"action"`
	Config      map[string]interface{} `json:"config" yaml:"config"`
}

// RollbackStrategy defines how to rollback a feature
type RollbackStrategy struct {
	Type       string                 `json:"type" yaml:"type"`
	Automatic  bool                   `json:"automatic" yaml:"automatic"`
	Triggers   []string               `json:"triggers" yaml:"triggers"`
	Config     map[string]interface{} `json:"config" yaml:"config"`
}

// TargetingRule represents a targeting rule for experiments
type TargetingRule struct {
	Attribute  string      `json:"attribute" yaml:"attribute"`
	Operator   string      `json:"operator" yaml:"operator"`
	Value      interface{} `json:"value" yaml:"value"`
	Type       string      `json:"type" yaml:"type"`
}

// ExperimentDefaults contains default experiment settings
type ExperimentDefaults struct {
	Duration               time.Duration `yaml:"duration"`
	TrafficPercent         float64       `yaml:"traffic_percent"`
	StatisticalSignificance float64      `yaml:"statistical_significance"`
	MinSampleSize          int           `yaml:"min_sample_size"`
}

// DefaultAdvancedConfig returns default advanced configuration
func DefaultAdvancedConfig() AdvancedConfig {
	return AdvancedConfig{
		ExperimentationEnabled:  true,
		ExperimentDefaults: ExperimentDefaults{
			Duration:                7 * 24 * time.Hour, // 1 week
			TrafficPercent:          10.0,               // 10%
			StatisticalSignificance: 0.95,               // 95%
			MinSampleSize:           1000,
		},
		StatisticalSignificance:    0.95,
		MinSampleSize:              1000,
		DefaultRolloutStrategy:     "percentage",
		SegmentationEnabled:        true,
		AnalyticsEnabled:           true,
		AnalyticsRetention:         30 * 24 * time.Hour, // 30 days
		DependencyTracking:         true,
		CacheEnabled:               true,
		CacheTTL:                   5 * time.Minute,
		EvaluationTimeout:          100 * time.Millisecond,
		KillSwitchEnabled:          true,
		CircuitBreakerEnabled:      true,
		MaxEvaluationsPerSecond:    10000,
		RolloutStrategies:          DefaultRolloutStrategies(),
		DefaultSegments:            DefaultUserSegments(),
	}
}

// DefaultRolloutStrategies returns default rollout strategies
func DefaultRolloutStrategies() []RolloutStrategy {
	return []RolloutStrategy{
		{
			ID:   "percentage_gradual",
			Name: "Gradual Percentage Rollout",
			Type: RolloutTypePercentage,
			Stages: []RolloutStage{
				{ID: "stage1", Name: "Initial", Percentage: 5.0, Duration: 24 * time.Hour, AutoAdvance: true},
				{ID: "stage2", Name: "Expand", Percentage: 25.0, Duration: 48 * time.Hour, AutoAdvance: true},
				{ID: "stage3", Name: "Majority", Percentage: 75.0, Duration: 72 * time.Hour, AutoAdvance: true},
				{ID: "stage4", Name: "Full", Percentage: 100.0, Duration: 0, AutoAdvance: false},
			},
			SafetyChecks: []SafetyCheck{
				{Type: "error_rate", Metric: "error_rate", Threshold: 5.0, Operator: "<", Window: time.Hour, Action: "pause"},
				{Type: "latency", Metric: "p95_latency", Threshold: 2000, Operator: "<", Window: 30 * time.Minute, Action: "pause"},
			},
			RollbackStrategy: RollbackStrategy{
				Type:      "immediate",
				Automatic: true,
				Triggers:  []string{"safety_check_failure", "manual"},
			},
		},
		{
			ID:   "canary_deployment",
			Name: "Canary Deployment",
			Type: RolloutTypeCanary,
			Stages: []RolloutStage{
				{ID: "canary", Name: "Canary", Percentage: 1.0, Duration: 2 * time.Hour, AutoAdvance: false},
				{ID: "production", Name: "Production", Percentage: 100.0, Duration: 0, AutoAdvance: false},
			},
		},
		{
			ID:   "ring_deployment",
			Name: "Ring-based Deployment",
			Type: RolloutTypeRing,
			Stages: []RolloutStage{
				{ID: "ring1", Name: "Internal Users", Percentage: 100.0, Duration: 12 * time.Hour, AutoAdvance: true},
				{ID: "ring2", Name: "Beta Users", Percentage: 100.0, Duration: 24 * time.Hour, AutoAdvance: true},
				{ID: "ring3", Name: "All Users", Percentage: 100.0, Duration: 0, AutoAdvance: false},
			},
		},
	}
}

// DefaultUserSegments returns default user segments
func DefaultUserSegments() []UserSegment {
	return []UserSegment{
		{
			ID:          "internal_users",
			Name:        "Internal Users",
			Description: "Internal company users",
			Rules: []SegmentRule{
				{Attribute: "email", Operator: "ends_with", Value: "@company.com", Type: "string"},
			},
			Tags: []string{"internal", "staff"},
		},
		{
			ID:          "beta_users",
			Name:        "Beta Users",
			Description: "Users enrolled in beta program",
			Rules: []SegmentRule{
				{Attribute: "beta_user", Operator: "equals", Value: true, Type: "boolean"},
			},
			Tags: []string{"beta", "early_adopters"},
		},
		{
			ID:          "premium_users",
			Name:        "Premium Users",
			Description: "Users with premium subscription",
			Rules: []SegmentRule{
				{Attribute: "subscription_tier", Operator: "in", Value: []string{"premium", "enterprise"}, Type: "string"},
			},
			Tags: []string{"premium", "paid"},
		},
	}
}

// NewAdvancedFlagManager creates a new advanced flag manager
func NewAdvancedFlagManager(
	basicManager *FlagManager,
	config AdvancedConfig,
	logger logger.Logger,
	telemetry *observability.TelemetryService,
) (*AdvancedFlagManager, error) {
	
	ctx, cancel := context.WithCancel(context.Background())
	
	manager := &AdvancedFlagManager{
		flagManager:       basicManager,
		config:            config,
		logger:            logger,
		telemetry:         telemetry,
		experiments:       make(map[string]*Experiment),
		segments:          make(map[string]*UserSegment),
		rolloutStrategies: make(map[string]*RolloutStrategy),
		flagDependencies:  make(map[string][]string),
		ctx:               ctx,
		cancel:            cancel,
	}
	
	// Initialize analytics
	if config.AnalyticsEnabled {
		manager.analytics = NewFeatureAnalytics(config, logger, telemetry)
	}
	
	// Initialize A/B testing engine
	if config.ExperimentationEnabled {
		manager.abTesting = NewABTestingEngine(config, logger, telemetry)
	}
	
	// Initialize default segments
	for i := range config.DefaultSegments {
		segment := &config.DefaultSegments[i]
		segment.CreatedAt = time.Now()
		segment.UpdatedAt = time.Now()
		manager.segments[segment.ID] = segment
	}
	
	// Initialize rollout strategies
	for i := range config.RolloutStrategies {
		strategy := &config.RolloutStrategies[i]
		manager.rolloutStrategies[strategy.ID] = strategy
	}
	
	// Start background tasks
	manager.startBackgroundTasks()
	
	logger.Info("Advanced flag manager initialized",
		"experimentation_enabled", config.ExperimentationEnabled,
		"segmentation_enabled", config.SegmentationEnabled,
		"analytics_enabled", config.AnalyticsEnabled,
		"segments_count", len(manager.segments),
		"rollout_strategies_count", len(manager.rolloutStrategies),
	)
	
	return manager, nil
}

// EvaluateFlag evaluates a feature flag with advanced targeting
func (afm *AdvancedFlagManager) EvaluateFlag(ctx context.Context, flagKey string, userContext UserContext, defaultValue interface{}) interface{} {
	start := time.Now()
	defer func() {
		// Record evaluation metrics
		if afm.telemetry != nil {
			afm.telemetry.RecordHistogram("feature_flag_evaluation_duration", float64(time.Since(start).Milliseconds()), map[string]string{
				"flag_key": flagKey,
			})
		}
	}()
	
	// Check if flag has an active experiment
	if experiment := afm.getActiveExperimentForFlag(flagKey); experiment != nil {
		if variant := afm.getExperimentVariant(experiment, userContext); variant != nil {
			// Record experiment exposure
			afm.recordExperimentExposure(experiment.ID, variant.ID, userContext)
			
			// Return variant configuration
			if value, exists := variant.Config["value"]; exists {
				return value
			}
		}
	}
	
	// Check rollout strategy
	if strategy := afm.getRolloutStrategyForFlag(flagKey); strategy != nil {
		if afm.shouldUserReceiveFlag(strategy, userContext) {
			// Get flag value from basic manager
			return afm.flagManager.EvaluateFlag(ctx, flagKey, userContext, defaultValue)
		}
		return defaultValue
	}
	
	// Fall back to basic evaluation
	return afm.flagManager.EvaluateFlag(ctx, flagKey, userContext, defaultValue)
}

// CreateExperiment creates a new A/B test experiment
func (afm *AdvancedFlagManager) CreateExperiment(experiment *Experiment) error {
	afm.mu.Lock()
	defer afm.mu.Unlock()
	
	// Validate experiment
	if err := afm.validateExperiment(experiment); err != nil {
		return fmt.Errorf("experiment validation failed: %w", err)
	}
	
	// Set timestamps
	experiment.StartTime = time.Now()
	if experiment.Duration > 0 {
		experiment.EndTime = experiment.StartTime.Add(experiment.Duration)
	}
	
	// Initialize results
	experiment.Results = &ExperimentResults{
		Status:         ResultStatusInProgress,
		VariantResults: make(map[string]VariantResult),
		StartTime:      experiment.StartTime,
	}
	
	// Store experiment
	afm.experiments[experiment.ID] = experiment
	
	afm.logger.Info("Experiment created",
		"experiment_id", experiment.ID,
		"experiment_name", experiment.Name,
		"variants_count", len(experiment.Variants),
		"traffic_percent", experiment.TrafficPercent,
	)
	
	return nil
}

// GetExperiment returns an experiment by ID
func (afm *AdvancedFlagManager) GetExperiment(experimentID string) (*Experiment, error) {
	afm.mu.RLock()
	defer afm.mu.RUnlock()
	
	experiment, exists := afm.experiments[experimentID]
	if !exists {
		return nil, fmt.Errorf("experiment not found: %s", experimentID)
	}
	
	// Return a copy to prevent external modifications
	experimentCopy := *experiment
	return &experimentCopy, nil
}

// UpdateExperiment updates an existing experiment
func (afm *AdvancedFlagManager) UpdateExperiment(experimentID string, updates *Experiment) error {
	afm.mu.Lock()
	defer afm.mu.Unlock()
	
	experiment, exists := afm.experiments[experimentID]
	if !exists {
		return fmt.Errorf("experiment not found: %s", experimentID)
	}
	
	// Only allow updates to certain fields based on status
	switch experiment.Status {
	case ExperimentStatusDraft:
		// All fields can be updated in draft
		*experiment = *updates
		experiment.ID = experimentID // Preserve ID
	case ExperimentStatusRunning:
		// Limited updates during running state
		experiment.Description = updates.Description
		experiment.EndTime = updates.EndTime
		// Cannot change variants or traffic allocation while running
	default:
		return fmt.Errorf("experiment cannot be updated in status: %s", experiment.Status)
	}
	
	return nil
}

// StartExperiment starts an experiment
func (afm *AdvancedFlagManager) StartExperiment(experimentID string) error {
	afm.mu.Lock()
	defer afm.mu.Unlock()
	
	experiment, exists := afm.experiments[experimentID]
	if !exists {
		return fmt.Errorf("experiment not found: %s", experimentID)
	}
	
	if experiment.Status != ExperimentStatusDraft && experiment.Status != ExperimentStatusScheduled {
		return fmt.Errorf("experiment cannot be started from status: %s", experiment.Status)
	}
	
	experiment.Status = ExperimentStatusRunning
	experiment.StartTime = time.Now()
	
	if experiment.Duration > 0 {
		experiment.EndTime = experiment.StartTime.Add(experiment.Duration)
	}
	
	afm.logger.Info("Experiment started",
		"experiment_id", experimentID,
		"experiment_name", experiment.Name,
	)
	
	return nil
}

// StopExperiment stops a running experiment
func (afm *AdvancedFlagManager) StopExperiment(experimentID string) error {
	afm.mu.Lock()
	defer afm.mu.Unlock()
	
	experiment, exists := afm.experiments[experimentID]
	if !exists {
		return fmt.Errorf("experiment not found: %s", experimentID)
	}
	
	if experiment.Status != ExperimentStatusRunning {
		return fmt.Errorf("experiment is not running: %s", experiment.Status)
	}
	
	experiment.Status = ExperimentStatusCompleted
	experiment.EndTime = time.Now()
	experiment.Results.EndTime = time.Now()
	experiment.Results.Duration = experiment.EndTime.Sub(experiment.StartTime)
	
	// Calculate final results
	afm.calculateExperimentResults(experiment)
	
	afm.logger.Info("Experiment stopped",
		"experiment_id", experimentID,
		"experiment_name", experiment.Name,
		"duration", experiment.Results.Duration,
	)
	
	return nil
}

// CreateUserSegment creates a new user segment
func (afm *AdvancedFlagManager) CreateUserSegment(segment *UserSegment) error {
	afm.mu.Lock()
	defer afm.mu.Unlock()
	
	segment.CreatedAt = time.Now()
	segment.UpdatedAt = time.Now()
	
	afm.segments[segment.ID] = segment
	
	afm.logger.Info("User segment created",
		"segment_id", segment.ID,
		"segment_name", segment.Name,
		"rules_count", len(segment.Rules),
	)
	
	return nil
}

// EvaluateUserSegment checks if a user belongs to a segment
func (afm *AdvancedFlagManager) EvaluateUserSegment(segmentID string, userContext UserContext) bool {
	afm.mu.RLock()
	segment, exists := afm.segments[segmentID]
	afm.mu.RUnlock()
	
	if !exists {
		return false
	}
	
	return afm.evaluateSegmentRules(segment.Rules, userContext)
}

// GetAnalytics returns feature analytics
func (afm *AdvancedFlagManager) GetAnalytics() *FeatureAnalytics {
	return afm.analytics
}

// GetABTestingEngine returns the A/B testing engine
func (afm *AdvancedFlagManager) GetABTestingEngine() *ABTestingEngine {
	return afm.abTesting
}

// Close gracefully shuts down the advanced flag manager
func (afm *AdvancedFlagManager) Close() error {
	afm.logger.Info("Shutting down advanced flag manager")
	
	// Cancel context and wait for background tasks
	afm.cancel()
	afm.wg.Wait()
	
	// Close analytics
	if afm.analytics != nil {
		afm.analytics.Close()
	}
	
	// Close A/B testing engine
	if afm.abTesting != nil {
		afm.abTesting.Close()
	}
	
	return nil
}

// Private methods

func (afm *AdvancedFlagManager) getActiveExperimentForFlag(flagKey string) *Experiment {
	afm.mu.RLock()
	defer afm.mu.RUnlock()
	
	for _, experiment := range afm.experiments {
		if experiment.Status == ExperimentStatusRunning {
			// Check if this experiment affects the flag
			if afm.experimentAffectsFlag(experiment, flagKey) {
				return experiment
			}
		}
	}
	
	return nil
}

func (afm *AdvancedFlagManager) experimentAffectsFlag(experiment *Experiment, flagKey string) bool {
	// This is simplified - in reality, you'd have explicit flag associations
	return strings.Contains(experiment.Name, flagKey) || strings.Contains(experiment.Description, flagKey)
}

func (afm *AdvancedFlagManager) getExperimentVariant(experiment *Experiment, userContext UserContext) *ExperimentVariant {
	// Check targeting rules first
	if !afm.evaluateTargetingRules(experiment.TargetingRules, userContext) {
		return nil
	}
	
	// Check user segments
	if !afm.userMatchesSegments(experiment.UserSegments, userContext) {
		return nil
	}
	
	// Check traffic allocation
	hash := afm.hashUser(experiment.ID, userContext.UserID)
	trafficBucket := hash % 100
	
	if float64(trafficBucket) >= experiment.TrafficPercent {
		return nil
	}
	
	// Assign variant based on weights
	return afm.assignVariant(experiment, userContext)
}

func (afm *AdvancedFlagManager) assignVariant(experiment *Experiment, userContext UserContext) *ExperimentVariant {
	// Calculate cumulative weights
	totalWeight := 0.0
	for _, variant := range experiment.Variants {
		totalWeight += variant.Weight
	}
	
	if totalWeight == 0 {
		return nil
	}
	
	// Hash user for consistent variant assignment
	hash := afm.hashUser(experiment.ID+"_variant", userContext.UserID)
	bucket := float64(hash%10000) / 100.0 // 0-100 with 2 decimal precision
	
	// Find variant based on bucket
	currentWeight := 0.0
	for _, variant := range experiment.Variants {
		currentWeight += (variant.Weight / totalWeight) * 100
		if bucket <= currentWeight {
			return &variant
		}
	}
	
	// Fallback to control variant
	for _, variant := range experiment.Variants {
		if variant.IsControl {
			return &variant
		}
	}
	
	// Fallback to first variant
	if len(experiment.Variants) > 0 {
		return &experiment.Variants[0]
	}
	
	return nil
}

func (afm *AdvancedFlagManager) hashUser(key, userID string) uint32 {
	data := fmt.Sprintf("%s:%s", key, userID)
	hash := md5.Sum([]byte(data))
	return uint32(hash[0])<<24 | uint32(hash[1])<<16 | uint32(hash[2])<<8 | uint32(hash[3])
}

func (afm *AdvancedFlagManager) evaluateTargetingRules(rules []TargetingRule, userContext UserContext) bool {
	if len(rules) == 0 {
		return true // No rules means everyone is targeted
	}
	
	for _, rule := range rules {
		if !afm.evaluateRule(rule, userContext) {
			return false
		}
	}
	
	return true
}

func (afm *AdvancedFlagManager) userMatchesSegments(segmentIDs []string, userContext UserContext) bool {
	if len(segmentIDs) == 0 {
		return true // No segments means all users match
	}
	
	for _, segmentID := range segmentIDs {
		if afm.EvaluateUserSegment(segmentID, userContext) {
			return true
		}
	}
	
	return false
}

func (afm *AdvancedFlagManager) evaluateSegmentRules(rules []SegmentRule, userContext UserContext) bool {
	if len(rules) == 0 {
		return true
	}
	
	for _, rule := range rules {
		if !afm.evaluateSegmentRule(rule, userContext) {
			return false
		}
	}
	
	return true
}

func (afm *AdvancedFlagManager) evaluateRule(rule TargetingRule, userContext UserContext) bool {
	userValue := afm.getUserAttribute(rule.Attribute, userContext)
	return afm.compareValues(userValue, rule.Operator, rule.Value, rule.Type)
}

func (afm *AdvancedFlagManager) evaluateSegmentRule(rule SegmentRule, userContext UserContext) bool {
	userValue := afm.getUserAttribute(rule.Attribute, userContext)
	return afm.compareValues(userValue, rule.Operator, rule.Value, rule.Type)
}

func (afm *AdvancedFlagManager) getUserAttribute(attribute string, userContext UserContext) interface{} {
	switch attribute {
	case "user_id":
		return userContext.UserID
	case "email":
		return userContext.Email
	case "country":
		return userContext.Country
	case "subscription_tier":
		return userContext.Attributes["subscription_tier"]
	case "beta_user":
		if val, exists := userContext.Attributes["beta_user"]; exists {
			if str, ok := val.(string); ok {
				return str == "true"
			}
			return val
		}
		return false
	default:
		return userContext.Attributes[attribute]
	}
}

func (afm *AdvancedFlagManager) compareValues(userValue interface{}, operator string, ruleValue interface{}, valueType string) bool {
	switch operator {
	case "equals", "==":
		return userValue == ruleValue
	case "not_equals", "!=":
		return userValue != ruleValue
	case "greater_than", ">":
		return afm.compareNumbers(userValue, ruleValue, ">")
	case "less_than", "<":
		return afm.compareNumbers(userValue, ruleValue, "<")
	case "greater_than_or_equal", ">=":
		return afm.compareNumbers(userValue, ruleValue, ">=")
	case "less_than_or_equal", "<=":
		return afm.compareNumbers(userValue, ruleValue, "<=")
	case "contains":
		return afm.stringContains(userValue, ruleValue)
	case "starts_with":
		return afm.stringStartsWith(userValue, ruleValue)
	case "ends_with":
		return afm.stringEndsWith(userValue, ruleValue)
	case "in":
		return afm.valueInList(userValue, ruleValue)
	case "not_in":
		return !afm.valueInList(userValue, ruleValue)
	default:
		return false
	}
}

func (afm *AdvancedFlagManager) compareNumbers(userValue, ruleValue interface{}, operator string) bool {
	userNum, err1 := afm.toFloat64(userValue)
	ruleNum, err2 := afm.toFloat64(ruleValue)
	
	if err1 != nil || err2 != nil {
		return false
	}
	
	switch operator {
	case ">":
		return userNum > ruleNum
	case "<":
		return userNum < ruleNum
	case ">=":
		return userNum >= ruleNum
	case "<=":
		return userNum <= ruleNum
	default:
		return false
	}
}

func (afm *AdvancedFlagManager) toFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("cannot convert to float64: %T", value)
	}
}

func (afm *AdvancedFlagManager) stringContains(userValue, ruleValue interface{}) bool {
	userStr := fmt.Sprintf("%v", userValue)
	ruleStr := fmt.Sprintf("%v", ruleValue)
	return strings.Contains(userStr, ruleStr)
}

func (afm *AdvancedFlagManager) stringStartsWith(userValue, ruleValue interface{}) bool {
	userStr := fmt.Sprintf("%v", userValue)
	ruleStr := fmt.Sprintf("%v", ruleValue)
	return strings.HasPrefix(userStr, ruleStr)
}

func (afm *AdvancedFlagManager) stringEndsWith(userValue, ruleValue interface{}) bool {
	userStr := fmt.Sprintf("%v", userValue)
	ruleStr := fmt.Sprintf("%v", ruleValue)
	return strings.HasSuffix(userStr, ruleStr)
}

func (afm *AdvancedFlagManager) valueInList(userValue, ruleValue interface{}) bool {
	list, ok := ruleValue.([]interface{})
	if !ok {
		// Try string slice
		if strList, ok := ruleValue.([]string); ok {
			userStr := fmt.Sprintf("%v", userValue)
			for _, item := range strList {
				if userStr == item {
					return true
				}
			}
		}
		return false
	}
	
	for _, item := range list {
		if userValue == item {
			return true
		}
	}
	
	return false
}

func (afm *AdvancedFlagManager) getRolloutStrategyForFlag(flagKey string) *RolloutStrategy {
	// This is simplified - in reality, you'd have explicit strategy associations
	afm.mu.RLock()
	defer afm.mu.RUnlock()
	
	// Return default strategy for demonstration
	return afm.rolloutStrategies[afm.config.DefaultRolloutStrategy]
}

func (afm *AdvancedFlagManager) shouldUserReceiveFlag(strategy *RolloutStrategy, userContext UserContext) bool {
	// Simplified rollout logic based on percentage
	if strategy.Type == RolloutTypePercentage && len(strategy.Stages) > 0 {
		// Get current stage
		currentStage := afm.getCurrentRolloutStage(strategy)
		if currentStage == nil {
			return false
		}
		
		// Hash user for consistent assignment
		hash := afm.hashUser("rollout_"+strategy.ID, userContext.UserID)
		bucket := hash % 100
		
		return float64(bucket) < currentStage.Percentage
	}
	
	return true
}

func (afm *AdvancedFlagManager) getCurrentRolloutStage(strategy *RolloutStrategy) *RolloutStage {
	// Simplified - return first stage
	if len(strategy.Stages) > 0 {
		return &strategy.Stages[0]
	}
	return nil
}

func (afm *AdvancedFlagManager) recordExperimentExposure(experimentID, variantID string, userContext UserContext) {
	if afm.analytics != nil {
		afm.analytics.RecordExposure(experimentID, variantID, userContext.UserID)
	}
	
	if afm.telemetry != nil {
		afm.telemetry.RecordCounter("experiment_exposures_total", 1, map[string]string{
			"experiment_id": experimentID,
			"variant_id":    variantID,
		})
	}
}

func (afm *AdvancedFlagManager) validateExperiment(experiment *Experiment) error {
	if experiment.ID == "" {
		return fmt.Errorf("experiment ID is required")
	}
	
	if experiment.Name == "" {
		return fmt.Errorf("experiment name is required")
	}
	
	if len(experiment.Variants) < 2 {
		return fmt.Errorf("experiment must have at least 2 variants")
	}
	
	if experiment.TrafficPercent < 0 || experiment.TrafficPercent > 100 {
		return fmt.Errorf("traffic percent must be between 0 and 100")
	}
	
	// Validate variant weights
	totalWeight := 0.0
	for _, variant := range experiment.Variants {
		if variant.Weight < 0 {
			return fmt.Errorf("variant weight cannot be negative")
		}
		totalWeight += variant.Weight
	}
	
	if totalWeight == 0 {
		return fmt.Errorf("total variant weight must be greater than 0")
	}
	
	return nil
}

func (afm *AdvancedFlagManager) calculateExperimentResults(experiment *Experiment) {
	// This is a simplified implementation
	// In reality, you'd perform comprehensive statistical analysis
	
	if afm.abTesting != nil {
		results := afm.abTesting.CalculateResults(experiment.ID)
		if results != nil {
			experiment.Results = results
		}
	}
}

func (afm *AdvancedFlagManager) startBackgroundTasks() {
	// Experiment monitoring task
	if afm.config.ExperimentationEnabled {
		afm.wg.Add(1)
		go afm.experimentMonitoringTask()
	}
	
	// Analytics collection task
	if afm.config.AnalyticsEnabled && afm.analytics != nil {
		afm.wg.Add(1)
		go afm.analyticsTask()
	}
	
	// Rollout monitoring task
	afm.wg.Add(1)
	go afm.rolloutMonitoringTask()
}

func (afm *AdvancedFlagManager) experimentMonitoringTask() {
	defer afm.wg.Done()
	
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-afm.ctx.Done():
			return
		case <-ticker.C:
			afm.monitorExperiments()
		}
	}
}

func (afm *AdvancedFlagManager) monitorExperiments() {
	afm.mu.RLock()
	experiments := make([]*Experiment, 0, len(afm.experiments))
	for _, exp := range afm.experiments {
		if exp.Status == ExperimentStatusRunning {
			experiments = append(experiments, exp)
		}
	}
	afm.mu.RUnlock()
	
	for _, experiment := range experiments {
		// Check if experiment should end
		if !experiment.EndTime.IsZero() && time.Now().After(experiment.EndTime) {
			afm.StopExperiment(experiment.ID)
		}
		
		// Calculate interim results
		afm.calculateExperimentResults(experiment)
		
		// Check for early stopping conditions
		if afm.shouldStopExperimentEarly(experiment) {
			afm.StopExperiment(experiment.ID)
		}
	}
}

func (afm *AdvancedFlagManager) shouldStopExperimentEarly(experiment *Experiment) bool {
	if experiment.Results == nil {
		return false
	}
	
	// Check for statistical significance
	if experiment.Results.SignificantResult && experiment.Results.Confidence >= afm.config.StatisticalSignificance {
		return true
	}
	
	// Check minimum sample size
	if experiment.Results.TotalSamples >= int64(afm.config.MinSampleSize) {
		return false // Let it run to completion
	}
	
	return false
}

func (afm *AdvancedFlagManager) analyticsTask() {
	defer afm.wg.Done()
	
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-afm.ctx.Done():
			return
		case <-ticker.C:
			if afm.analytics != nil {
				afm.analytics.ProcessMetrics()
			}
		}
	}
}

func (afm *AdvancedFlagManager) rolloutMonitoringTask() {
	defer afm.wg.Done()
	
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	
	for {
		select {
		case <-afm.ctx.Done():
			return
		case <-ticker.C:
			afm.monitorRollouts()
		}
	}
}

func (afm *AdvancedFlagManager) monitorRollouts() {
	// Monitor rollout strategies and advance stages if needed
	afm.mu.RLock()
	strategies := make([]*RolloutStrategy, 0, len(afm.rolloutStrategies))
	for _, strategy := range afm.rolloutStrategies {
		strategies = append(strategies, strategy)
	}
	afm.mu.RUnlock()
	
	for _, strategy := range strategies {
		afm.checkRolloutProgress(strategy)
	}
}

func (afm *AdvancedFlagManager) checkRolloutProgress(strategy *RolloutStrategy) {
	// Check safety checks
	for _, safetyCheck := range strategy.SafetyChecks {
		if afm.evaluateSafetyCheck(safetyCheck) {
			afm.logger.Warn("Safety check failed for rollout strategy",
				"strategy_id", strategy.ID,
				"safety_check", safetyCheck.Type,
			)
			
			// Take action based on safety check
			switch safetyCheck.Action {
			case "pause":
				// Pause rollout
			case "rollback":
				// Rollback rollout
			}
		}
	}
}

func (afm *AdvancedFlagManager) evaluateSafetyCheck(check SafetyCheck) bool {
	// This would integrate with your metrics system
	// For now, return false (no safety violations)
	return false
}