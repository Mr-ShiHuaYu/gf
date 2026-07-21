package otel

import (
	"context"

	"github.com/Mr-ShiHuaYu/otel-go111/attribute"
	"github.com/Mr-ShiHuaYu/otel-go111/codes"
	"github.com/Mr-ShiHuaYu/otel-go111/propagation"
	"github.com/Mr-ShiHuaYu/otel-go111/trace"
)

var (
	tracerProvider    trace.TracerProvider          = &defaultTracerProvider{}
	textMapPropagator propagation.TextMapPropagator = propagation.NewCompositeTextMapPropagator()
)

type defaultTracerProvider struct{}

func (p *defaultTracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return &defaultTracer{}
}

type defaultTracer struct{}

func (t *defaultTracer) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return ctx, &defaultSpan{}
}

type defaultSpan struct{}

func (s *defaultSpan) End(opts ...trace.SpanEndOption) {}

func (s *defaultSpan) AddEvent(name string, opts ...trace.EventOption) {}

func (s *defaultSpan) IsRecording() bool {
	return false
}

func (s *defaultSpan) RecordError(err error, opts ...trace.EventOption) {}

func (s *defaultSpan) SetStatus(code codes.Code, description string) {}

func (s *defaultSpan) SetName(name string) {}

func (s *defaultSpan) SetAttributes(kv ...attribute.KeyValue) {}

func (s *defaultSpan) SpanContext() trace.SpanContext {
	return trace.SpanContext{}
}

func (s *defaultSpan) TracerProvider() trace.TracerProvider {
	return nil
}

func SetTracerProvider(tp trace.TracerProvider) {
	tracerProvider = tp
}

func GetTracerProvider() trace.TracerProvider {
	return tracerProvider
}

func Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return tracerProvider.Tracer(name, opts...)
}

func SetTextMapPropagator(p propagation.TextMapPropagator) {
	textMapPropagator = p
}

func GetTextMapPropagator() propagation.TextMapPropagator {
	return textMapPropagator
}
