// Exemplo simples para testar o simulador sem gRPC
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sensor-data-sim/pkg/sensor"
)

func main() {
	fmt.Println("🌟 Testando Simulador de Sensor IoT")
	fmt.Println("=====================================")

	// Cria uma localização de exemplo
	location := &sensor.Location{
		Latitude:    -23.5505,
		Longitude:   -46.6333,
		Altitude:    760,
		Description: "São Paulo - Centro",
	}

	// Cria o simulador
	simulator := sensor.NewSensorSimulator("sensor-test-001", location)

	fmt.Printf("📍 Sensor ID: %s\n", simulator.SensorID)
	fmt.Printf("📍 Localização: %s\n", location.Description)
	fmt.Printf("📍 Coordenadas: %.6f, %.6f (Alt: %.0fm)\n\n",
		location.Latitude, location.Longitude, location.Altitude)

	// Gera algumas leituras de exemplo
	fmt.Println("📊 Gerando dados simulados:")
	fmt.Println("--------------------------------------------------")

	for i := 1; i <= 5; i++ {
		reading := simulator.GenerateReading()

		fmt.Printf("\n🔄 Leitura #%d - %s\n", i, reading.Timestamp.Format("15:04:05"))
		fmt.Printf("   🌡️  Temperatura: %.2f°C\n", reading.Temperature)
		fmt.Printf("   🌊 Pressão: %.2f hPa\n", reading.Pressure)
		fmt.Printf("   💧 Umidade: %.2f%%\n", reading.Humidity)
		fmt.Printf("   🔊 Ruído: %.2f dB\n", reading.NoiseLevel)
		fmt.Printf("   💡 Luz: %.2f lux\n", reading.LightIntensity)
		fmt.Printf("   🌬️  Qualidade do ar: %.2f\n", reading.AirQualityIndex)
		fmt.Printf("   📳 Vibração: %.2f\n", reading.VibrationLevel)
		fmt.Printf("   🔋 Bateria: %.2f%% (%.2fV)\n", reading.BatteryPercent, reading.BatteryVoltage)

		// Métricas customizadas
		fmt.Printf("   📈 CO2: %.1f ppm\n", reading.CustomMetrics["co2_ppm"])
		fmt.Printf("   ☀️  UV: %.1f\n", reading.CustomMetrics["uv_index"])
		fmt.Printf("   💨 Vento: %.1f m/s\n", reading.CustomMetrics["wind_speed_ms"])
		fmt.Printf("   🌧️  Chuva: %.1f mm/h\n", reading.CustomMetrics["rainfall_mm"])

		time.Sleep(2 * time.Second) // Pausa entre leituras
	}

	fmt.Println("\n✅ Teste concluído com sucesso!")
	fmt.Println("\n💡 Para executar o cliente gRPC completo:")
	fmt.Println("   go run ./cmd/sensor")
}
