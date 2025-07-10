package main

import (
	"context"
	"flag"
	"log"

	sensor "github.com/sensor-data-sim/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	client, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("An error occurred while creating the client:", err)
	}
	defer client.Close()

	ctx := context.Background()
	req := &sensor.HelloRequest{
		Name: "Hello, I'm Sensor Client 1",
	}

	sensorGreeterClient := sensor.NewGreeterClient(client)
	res, err := sensorGreeterClient.SayHello(ctx, req)
	if err != nil {
		log.Fatal("An error occurred while calling the SayHello method:", err)
	}

	log.Printf("Received response: %s", res.GetMessage())
}
