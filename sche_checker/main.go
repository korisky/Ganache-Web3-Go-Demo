package main

import (
	"log"
	"net/http"
	"time"
)

// client
var client = &http.Client{
	Timeout: 30 * time.Second,
}

func main() {

}

// checkHttpRespCode200 only check http response code is 200 or not
// not care the resp body
func checkHttpRespCode200(url string) bool {

	// request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Request failed with %v\n", err)
		return false
	}

	// only check http-resp code
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with %v\n", resp.StatusCode)
		return false
	}
	return true
}
