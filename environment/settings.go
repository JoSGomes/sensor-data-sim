package environment

type Settings struct {
	Sensor struct {
		Location struct {
			Latitude    float64 `envconfig:"LATITUDE" default:"-23.5505"`  // São Paulo
			Longitude   float64 `envconfig:"LONGITUDE" default:"-46.6333"` // São Paulo
			Altitude    float64 `envconfig:"ALTITUDE" default:"760"`       // Altitude in meters
			Description string  `envconfig:"DESCRIPTION" default:"São Paulo - Centro"`
		}
		Interval int `envconfig:"INTERVAL" default:"5000"` // Interval in milli to generate sensor data
	}
}
