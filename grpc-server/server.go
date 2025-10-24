package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	log.Println("server is running on port", port)
	grpcserver := grpc.NewServer()
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
