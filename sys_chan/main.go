package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// before go 1.16
	// buffered channel with 1 capacity is usually a good choice
	PureIntTerSigHandling()

	// after go 1.16, ties the signal handling with Ctx

}

// PureIntTerSigHandling before go ver 1.16, purely handle SIG_INT & SIG_TERM
func PureIntTerSigHandling() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("Server running...")

	// hangUp waiting for interrupt / termination OS signal
	<-signalChan

	// for handling second Ctrl+C
	signal.Stop(signalChan)
	fmt.Println("\nReceived termination signal, shutting down...")
}
