package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string // dev | prod | test
}

// Load lee variables de entorno, aplicando defaults y .env en desarrollo
func Load() Config {
	_ = godotenv.Load() // .env es opcional

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
