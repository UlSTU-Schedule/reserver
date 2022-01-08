package config

import (
	"github.com/spf13/viper"
)

const reserverConfigFileName = "reserver"

type Config struct {
	DatabaseURL string

	LogLevel            string `mapstructure:"log_level"`
	ReservationInterval int    `mapstructure:"reservation_interval"` // in hours
}

func New(configsPath string) (*Config, error) {
	cfg := &Config{}

	if err := parseFromYml(configsPath, cfg); err != nil {
		return nil, err
	}

	if err := parseFromEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseFromYml(configPath string, cfg *Config) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(reserverConfigFileName)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg)
}

func parseFromEnv(cfg *Config) error {
	if err := viper.BindEnv("DATABASE_URL"); err != nil {
		return err
	}

	cfg.DatabaseURL = viper.GetString("DATABASE_URL")
	return nil
}
