package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"time"
)

func main() {

	// init openTelemetry tracer
	cleanup, err := initTracer()
	if err != nil {
		log.Fatalf("Failed to init tracer: %v", err)
	}
	defer cleanup(context.Background())

	// web-service config
	router := gin.Default()

	// add tracing as a gin's middleware
	router.Use(func(c *gin.Context) {
		// start the tracer & defer the end
		tracer := otel.Tracer("agent-demo-tracer")
		ctx, span := tracer.Start(c.Request.Context(), c.FullPath())
		defer span.End()

		// pass the request with 'enhanced' context
		c.Request = c.Request.WithContext(ctx)
		c.Next() // -> similar to Netty's handler chain
	})

	// simple endpoint
	router.GET("/hello", func(c *gin.Context) {
		tracer := otel.Tracer("hello-tracer")
		_, span := tracer.Start(c.Request.Context(), "hello-handler")
		defer span.End()

		time.Sleep(50 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"msg": "hello, golang agent"})
	})

	// server start
	router.Run(":8080")
}
