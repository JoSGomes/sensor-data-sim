package sensor

import (
	"math"
	"math/rand"
	"time"
)

// SensorSimulator simula dados de sensores
type SensorSimulator struct {
	SensorID     string
	Location     *Location
	baseTemp     float64
	basePressure float64
	baseHumidity float64
	random       *rand.Rand
}

// Location representa a localização do sensor
type Location struct {
	Latitude    float64
	Longitude   float64
	Altitude    float64
	Description string
}

// SensorReading representa uma leitura do sensor
type SensorReading struct {
	SensorID        string
	Timestamp       time.Time
	Location        *Location
	Temperature     float64 // Celsius
	Pressure        float64 // hPa
	Humidity        float64 // %
	NoiseLevel      float64 // dB
	LightIntensity  float64 // lux
	AirQualityIndex float64
	VibrationLevel  float64
	BatteryVoltage  float64
	BatteryPercent  float64
	IsCharging      bool
	BatteryTemp     float64
	CustomMetrics   map[string]float64
}

// NewSensorSimulator cria um novo simulador de sensor
func NewSensorSimulator(sensorID string, location *Location) *SensorSimulator {
	return &SensorSimulator{
		SensorID:     sensorID,
		Location:     location,
		baseTemp:     20.0 + rand.Float64()*15.0, // Temperatura base entre 20-35°C
		basePressure: 1013.25,                    // Pressão atmosférica padrão
		baseHumidity: 50.0 + rand.Float64()*30.0, // Umidade base entre 50-80%
		random:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateReading gera uma leitura simulada do sensor
func (s *SensorSimulator) GenerateReading() *SensorReading {
	now := time.Now()

	// Simula variações temporais (ciclo diário, sazonal, etc.)
	hourOfDay := float64(now.Hour())
	dayOfYear := float64(now.YearDay())

	// Temperatura com variação diária e ruído
	tempVariation := 5.0 * math.Sin(2*math.Pi*hourOfDay/24.0)       // Variação diária
	seasonalVariation := 10.0 * math.Sin(2*math.Pi*dayOfYear/365.0) // Variação sazonal
	tempNoise := (s.random.Float64() - 0.5) * 2.0                   // Ruído ±1°C
	temperature := s.baseTemp + tempVariation + seasonalVariation + tempNoise

	// Pressão atmosférica com pequenas variações
	pressureNoise := (s.random.Float64() - 0.5) * 20.0 // Variação de ±10 hPa
	pressure := s.basePressure + pressureNoise

	// Umidade inversamente relacionada à temperatura
	humidityVariation := -tempVariation * 2.0 // Umidade diminui com temperatura
	humidityNoise := (s.random.Float64() - 0.5) * 10.0
	humidity := s.baseHumidity + humidityVariation + humidityNoise
	humidity = math.Max(0, math.Min(100, humidity)) // Limita entre 0-100%

	// Nível de ruído (varia com hora do dia - mais ruído durante o dia)
	baseNoise := 35.0                                          // dB base
	timeNoise := 15.0 * math.Sin(2*math.Pi*(hourOfDay-6)/12.0) // Mais ruído 6h-18h
	randomNoise := s.random.Float64() * 10.0
	noiseLevel := baseNoise + math.Max(0, timeNoise) + randomNoise

	// Intensidade da luz (relacionada com hora do dia)
	var lightIntensity float64
	if hourOfDay >= 6 && hourOfDay <= 18 {
		// Simula curva de luz solar durante o dia
		sunAngle := math.Sin(math.Pi * (hourOfDay - 6) / 12.0)
		lightIntensity = sunAngle * 50000.0 * (0.8 + s.random.Float64()*0.4) // 0-50000 lux
	} else {
		lightIntensity = s.random.Float64() * 100.0 // Luz artificial noturna
	}

	// Índice de qualidade do ar (varia com hora e ruído)
	baseAQI := 50.0 + s.random.Float64()*50.0                     // AQI base 50-100
	trafficEffect := 20.0 * math.Sin(2*math.Pi*(hourOfDay-8)/8.0) // Pico no rush
	airQuality := baseAQI + math.Max(0, trafficEffect)

	// Nível de vibração (relacionado com ruído)
	vibrationLevel := (noiseLevel - 30.0) / 10.0 * s.random.Float64()
	vibrationLevel = math.Max(0, vibrationLevel)

	// Simulação da bateria
	batteryPercent := 20.0 + s.random.Float64()*80.0    // 20-100%
	isCharging := s.random.Float64() < 0.1              // 10% de chance de estar carregando
	batteryVoltage := 3.0 + (batteryPercent/100.0)*1.2  // 3.0-4.2V
	batteryTemp := temperature + s.random.Float64()*5.0 // Bateria um pouco mais quente

	// Métricas customizadas
	customMetrics := map[string]float64{
		"co2_ppm":        400.0 + s.random.Float64()*200.0, // CO2 em ppm
		"dust_particles": s.random.Float64() * 100.0,       // Partículas de poeira
		"uv_index":       s.random.Float64() * 11.0,        // Índice UV 0-11
		"wind_speed_ms":  s.random.Float64() * 15.0,        // Velocidade do vento m/s
		"rainfall_mm":    s.random.Float64() * 5.0,         // Precipitação mm/h
	}

	return &SensorReading{
		SensorID:        s.SensorID,
		Timestamp:       now,
		Location:        s.Location,
		Temperature:     temperature,
		Pressure:        pressure,
		Humidity:        humidity,
		NoiseLevel:      noiseLevel,
		LightIntensity:  lightIntensity,
		AirQualityIndex: airQuality,
		VibrationLevel:  vibrationLevel,
		BatteryVoltage:  batteryVoltage,
		BatteryPercent:  batteryPercent,
		IsCharging:      isCharging,
		BatteryTemp:     batteryTemp,
		CustomMetrics:   customMetrics,
	}
}

// StartContinuousReading inicia a geração contínua de leituras
func (s *SensorSimulator) StartContinuousReading(interval time.Duration, callback func(*SensorReading)) chan bool {
	stop := make(chan bool)

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				reading := s.GenerateReading()
				callback(reading)
			case <-stop:
				return
			}
		}
	}()

	return stop
}
