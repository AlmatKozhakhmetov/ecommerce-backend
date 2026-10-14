package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func MustLoad() *Config {
	// Pytayemsya zagruzit .env fayl (esli ego net, nastroiki vozmutsa iz peremennyh okruzheniya OS)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL environment variable is required")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}
}
