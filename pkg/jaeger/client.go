package jaeger

import (
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

type Tracer struct {
	url string
}

func (t *Tracer) NewTracerProvider(name string) (*tracesdk.TracerProvider, error) {

	// Create the Jaeger exporter
	exp, err := jaegerExporter.New(jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(t.url)))
	if err != nil {
		return nil, err
	}
	tracer := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	)
	return tracer, nil
}

func NewTracer(url string) *Tracer {
	return &Tracer{url: url}
}
