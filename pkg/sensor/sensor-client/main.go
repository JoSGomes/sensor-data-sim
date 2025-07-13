package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

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
		log.Fatal("An error occurred while creating the client", err)
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

	sensorSim := service.NewSensorSimulator(fmt.Sprintf("Sensor - %s", location.Description), location)
	sensorSim.StartGeneratingReadings(ctx, time.Duration(sensorConf.Sensor.Interval)*time.Second)

	sensorGreeterClient := sensor.NewGreeterClient(client)

	for {
		time.Sleep(5 * time.Second)
		req := sensorSim.GetSensorData()
		res, err := sensorGreeterClient.SendSensorData(ctx, req)
		if err != nil {
			log.Fatal("An error occurred while calling the SendSensorData method", err)
		}
		if !res.GetSuccess() {
			log.Printf("Failed to send sensor data: %v", res.GetMessage())
			continue
		}

		log.Printf("Success response from server: %v", res.GetMessage())
	}

}
