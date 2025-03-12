package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buff[:n]))

	fmt.Fprintf(conn, "Echo - "+string(buff[:n]))
}

func handleNonBlockingConn(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetDeadline(time.Now().Add(time.Second))

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
				log.Println("Connection closed", err)
				break
			}
		}

		// Print that the message was receieved
		fmt.Println("Received: ", string(buff[:n]))

		// set a deadline for writing data
		conn.SetDeadline(time.Now().Add(time.Second))

		_, err = fmt.Fprintf(conn, "Echo\n")
		if err != nil {
			fmt.Println("There was an error when writing")
		}
	}
}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Couldn't connect", err)
	}
	defer lst.Close()

	for {
		conn, err := lst.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleNonBlockingConn(conn)
	}
}
