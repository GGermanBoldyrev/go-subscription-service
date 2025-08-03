package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	DatabaseDSN string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("Environment variable DB_URL is required but not set")
	}

	return &Config{
		Port:        port,
		DatabaseDSN: dsn,
	}
}
