package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/sensor-data-sim/environment"
	"github.com/sensor-data-sim/pkg/service"
	"github.com/sensor-data-sim/proto"

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

	var sensorConf environment.Settings

	err = envconfig.Process("SENSOR", &sensorConf)
	if err != nil {
		log.Fatal("couldn't load sensor environment variables")
	}

	ctx := context.Background()

	location := &sensor.Location{
		Latitude:       sensorConf.Sensor.Location.Latitude,
		Longitude:      sensorConf.Sensor.Location.Longitude,
		AltitudeMeters: sensorConf.Sensor.Location.Altitude,
		Description:    sensorConf.Sensor.Location.Description,
	}

	SensorSimulatorService := service.NewSensorSimulator(fmt.Sprintf("Sensor - %s", location.Description), location)

	sensorGreeterClient := sensor.NewGreeterClient(client)

	//res, err := sensorGreeterClient.SayHello(ctx, req)
	//if err != nil {
	//	log.Fatal("An error occurred while calling the SayHello method:", err)
	//}
	//
	//log.Printf("Received response: %s", res.GetMessage())
}
