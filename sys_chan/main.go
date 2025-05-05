package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// buffered channel with 1 capacity is usually a good choice
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("Server running...")

	// hangUp waiting for interrupt / termination OS signal
	<-signalChan
	fmt.Println("\nReceived termination signal, shutting down...")
}
