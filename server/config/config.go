package config

import (
	"os"
)

// database Config
var (
	Host     = getEnvOrDefault("DB_HOST", "postgres")
	Username = getEnvOrDefault("DB_USER", "user")
	Password = getEnvOrDefault("DB_PASSWORD", "password123456")
	Dbname   = getEnvOrDefault("DB_NAME", "taskManagement")
)

const (
	Port = 5432
)

// main  config
const (
	ServerAddress = "0.0.0.0:8000"
)

// secret key
var (
	JwtSecretKey = getEnvOrDefault("JWT_SECRET", "SecretKey")
)

// get environment variable function
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
