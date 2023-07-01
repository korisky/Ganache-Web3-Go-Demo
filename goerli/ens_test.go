package main

import (
	"github.com/wealdtech/go-ens"
	"log"
	"testing"
)

// Test_EthereumNameService for testing the ens name resolve
func Test_EthereumNameService(t *testing.T) {
	address, err := ens.Resolve(client, "bsc.eth")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(address)
}
