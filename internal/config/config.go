// Package config handles application configuration loading.
package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds all application configuration values.
type Config struct {
	FilePath string `default:"./api/openapi.yml" envconfig:"FILE_PATH"`
}

// Load reads configuration from environment variables and returns a Config instance.
func Load() *Config {
	_ = godotenv.Load()

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to load server config: %s", err)
	}

	return &cfg
}
