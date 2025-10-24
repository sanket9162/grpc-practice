package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sanket9162/go-grpc/proto/gen"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	grpcserver := grpc.NewServer()

	pb.RegisterCalculatorServer(grpcserver, &server{})

	log.Println("server is running on port", port)
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
