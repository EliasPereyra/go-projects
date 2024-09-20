package main

import (
	"context"
	"log"
	"net"

	pb "basic-grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (grpcServer *server) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.SuccessResponse, error) {
	return &pb.SuccessResponse{Response: "Done!"}, nil
}

func main() {
	tcpSer, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, &server{})
	log.Printf("gRPC server listening at %v", tcpSer.Addr())
	if err := grpcServer.Serve(tcpSer); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
