package main

import (
	"fmt"
	"log"
	"os"
	"protobuf/models"

	"google.golang.org/protobuf/proto"
)

func main() {
	data, err := os.ReadFile("tmp/person.bin")
	if err != nil {
		log.Fatal("Failed to read file: %v", err)
		return
	}

	person := &models.Person{}
	err = proto.Unmarshal(data, person)

	if err != nil {
		log.Fatal("Failed to unmarshal: %v", err)
	}

	fmt.Println("Deserialized Person: ", person)
}
