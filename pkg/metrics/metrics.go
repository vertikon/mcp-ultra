package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Re-export prometheus types and functions
type (
	Counter       = prometheus.Counter
	CounterVec    = prometheus.CounterVec
	CounterOpts   = prometheus.CounterOpts
	Gauge         = prometheus.Gauge
	GaugeVec      = prometheus.GaugeVec
	GaugeOpts     = prometheus.GaugeOpts
	Histogram     = prometheus.Histogram
	HistogramVec  = prometheus.HistogramVec
	HistogramOpts = prometheus.HistogramOpts
	Summary       = prometheus.Summary
	SummaryVec    = prometheus.SummaryVec
	SummaryOpts   = prometheus.SummaryOpts
	Registry      = prometheus.Registry
	Registerer    = prometheus.Registerer
	Gatherer      = prometheus.Gatherer
	Collector     = prometheus.Collector
)

// Re-export constructor functions
var (
	NewCounter        = prometheus.NewCounter
	NewCounterVec     = prometheus.NewCounterVec
	NewGauge          = prometheus.NewGauge
	NewGaugeVec       = prometheus.NewGaugeVec
	NewHistogram      = prometheus.NewHistogram
	NewHistogramVec   = prometheus.NewHistogramVec
	NewSummary        = prometheus.NewSummary
	NewSummaryVec     = prometheus.NewSummaryVec
	NewRegistry       = prometheus.NewRegistry
	MustRegister      = prometheus.MustRegister
	Register          = prometheus.Register
	Unregister        = prometheus.Unregister
	DefaultRegisterer = prometheus.DefaultRegisterer
	DefaultGatherer   = prometheus.DefaultGatherer
)

// Re-export promauto functions
var (
	NewCounterFunc      = promauto.NewCounterFunc
	NewGaugeFunc        = promauto.NewGaugeFunc
	NewCounterAuto      = promauto.NewCounter
	NewCounterVecAuto   = promauto.NewCounterVec
	NewGaugeAuto        = promauto.NewGauge
	NewGaugeVecAuto     = promauto.NewGaugeVec
	NewHistogramAuto    = promauto.NewHistogram
	NewHistogramVecAuto = promauto.NewHistogramVec
	With                = promauto.With
)

// Re-export promhttp functions
var (
	Handler                   = promhttp.Handler
	HandlerFor                = promhttp.HandlerFor
	InstrumentHandlerDuration = promhttp.InstrumentHandlerDuration
	InstrumentHandlerCounter  = promhttp.InstrumentHandlerCounter
	InstrumentHandlerInFlight = promhttp.InstrumentHandlerInFlight
)

// Utility functions
var (
	DefBuckets              = prometheus.DefBuckets
	LinearBuckets           = prometheus.LinearBuckets
	ExponentialBuckets      = prometheus.ExponentialBuckets
	ExponentialBucketsRange = prometheus.ExponentialBucketsRange
)
