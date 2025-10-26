package main

import (
	"context"
	mainpb "grpcstreamsclient/proto/gen"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mainpb.NewCalculatorClient(conn)

	ctx := context.Background()

	req := &mainpb.FibonacciRequest{
		N: 10,
	}
	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln("Error calling GenerateFibonacci func:", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			break
		}
		if err != nil {
			log.Fatal("Error receiving data from GenerateFibonacci")
		}
		log.Println("Fibonacci number: ", resp.GetNumber())
	}

	stream1, err := client.SendNumber(ctx)
	if err != nil {
		log.Fatalln("Error creating stream:", err)
	}

	for num := range 9 {
		log.Println("Sending:", num)
		err := stream1.Send(&mainpb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln("Error sending number:", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error receiving response:", err)
	}
	log.Println("SUM:", res.Sum)

}
