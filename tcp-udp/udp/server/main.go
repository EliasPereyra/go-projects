package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	add, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", add)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for {
		buff := make([]byte, 1024)
		n, clientAdd, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error when reading: ", err)
			continue
		}

		fmt.Println("Received from ", clientAdd, string(buff[:n]))

		_, err = conn.WriteToUDP([]byte("Hey client"), clientAdd)
		if err != nil {
			fmt.Println("Error when writing", err)
		}

	}
}
