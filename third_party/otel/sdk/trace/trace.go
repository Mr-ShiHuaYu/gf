package trace

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type TracerProvider struct {
	idGenerator IDGenerator
}

type IDGenerator interface {
	NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID)
	NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID
}

type TracerProviderOption func(*TracerProvider)

func WithIDGenerator(g IDGenerator) TracerProviderOption {
	return func(tp *TracerProvider) {
		tp.idGenerator = g
	}
}

func NewTracerProvider(opts ...TracerProviderOption) *TracerProvider {
	tp := &TracerProvider{}
	for _, opt := range opts {
		opt(tp)
	}
	return tp
}

func (tp *TracerProvider) Tracer(name string, opts ...trace.TracerOption) trace.Tracer {
	return &tracerImpl{tp: tp}
}

type tracerImpl struct {
	tp *TracerProvider
}

func (t *tracerImpl) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return ctx, &spanImpl{}
}

type spanImpl struct{}

func (s *spanImpl) End(opts ...trace.SpanEndOption) {}

func (s *spanImpl) AddEvent(name string, opts ...trace.EventOption) {}

func (s *spanImpl) IsRecording() bool {
	return false
}

func (s *spanImpl) RecordError(err error, opts ...trace.EventOption) {}

func (s *spanImpl) SetStatus(code codes.Code, description string) {}

func (s *spanImpl) SetName(name string) {}

func (s *spanImpl) SetAttributes(kv ...attribute.KeyValue) {}

func (s *spanImpl) SpanContext() trace.SpanContext {
	return trace.SpanContext{}
}

func (s *spanImpl) TracerProvider() trace.TracerProvider {
	return nil
}

type Sampler interface{}

type SpanProcessor interface{}
