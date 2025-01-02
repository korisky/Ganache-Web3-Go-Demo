package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() (func(context.Context), error) {

	// new a exporter by constructor
	exporter, err := otlptracehttp.New(context.Background())
	if err != nil {
		return nil, err
	}

	// construct a tracer provider & set to global trace provider
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.Default()),
	)
	otel.SetTracerProvider(tracerProvider)

	// return the function -> for shutting down (clean up)
	return func(ctx context.Context) {
		_ = tracerProvider.Shutdown(ctx)
	}, nil
}
