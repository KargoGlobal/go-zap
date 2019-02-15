package utils

import (
	"fmt"

	trace "go.opencensus.io/trace"
	"go.uber.org/zap/zapcore"
)

// ZapFieldsToOpenCensus returns a table of standard opentracing field based on
// the inputed table of Zap field.
func ZapFieldsToOpenCensus(zapFields ...zapcore.Field) []trace.Attribute {
	traceAttributes := []trace.Attribute{}

	for _, zapField := range zapFields {
		attribute, err := ZapFieldToOpenCensus(zapField)
		if err == nil {
			traceAttributes = append(traceAttributes, attribute)
		}
	}

	return traceAttributes
}

// ZapFieldToOpenCensus returns a standard opentracing field based on the
// input Zap field.
func ZapFieldToOpenCensus(zapField zapcore.Field) (trace.Attribute, error) {
	switch zapField.Type {

	case zapcore.BoolType:
		val := false
		if zapField.Integer >= 1 {
			val = true
		}
		return trace.BoolAttribute(zapField.Key, val), nil
	case zapcore.Int64Type:
		return trace.Int64Attribute(zapField.Key, int64(zapField.Integer)), nil
	case zapcore.StringType:
		return trace.StringAttribute(zapField.Key, zapField.String), nil
	default:
		return trace.Attribute{}, fmt.Errorf("invalid zap attribue type")
	}
}
