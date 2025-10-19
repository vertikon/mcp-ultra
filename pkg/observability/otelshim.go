package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// Tracer reexport: centraliza criação de tracer via facade.
func Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return otel.Tracer(name, opts...)
}

// KeyValue re-export para uso sem importar attribute diretamente
type KeyValue = attribute.KeyValue

// Atributos (helpers)
func AttrString(key, val string) KeyValue          { return attribute.String(key, val) }
func AttrInt(key string, val int) KeyValue         { return attribute.Int(key, val) }
func AttrInt64(key string, val int64) KeyValue     { return attribute.Int64(key, val) }
func AttrFloat64(key string, val float64) KeyValue { return attribute.Float64(key, val) }
func AttrBool(key string, val bool) KeyValue       { return attribute.Bool(key, val) }

// Baggage (helpers mínimos usados pelo projeto)
func BaggageFromContext(ctx context.Context) baggage.Baggage {
	return baggage.FromContext(ctx)
}

func BaggageParse(s string) (baggage.Baggage, error) {
	return baggage.Parse(s)
}

func BaggageNewMember(key, value string) (baggage.Member, error) {
	return baggage.NewMember(key, value)
}

func BaggageContextWithBaggage(ctx context.Context, bag baggage.Baggage) context.Context {
	return baggage.ContextWithBaggage(ctx, bag)
}

// Type re-exports for trace
type (
	TracerType   = trace.Tracer
	Span         = trace.Span
	SpanContext  = trace.SpanContext
	TracerOption = trace.TracerOption
	SpanKind     = trace.SpanKind
)

// Span status codes
var (
	StatusCodeUnset = codes.Unset
	StatusCodeError = codes.Error
	StatusCodeOK    = codes.Ok
)

// Span kinds
const (
	SpanKindInternal = trace.SpanKindInternal
	SpanKindServer   = trace.SpanKindServer
	SpanKindClient   = trace.SpanKindClient
	SpanKindProducer = trace.SpanKindProducer
	SpanKindConsumer = trace.SpanKindConsumer
)

// Metric types
type (
	MeterType              = metric.Meter
	MeterOption            = metric.MeterOption
	Int64Counter           = metric.Int64Counter
	Float64Counter         = metric.Float64Counter
	Int64Histogram         = metric.Int64Histogram
	Float64Histogram       = metric.Float64Histogram
	Int64ObservableGauge   = metric.Int64ObservableGauge
	Float64ObservableGauge = metric.Float64ObservableGauge
)

// GetMeter creates a new meter
func GetMeter(name string, opts ...metric.MeterOption) metric.Meter {
	return otel.Meter(name, opts...)
}

// GetTracerProvider returns the global TracerProvider
func GetTracerProvider() trace.TracerProvider {
	return otel.GetTracerProvider()
}

// SetTracerProvider sets the global TracerProvider
func SetTracerProvider(provider trace.TracerProvider) {
	otel.SetTracerProvider(provider)
}

// GetMeterProvider returns the global MeterProvider
func GetMeterProvider() metric.MeterProvider {
	return otel.GetMeterProvider()
}

// SetMeterProvider sets the global MeterProvider
func SetMeterProvider(provider metric.MeterProvider) {
	otel.SetMeterProvider(provider)
}
