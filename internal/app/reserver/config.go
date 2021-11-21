package reserver

// Config ...
type Config struct {
	LogLevel            string `toml:"log_level"`
	ReservationInterval int    `toml:"reservation_interval"` // in hours
	DatabaseURL         string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		LogLevel:            "debug",
		ReservationInterval: 2,
	}
}
