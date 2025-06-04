package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string // mod for environment (dev, test, prod)
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: mustEnv("DATABASE_URL"),
		Env:         getEnv("ENV", "dev"),
	}
	return cfg
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing required env var: %s", key)
	}
	return val
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
