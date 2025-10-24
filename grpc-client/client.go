package main

import (
	"context"
	"log"
	"time"

	mainapipb "grpcclient/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := mainapipb.AddRequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatalln("could not add", err)
	}

	log.Println("Sum:", res.Sum)

}
