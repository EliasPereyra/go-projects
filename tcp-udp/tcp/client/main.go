package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(conn, "Hello!")
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buff[:n]))
}
