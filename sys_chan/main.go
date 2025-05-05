package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	_shutdownPeriod      = 15 * time.Second
	_shutdownHardPeriod  = 3 * time.Second
	_readinessDrainDelay = 5 * time.Second
)

var isShuttingDown atomic.Bool

// main a graceful full http app example
func main() {
	// 1. setup signal ctx
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// 2. readiness endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if isShuttingDown.Load() {
			http.Error(w, "Server shutting down", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintln(w, "OK")
	})

	// sample business logic
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Fprintln(w, "Hello Graceful Server")
		case <-r.Context().Done():
			http.Error(w, "Request cancelled", http.StatusRequestTimeout)
		}
	})

	// 3. ensure in-flight requests aren't cancelled immediately
	ongoingCtx, stopOngoingGracefully := context.WithCancel(context.Background())
	server := &http.Server{
		Addr: ":8099",
		BaseContext: func(_ net.Listener) context.Context {
			return ongoingCtx
		},
	}

	// server start by another routine
	go func() {
		log.Println("Server starting on :8099")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServer: %v", err)
		}
	}()

	// 4. waiting for shutdown signal
	<-rootCtx.Done()
	stop()
	isShuttingDown.Store(true)
	log.Println("Received shutdown signal, shutting down gracefully")

	// 5. give time for readiness check to propagate
	time.Sleep(_readinessDrainDelay)
	log.Println("Readiness check propagated, waiting for ongoing requests")

	// 6. waiting ongoing request consumption
	shutdownCtx, cancel := context.WithTimeout(context.Background(), _shutdownPeriod)
	defer cancel()
	err := server.Shutdown(shutdownCtx)
	stopOngoingGracefully()

	if err != nil {
		log.Println("Failed to wait for ongoing requests to finish, waiting for forced cancellation")
		time.Sleep(_shutdownHardPeriod)
	}
	log.Println("Server shutdown gracefully")
}

// BindIntTerSigHandlingWithCtx after go ver 1.16, OS signal handling could bind with ctx
func BindIntTerSigHandlingWithCtx() {
	// bind notify with context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	fmt.Println("Server running...")

	// tasks here

	<-ctx.Done()
	fmt.Println("\nReceived termination signal, shutting down...")
	stop()
}

// readinessHandler 为了gracefully关闭容器, 提前告知probe探针服务不可用,
// 使得k8s将流量不再打入, 然后再进行服务的OS-Signal处理的结束, 会更顺畅的关闭
func readinessHandler(w http.ResponseWriter, r *http.Request) {

	// shutting down -> let the K8S probe failed first
	if isShuttingDown.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("service is shutting down"))
		return
	}

	// normal stage -> just ok
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// PureIntTerSigHandling before go ver 1.16, purely handle SIG_INT & SIG_TERM
func PureIntTerSigHandling() {
	// buffered channel with 1 capacity is usually a good choice
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("Server running...")

	// tasks here

	// hangUp waiting for interrupt / termination OS signal
	<-signalChan

	// for handling second Ctrl+C
	signal.Stop(signalChan)
	fmt.Println("\nReceived termination signal, shutting down...")
}
