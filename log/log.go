package log

import (
	"context"

	"go.uber.org/zap/zapcore"

	"github.com/KargoGlobal/go-zap/utils"
	trace "go.opencensus.io/trace"
	"go.uber.org/zap"
)

var (
	logger = zap.L()
)

// SetLogger sets the logger of the package
func SetLogger(log *zap.Logger) {
	logger = log
}

// DEBUG

// DebugWithContext logs on debug level and trace based on the context span if it exists
func DebugWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	DebugWithSpan(trace.FromContext(ctx), log, fields...)
}

// DebugWithSpan logs on debug level and add the logs on the trace if span exists.
func DebugWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	Debug(log, fields...)
	logSpan(span, log, fields...)
}

// Debug logs on debug level
func Debug(log string, fields ...zapcore.Field) {
	logger.Debug(log, fields...)
}

// INFO

// InfoWithContext logs on info level and trace based on the context span if it exists
func InfoWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	InfoWithSpan(trace.FromContext(ctx), log, fields...)
}

// InfoWithSpan logs on info level and add the logs on the trace if span exists.
func InfoWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	Info(log, fields...)
	logSpan(span, log, fields...)

}

// Info logs on info level
func Info(log string, fields ...zapcore.Field) {
	logger.Info(log, fields...)
}

// WARN

// WarnWithContext logs on warn level and trace based on the context span if it exists
func WarnWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	WarnWithSpan(trace.FromContext(ctx), log, fields...)
}

// WarnWithSpan logs on warn level and add the logs on the trace if span exists.
func WarnWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	Warn(log, fields...)
	logSpan(span, log, fields...)

}

// Warn logs on warn level
func Warn(log string, fields ...zapcore.Field) {
	logger.Warn(log, fields...)
}

// ERROR

// ErrorWithContext logs on error level and trace based on the context span if it exists
func ErrorWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	ErrorWithSpan(trace.FromContext(ctx), log, fields...)
}

// ErrorWithSpan logs on error level and add the logs on the trace if span exists.
func ErrorWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	Error(log, fields...)
	logSpan(span, log, fields...)
}

// Error logs on error level
func Error(log string, fields ...zapcore.Field) {
	logger.Error(log, fields...)
}

// DPANIC

// DPanicWithContext logs on dPanic level and trace based on the context span if it exists
func DPanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	DPanicWithSpan(trace.FromContext(ctx), log, fields...)
}

// DPanicWithSpan logs on dPanic level and add the logs on the trace if span exists.
func DPanicWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	DPanic(log, fields...)
}

// DPanic logs on dPanic level
func DPanic(log string, fields ...zapcore.Field) {
	logger.DPanic(log, fields...)
}

// PANIC

// PanicWithContext logs on panic level and trace based on the context span if it exists
func PanicWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	PanicWithSpan(trace.FromContext(ctx), log, fields...)
}

// PanicWithSpan logs on panic level and add the logs on the trace if span exists.
func PanicWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	Panic(log, fields...)
}

// Panic logs on panic level
func Panic(log string, fields ...zapcore.Field) {
	logger.Panic(log, fields...)
}

// FatalWithContext logs on fatal level and trace based on the context span if it exists
func FatalWithContext(ctx context.Context, log string, fields ...zapcore.Field) {
	FatalWithSpan(trace.FromContext(ctx), log, fields...)
}

// FatalWithSpan logs on fatal level and add the logs on the trace if span exists.
func FatalWithSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	logSpan(span, log, fields...)
	Fatal(log, fields...)
}

// Fatal logs on fatal level
func Fatal(log string, fields ...zapcore.Field) {
	logger.Fatal(log, fields...)
}

func logSpan(span *trace.Span, log string, fields ...zapcore.Field) {
	if span != nil {
		traceAttributes := []trace.Attribute{}
		if len(fields) > 0 {
			traceAttributes = append(traceAttributes, utils.ZapFieldsToOpenCensus(fields...)...)
		}
		span.Annotate(traceAttributes, log)
	}
}
