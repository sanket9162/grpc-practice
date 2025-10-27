package main

import (
	"context"
	"log"
	"time"

	mainapipb "grpcclient/proto/gen"
	farewellpb "grpcclient/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculatorClient(conn)
	client2 := mainapipb.NewGreeterClient(conn)
	fwClient := farewellpb.NewAufWiedersehenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := mainapipb.AddRequest{
		A: 10,
		B: 20,
	}
	md := metadata.Pairs("authorization", "Bearer=ahskgnoeqhvnaodeaebhre", "test", "testing", "test2", "testing2")
	ctx = metadata.NewOutgoingContext(ctx, md)
	var resHeader metadata.MD
	res, err := client.Add(ctx, &req, grpc.Header(&resHeader))
	if err != nil {
		log.Fatalln("could not add", err)
	}
	log.Println("resHeader", resHeader)
	log.Println("resHeader[test]:", resHeader["test"])

	reqGreet := mainapipb.HelloRequest{
		Name: "Sanket",
	}
	res1, err := client2.Greet(ctx, &reqGreet)
	if err != nil {
		log.Fatalln("could not greet", err)
	}
	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "sanket",
	}
	resFw, err := fwClient.BidGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatalln("Could not bid Goodbye", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("Sum:", res1.Message)
	log.Println("Goodbye message:", resFw.Message)
	state := conn.GetState()
	log.Println("Connection State:", state)

}
