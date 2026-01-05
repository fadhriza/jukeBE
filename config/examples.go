package config

// examples.go
//
// This directory contains configuration logic for the application.
// It uses `godotenv` to load environment variables from .env files.
//
// Example Usage:
//
// func main() {
//     cfg := config.LoadConfig()
//     fmt.Printf("Running in %s mode", cfg.AppEnv)
//     fmt.Printf("Connecting to DB at %s:%s", cfg.DBHost, cfg.DBPort)
// }
//
// By default, it looks for APP_ENV environment variable to decide which .env file to load.
// APP_ENV=local   -> loads .env.local
// APP_ENV=staging -> loads .env.staging
// APP_ENV=prod    -> loads .env.prod
