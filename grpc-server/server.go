package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sanket9162/go-grpc/proto/gen"
	farewellpb "github.com/sanket9162/go-grpc/proto/gen/farewell"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
	pb.UnimplementedGreeterServer
	farewellpb.AufWiedersehenServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprint("Hello %s. Nice to receive request from you", req.Name),
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{
		Message: fmt.Sprint("Hello %s!. Nice to receive request from you farewell my friend", req.Name),
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
	pb.RegisterGreeterServer(grpcserver, &server{})
	farewellpb.RegisterAufWiedersehenServer(grpcserver, &server{})

	log.Println("server is running on port", port)
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
