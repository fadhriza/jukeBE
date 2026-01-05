package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

// LoadConfig loads configuration from .env file based on APP_ENV or defaults to .env.local
func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Printf("Warning: Error loading .env.%s file. Using system environment variables.", env)
	}

	return &Config{
		AppEnv:     getEnv("APP_ENV", "local"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "juke"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
