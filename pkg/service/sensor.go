package service

import (
	"context"
	"github.com/sensor-data-sim/proto"
	"math/rand"
	"time"
)

type SensorSimulatorService struct {
	SensorDataReq *sensor.SensorDataRequest
	ticker        *time.Ticker
	done          chan bool
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
		done: make(chan bool),
	}
}

func (s *SensorSimulatorService) GetSensorData() *sensor.SensorDataRequest {
	return s.SensorDataReq
}

func (s *SensorSimulatorService) StartGeneratingReadings(ctx context.Context, interval time.Duration) {
	s.ticker = time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-s.ticker.C:
				s.updateSensorData()
			case <-s.done:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *SensorSimulatorService) updateSensorData() {
	s.SensorDataReq.Data.TemperatureCelsius = 20.0 + rand.Float64()*15.0
	s.SensorDataReq.Data.HumidityPercent = 50.0 + rand.Float64()*30.0
	s.SensorDataReq.Data.NoiseLevelDb = 30.0 + rand.Float64()*20.0
	s.SensorDataReq.Data.LightIntensityLux = 100 + rand.Float64()*900
	s.SensorDataReq.Data.AirQualityIndex = 0 + rand.Float64()*100
	s.SensorDataReq.Data.VibrationLevel = 0 + rand.Float64()*10
	s.SensorDataReq.Data.Battery.Voltage = 3.7 + rand.Float64()*0.5
	s.SensorDataReq.Data.Battery.ChargePercent = 20 + rand.Float64()*80
	s.SensorDataReq.Data.Battery.IsCharging = rand.Intn(2) == 0
	s.SensorDataReq.Data.Battery.TemperatureCelsius = 20.0 + rand.Float64()*15.0

	customMetrics := map[string]float64{
		"co2_ppm":        400.0 + rand.Float64()*200.0, // CO2 em ppm
		"dust_particles": rand.Float64() * 100.0,       // Partículas de poeira
		"uv_index":       rand.Float64() * 11.0,        // Índice UV 0-11
		"wind_speed_ms":  rand.Float64() * 15.0,        // Velocidade do vento m/s
		"rainfall_mm":    rand.Float64() * 5.0,         // Precipitação mm/h
	}

	s.SensorDataReq.Data.CustomMetrics = customMetrics
}

func (s *SensorSimulatorService) Stop() {
	if s.ticker != nil {
		s.ticker.Stop()
	}
	close(s.done)
}
