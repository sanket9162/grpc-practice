package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "grpc-gateway-restapi/proto/gen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// server is used to implement mainapi.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// Greet implements mainapi.GreeterServer
func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Validate the incoming request
	err := req.Validate()
	if err != nil {
		log.Printf("Validation failed: %v", err)
		// Return an invalid argument error if validation fails
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request: %v", err)
	}

	return &pb.HelloResponse{Message: "Hello, " + req.GetName()}, nil
}

func runGRPCSever() {
	// Create a listener on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port :50051...")
	// Start the server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

// Rest api
func runGatewayServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatal("Failed to register gRPC-Gateway handler:", err)
	}

	log.Println("HTTP Server is running on port: 8080..")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to serve HTTP:", err)
	}
}

func main() {
	go runGRPCSever()
	runGatewayServer()
}
