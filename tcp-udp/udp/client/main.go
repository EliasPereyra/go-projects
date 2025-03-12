package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	message := "Hello to the server"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sent: ", message)

	// We set a dealine to avoid blocking indefenitely
	conn.SetReadDeadline(time.Now().Add(time.Second))

	// we receive a response from the server
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Received: ", string(buff[:n]))
}
