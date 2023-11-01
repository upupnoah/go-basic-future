package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendHelloworld() {
	// helloworld (GET http://localhost:8080/hello)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	sendHelloworld()
}
