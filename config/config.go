package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		GRPC `yaml:"grpc"`
		Log  `yaml:"logger"`
		MQTT `yaml:"mqtt"`
		// TokenTTL time.Duration `yaml:"tocken_ttl" env-required:"true"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// GRPC -.
	GRPC struct {
		Port    string        `env-required:"true" yaml:"port"    env:"GRPC_PORT"`
		Timeout time.Duration `env-required:"true" yaml:"timeout" env:"GRPC_TIMEOUT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	// MQTT -.
	MQTT struct {
		URL      string `env-required:"true" yaml:"url"       env:"MQTT_URL"`
		ClientID string `env-required:"true" yaml:"client_id" env:"MQTT_CLIENT_ID"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
