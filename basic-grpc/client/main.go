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

	fmt.Println("hello")
	generateReq := &pb.HelloReq{Message: "Wow"}
	generateRes, err := client.Sayhello(ctx, generateReq)
	if err != nil {
		log.Fatalf("Error when generating the req: %v", err)
	}
	fmt.Print("Greeting: ", generateRes)
}
