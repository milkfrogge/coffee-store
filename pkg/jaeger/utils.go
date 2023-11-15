package jaeger

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

func ExtractMetaFromGRPC(ctx context.Context) (context.Context, error) {
	// Extract TraceID from header
	md, _ := metadata.FromIncomingContext(ctx)

	if len(md["x-trace-id"]) < 1 {
		return nil, errors.New("no metadata in grpc header")
	}

	traceIdString := md["x-trace-id"][0]
	// Convert string to byte array
	traceId, err := trace.TraceIDFromHex(traceIdString)
	if err != nil {
		return nil, err
	}
	// Creating a span context with a predefined trace-id
	spanContext := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	})
	// Embedding span config into the context
	ctx = trace.ContextWithSpanContext(ctx, spanContext)

	return ctx, nil
}
