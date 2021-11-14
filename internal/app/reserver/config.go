package reserver

import "github.com/ulstu-schedule/reserver/internal/app/repository"

// Config ...
type Config struct {
	LogLevel            string `toml:"log_level"`
	ReservationInterval int    `toml:"reservation_interval"` // in hours
	Repository          *repository.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		LogLevel:            "debug",
		ReservationInterval: 2,
		Repository:          repository.NewConfig(),
	}
}
