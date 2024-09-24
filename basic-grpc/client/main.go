package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "basic-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// conect to the grpc server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)

	// we add timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	generateCreateReq := &pb.CreatePersonRequest{Name: "elias", Email: "elias@mail.com", PhoneNumber: "123456"}
	generateCreateRes, err := client.Create(ctx, generateCreateReq)
	if err != nil {
		log.Fatalf("Error when generating the req: %v", err)
	}
	fmt.Println("--Person created: ", generateCreateRes)

	fmt.Println("--Getting a user by ID")
	readReq := &pb.SinglePersonRequest{Id: 1}
	readRes, err := client.Read(ctx, readReq)
	if err != nil {
		log.Fatalf("Error when getting a user: %v", err)
	}
	fmt.Println("The person requested: ", readRes)

	fmt.Println("--Update a user:")
	updatedReq := &pb.UpdatePersonRequest{
		Id:          1,
		Name:        "eliseo",
		Email:       "eliseo@mail.com",
		PhoneNumber: "123151515",
	}
	updateRes, err := client.Update(ctx, updatedReq)
	if err != nil {
		log.Fatalf("Error when updating a user: %v", err)
	}
	newUserReq := &pb.SinglePersonRequest{Id: 1}
	newUserRes, err := client.Read(ctx, newUserReq)
	if err != nil {
		log.Fatalf("Error when getting a user: %v", err)
	}
	fmt.Println(updateRes)
	fmt.Println("User updated: ", newUserRes)

	fmt.Println("--Delete a user by ID:")
	delUserReq := pb.SinglePersonRequest{Id: 1}
	delUserRes, err := client.Delete(ctx, &delUserReq)
	if err != nil {
		log.Fatalf("Error when trying to delete a user: %v", err)
	}
	fmt.Println(delUserRes)
}
