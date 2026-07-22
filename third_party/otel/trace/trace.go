package trace

import (
	"context"
	"encoding/hex"

	"github.com/Mr-ShiHuaYu/gf/v2/third_party/otel/attribute"
	"github.com/Mr-ShiHuaYu/gf/v2/third_party/otel/codes"
)

type TraceID [16]byte

func (t TraceID) IsValid() bool {
	return t != TraceID{}
}

func (t TraceID) String() string {
	return hex.EncodeToString(t[:])
}

func (t TraceID) HexString() string {
	return t.String()
}

func TraceIDFromHex(h string) (TraceID, error) {
	var id TraceID
	b, err := hex.DecodeString(h)
	if err != nil {
		return id, err
	}
	copy(id[:], b)
	return id, nil
}

type SpanID [8]byte

func (s SpanID) IsValid() bool {
	return s != SpanID{}
}

func (s SpanID) String() string {
	return hex.EncodeToString(s[:])
}

func (s SpanID) HexString() string {
	return s.String()
}

type TraceFlags byte

const (
	FlagsSampled TraceFlags = 1
)

func (tf TraceFlags) IsSampled() bool {
	return tf&FlagsSampled != 0
}

type TraceState struct{}

func (ts TraceState) String() string {
	return ""
}

func (ts TraceState) Get(key string) (string, bool) {
	return "", false
}

func (ts TraceState) Insert(key, value string) (TraceState, error) {
	return ts, nil
}

func (ts TraceState) Update(key, value string) (TraceState, error) {
	return ts, nil
}

func (ts TraceState) Delete(key string) (TraceState, error) {
	return ts, nil
}

type SpanContextConfig struct {
	TraceID    TraceID
	SpanID     SpanID
	TraceFlags TraceFlags
	TraceState TraceState
	Remote     bool
}

func NewSpanContext(cfg SpanContextConfig) SpanContext {
	return SpanContext{
		traceID:    cfg.TraceID,
		spanID:     cfg.SpanID,
		traceFlags: cfg.TraceFlags,
		remote:     cfg.Remote,
	}
}

type SpanContext struct {
	traceID    TraceID
	spanID     SpanID
	traceFlags TraceFlags
	remote     bool
}

func (sc SpanContext) TraceID() TraceID {
	return sc.traceID
}

func (sc SpanContext) SpanID() SpanID {
	return sc.spanID
}

func (sc SpanContext) TraceFlags() TraceFlags {
	return sc.traceFlags
}

func (sc SpanContext) TraceState() TraceState {
	return TraceState{}
}

func (sc SpanContext) IsSampled() bool {
	return sc.traceFlags.IsSampled()
}

func (sc SpanContext) IsValid() bool {
	return sc.traceID.IsValid() && sc.spanID.IsValid()
}

func (sc SpanContext) HasTraceID() bool {
	return sc.traceID.IsValid()
}

func (sc SpanContext) HasSpanID() bool {
	return sc.spanID.IsValid()
}

func (sc SpanContext) IsRemote() bool {
	return sc.remote
}

type spanContextKeyType struct{}

var spanContextKey = spanContextKeyType{}

func ContextWithSpanContext(ctx context.Context, sc SpanContext) context.Context {
	return context.WithValue(ctx, spanContextKey, sc)
}

func SpanContextFromContext(ctx context.Context) SpanContext {
	if ctx == nil {
		return SpanContext{}
	}
	sc, ok := ctx.Value(spanContextKey).(SpanContext)
	if !ok {
		return SpanContext{}
	}
	return sc
}

func ContextWithRemoteSpanContext(ctx context.Context, sc SpanContext) context.Context {
	sc.remote = true
	return ContextWithSpanContext(ctx, sc)
}

type Span interface {
	End(options ...SpanEndOption)
	AddEvent(name string, options ...EventOption)
	IsRecording() bool
	RecordError(err error, options ...EventOption)
	SetStatus(code codes.Code, description string)
	SetName(name string)
	SetAttributes(kv ...attribute.KeyValue)
	SpanContext() SpanContext
	TracerProvider() TracerProvider
}

type noopSpan struct{}

func (noopSpan) End(...SpanEndOption)                {}
func (noopSpan) AddEvent(string, ...EventOption)     {}
func (noopSpan) IsRecording() bool                   { return false }
func (noopSpan) RecordError(error, ...EventOption)   {}
func (noopSpan) SetStatus(codes.Code, string)        {}
func (noopSpan) SetName(string)                      {}
func (noopSpan) SetAttributes(...attribute.KeyValue) {}
func (noopSpan) SpanContext() SpanContext            { return SpanContext{} }
func (noopSpan) TracerProvider() TracerProvider      { return nil }

type Tracer interface {
	Start(ctx context.Context, spanName string, opts ...SpanStartOption) (context.Context, Span)
}

type noopTracer struct{}

func (noopTracer) Start(ctx context.Context, name string, opts ...SpanStartOption) (context.Context, Span) {
	return ctx, noopSpan{}
}

type TracerProvider interface {
	Tracer(instrumentationName string, opts ...TracerOption) Tracer
}

type noopTracerProvider struct{}

func (noopTracerProvider) Tracer(string, ...TracerOption) Tracer {
	return noopTracer{}
}

type SpanStartOption interface {
	applySpanStart(*SpanConfig)
}

type SpanEndOption interface {
	applySpanEnd(*SpanEndConfig)
}

type EventOption interface {
	applyEvent(*EventConfig)
}

type TracerOption interface {
	applyTracer(*TracerConfig)
}

type SpanConfig struct {
	Attributes []attribute.KeyValue
	Timestamp  interface{}
	Links      []Link
	NewRoot    bool
	SpanKind   SpanKind
}

type SpanEndConfig struct {
	Timestamp  interface{}
	StackTrace bool
}

type EventConfig struct {
	Attributes []attribute.KeyValue
	Timestamp  interface{}
}

type TracerConfig struct {
	InstrumentationVersion string
	SchemaURL              string
}

type SpanKind int

const (
	SpanKindUnspecified SpanKind = iota
	SpanKindInternal
	SpanKindServer
	SpanKindClient
	SpanKindProducer
	SpanKindConsumer
)

func (sk SpanKind) String() string {
	switch sk {
	case SpanKindInternal:
		return "internal"
	case SpanKindServer:
		return "server"
	case SpanKindClient:
		return "client"
	case SpanKindProducer:
		return "producer"
	case SpanKindConsumer:
		return "consumer"
	}
	return "unspecified"
}

type Link struct {
	SpanContext
	Attributes []attribute.KeyValue
}

func WithAttributes(attributes ...attribute.KeyValue) withAttributes {
	return withAttributes{attrs: attributes}
}

type withAttributes struct {
	attrs []attribute.KeyValue
}

func (o withAttributes) applySpanStart(c *SpanConfig) {
	c.Attributes = append(c.Attributes, o.attrs...)
}

func (o withAttributes) applyEvent(c *EventConfig) {
	c.Attributes = append(c.Attributes, o.attrs...)
}

func (o withAttributes) applySpanEnd(c *SpanEndConfig) {
}

func WithTimestamp(t interface{}) SpanStartOption {
	return nil
}

func WithLinks(links ...Link) SpanStartOption {
	return nil
}

func WithNewRoot() SpanStartOption {
	return nil
}

func WithSpanKind(kind SpanKind) SpanStartOption {
	return nil
}

func WithEventAttributes(attrs ...attribute.KeyValue) EventOption {
	return WithAttributes(attrs...)
}

func WithStackTrace(b bool) SpanEndOption {
	return nil
}

func WithTimestampEnd(t interface{}) SpanEndOption {
	return nil
}

func WithInstrumentationVersion(version string) TracerOption {
	return withInstrumentationVersion(version)
}

type withInstrumentationVersion string

func (w withInstrumentationVersion) applyTracer(cfg *TracerConfig) {
	cfg.InstrumentationVersion = string(w)
}
