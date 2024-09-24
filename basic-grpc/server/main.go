package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "basic-grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPersonServiceServer
}

type Person struct {
	ID          int32
	Name        string
	Email       string
	PhoneNumber string
}

var NextID int32 = 1

var persons = make(map[int32]Person)

func (grpcServer *server) Create(ctx context.Context, in *pb.CreatePersonRequest) (*pb.PersonProfileResponse, error) {
	person := Person{Name: in.GetName(), Email: in.GetEmail(), PhoneNumber: in.GetPhoneNumber()}
	if person.Name == "" || person.Email == "" || person.PhoneNumber == "" {
		return &pb.PersonProfileResponse{}, errors.New("Fields are missing")
	}

	person.ID = NextID
	persons[NextID] = person
	NextID = NextID + 1

	return &pb.PersonProfileResponse{Id: person.ID, Name: person.Name, Email: person.Email, PhoneNumber: person.PhoneNumber}, nil
}

func (grpcServer *server) Read(ctx context.Context, in *pb.SinglePersonRequest) (*pb.PersonProfileResponse, error) {
	id := in.GetId()
	person := persons[id]
	if person.ID == 0 {
		return &pb.PersonProfileResponse{}, errors.New("User not found")
	}

	return &pb.PersonProfileResponse{Id: person.ID, Name: person.Name, Email: person.Email, PhoneNumber: person.PhoneNumber}, nil
}

func (grpcServer *server) Update(ctx context.Context, in *pb.UpdatePersonRequest) (*pb.SuccessResponse, error) {
	id := in.GetId()
	if id == 0 {
		return &pb.SuccessResponse{}, errors.New("User not found")
	}
	personReq := persons[id]
	personReq.Name = in.GetName()
	personReq.Email = in.GetEmail()
	personReq.PhoneNumber = in.GetPhoneNumber()

	if personReq.Name == "" || personReq.Email == "" || personReq.PhoneNumber == "" {
		return &pb.SuccessResponse{}, errors.New("Fields are missing")
	}
	persons[id] = personReq

	return &pb.SuccessResponse{Response: "Person updated"}, nil
}

func (grpcServer *server) Delete(ctx context.Context, in *pb.SinglePersonRequest) (*pb.SuccessResponse, error) {
	id := in.GetId()
	if id == 0 {
		return &pb.SuccessResponse{}, errors.New("User not found")
	}

	delete(persons, id)
	return &pb.SuccessResponse{Response: "Deleted!"}, nil
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
