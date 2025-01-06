package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	serviceName    = "demo-golang-agent"
	serviceVersion = "0.0.1"
	otlpEndpoint   = "localhost"
	environment    = "dev"
)

// initTracer ref: https://axiom.co/docs/guides/opentelemetry-go
func initTracer() (func(context.Context), error) {

	// new a exporter by constructor
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// construct a tracer provider & set to global trace provider
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(
			// attributes for service -> tags in Jaeger
			resource.NewWithAttributes(
				otlpEndpoint,
				attribute.String("serviceName", serviceName),
				attribute.String("serviceVersion", serviceVersion),
				attribute.String("environment", environment)),
		),
	)
	otel.SetTracerProvider(tracerProvider)

	// return the function -> for shutting down (clean up)
	return func(ctx context.Context) {
		_ = tracerProvider.Shutdown(ctx)
	}, nil
}
