package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// sendAlert is for sending alert for specific url not response 200
func sendAlert(webhookUrl, monitoredUrl string) {

	// payload -> request body
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("Alert: URL check failed for %s", monitoredUrl),
		},
	}
	payLoadBytes, _ := json.Marshal(payload)

	// form request
	request, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(payLoadBytes))
	if err != nil {
		log.Fatalf("Request failed with %v\n", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	// exec
	resp, err := client.Do(request)
	defer resp.Body.Close()

	// resp decoding
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with %v\n", resp.StatusCode)
	} else {
		log.Printf("Successfully called the alert")
	}

}
