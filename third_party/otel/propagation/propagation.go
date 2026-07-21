package propagation

import (
	"context"
)

type TextMapCarrier interface {
	Get(key string) string
	Set(key, value string)
	Keys() []string
}

type TextMapPropagator interface {
	Inject(ctx context.Context, carrier TextMapCarrier)
	Extract(ctx context.Context, carrier TextMapCarrier) context.Context
	Fields() []string
}

type TraceContext struct{}

func NewTraceContext() TraceContext {
	return TraceContext{}
}

func (tc TraceContext) Inject(ctx context.Context, carrier TextMapCarrier) {}

func (tc TraceContext) Extract(ctx context.Context, carrier TextMapCarrier) context.Context {
	return ctx
}

func (tc TraceContext) Fields() []string {
	return []string{}
}

type Baggage struct{}

func NewBaggagePropagator() Baggage {
	return Baggage{}
}

func (b Baggage) Inject(ctx context.Context, carrier TextMapCarrier) {}

func (b Baggage) Extract(ctx context.Context, carrier TextMapCarrier) context.Context {
	return ctx
}

func (b Baggage) Fields() []string {
	return []string{}
}

type compositeTextMapPropagator struct {
	props []TextMapPropagator
}

func NewCompositeTextMapPropagator(props ...TextMapPropagator) TextMapPropagator {
	return &compositeTextMapPropagator{props: props}
}

func (p *compositeTextMapPropagator) Inject(ctx context.Context, carrier TextMapCarrier) {
	for _, prop := range p.props {
		prop.Inject(ctx, carrier)
	}
}

func (p *compositeTextMapPropagator) Extract(ctx context.Context, carrier TextMapCarrier) context.Context {
	for _, prop := range p.props {
		ctx = prop.Extract(ctx, carrier)
	}
	return ctx
}

func (p *compositeTextMapPropagator) Fields() []string {
	var fields []string
	for _, prop := range p.props {
		fields = append(fields, prop.Fields()...)
	}
	return fields
}

func MapCarrier(carrier map[string]string) TextMapCarrier {
	return mapCarrier(carrier)
}

type mapCarrier map[string]string

func (m mapCarrier) Get(key string) string {
	return m[key]
}

func (m mapCarrier) Set(key, value string) {
	m[key] = value
}

func (m mapCarrier) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func HeaderCarrier(carrier map[string][]string) TextMapCarrier {
	return headerCarrier(carrier)
}

type headerCarrier map[string][]string

func (h headerCarrier) Get(key string) string {
	if vals, ok := h[key]; ok && len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (h headerCarrier) Set(key, value string) {
	h[key] = []string{value}
}

func (h headerCarrier) Keys() []string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	return keys
}
