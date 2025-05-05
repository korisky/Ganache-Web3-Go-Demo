package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var isShuttingDown atomic.Bool

func main() {

	// Before go 1.16
	//PureIntTerSigHandling()

	// After go 1.16, ties the signal handling with Ctx
	BindIntTerSigHandlingWithCtx()

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
