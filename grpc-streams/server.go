package main

import (
	mainpb "grpcstreams/proto/gen"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	mainpb.UnimplementedCalculatorServer
}

func (s *server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculator_GenerateFibonacciServer) error {
	n := req.N
	a, b := 0, 1

	for i := 0; i < int(n); i++ {
		err := stream.Send(&mainpb.FibonacciResponse{
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServe := grpc.NewServer()
	mainpb.RegisterCalculatorServer(grpcServe, &server{})

	log.Println("server is running on port 50051")
	err = grpcServe.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
