package config

import (
	"os"
	"strconv"
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Port string
	}

	Database struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     string
	}

	YouTube struct {
		UploadTime int
	}
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	cfg := &Config{}

	// Load server configuration
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// Load database configuration
	cfg.Database.User = getEnv("DB_USER", "default_user")
	cfg.Database.Password = getEnv("DB_PASSWORD", "default_password")
	cfg.Database.DBName = getEnv("DB_NAME", "default_db")
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")

	return cfg
}

// Helper function to get environment variables with a fallback
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// Helper function to parse integers with a fallback
func parseInt(value string, fallback int) int {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intVal
}
