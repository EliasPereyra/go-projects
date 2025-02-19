package main

import (
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	url := "http://localhost:8080/"

	for i := 1; i <= 150; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("There was an error when making a request", err)
			continue
		}

		fmt.Printf("Request %2d: Status %d\n", i, resp.StatusCode)
		time.Sleep(100 * time.Millisecond)
	}
}
