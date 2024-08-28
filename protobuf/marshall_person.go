package main

import (
	"log"
	"os"
	"protobuf/models"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := &models.Person{
		Name:  "John Doe",
		Email: "yJp4x@example.com",
		Id:    1,
		Phones: []*models.PhoneNumber{
			&models.PhoneNumber{
				Number: "1234567890",
				Type:   models.PhoneType_HOME,
			},
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Failed to marshal: %v", err)
	}

	os.WriteFile("tmp/person.bin", data, 0644)
}
