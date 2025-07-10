package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	sensor "github.com/sensor-data-sim/pkg/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	sensor.UnimplementedGreeterServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sensor.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(_ context.Context, input *sensor.HelloRequest) (*sensor.HelloReply, error) {
	log.Printf("Received: %v", input.GetName())
	return &sensor.HelloReply{Message: "Hello Sensor Client!\n I received: " + input.GetName()}, nil
}
