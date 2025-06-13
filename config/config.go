package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Port     string
	LogLevel string
}

var AppConfig *Config

func Load() *Config {
	_ = godotenv.Load()

	AppConfig = &Config{
		Env:      getEnv("APP_ENV", "development"),
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}
	return AppConfig
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
