package service

import (
	"github.com/sensor-data-sim/proto"
	"math/rand"
)

type SensorSimulatorService struct {
	SensorDataReq *sensor.SensorDataRequest
}

func NewSensorSimulator(sensorID string, location *sensor.Location) *SensorSimulatorService {
	return &SensorSimulatorService{
		SensorDataReq: &sensor.SensorDataRequest{
			Data: &sensor.SensorData{
				SensorId:           sensorID,
				Location:           location,
				TemperatureCelsius: 20.0 + rand.Float64()*15.0, // Temperatura base entre 20-35°C
				PressureHpa:        1013.25,                    // Pressão atmosférica padrão
				HumidityPercent:    50.0 + rand.Float64()*30.0, // Umidade base entre 50-80%
				NoiseLevelDb:       30.0 + rand.Float64()*20.0, // Nível de ruído entre 30-50 dB
				LightIntensityLux:  100 + rand.Float64()*900,   // Intensidade de luz entre 100-1000 lux
				AirQualityIndex:    0 + rand.Float64()*100,     // Índice de qualidade do ar entre 0-100
				VibrationLevel:     0 + rand.Float64()*10,      // Nível de vibração entre 0-10
				Battery: &sensor.BatteryInfo{
					Voltage:            3.7 + rand.Float64()*0.5,   // Tensão da bateria entre 3.7-4.2V
					ChargePercent:      20 + rand.Float64()*80,     // Percentual da bateria entre 20-100%
					IsCharging:         rand.Intn(2) == 0,          // Aleatoriamente se está carregando ou não
					TemperatureCelsius: 20.0 + rand.Float64()*15.0, // Temperatura da bateria entre 20-35°C
				},
				CustomMetrics: map[string]float64{},
			},
		},
	}
}

func (s *SensorSimulatorService) GenerateReading() *sensor.SensorDataRequest {
	// Simulate sensor data generation logic here if needed
	// For now, we return the existing SensorDataReq
	return s.SensorDataReq
}

func (s *SensorSimulatorService) GetSensorData() *sensor.SensorDataRequest {
	return s.SensorDataReq
}
