package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
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
