package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	sensor "github.com/sensor-data-sim/proto"

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

func (s *server) SendSensorData(_ context.Context, input *sensor.SensorDataRequest) (*sensor.SensorDataResponse, error) {
	log.Printf("Received: %s", input.GetData().SensorId)
	response := &sensor.SensorDataResponse{
		Message: fmt.Sprintf("Hello %s, I received your message!", input.GetData().SensorId),
		Success: true,
	}
	return response, nil
}
