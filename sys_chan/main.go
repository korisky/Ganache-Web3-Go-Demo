package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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
