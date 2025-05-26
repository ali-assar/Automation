package config

import (
	"os"
)

type Config struct {
	AppHost      string
	AppPort      string
	DBServiceUrl string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	LogLevel     string
	LogFormat    string
	JWTSecret    string
}

func Load() *Config {
	return &Config{
		AppHost:      getEnv("APP_HOST", "0.0.0.0"),
		AppPort:      getEnv("APP_PORT", "8080"),
		DBServiceUrl: getEnv("DB_SERVICE_URL", "http://localhost:8081"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "mydb"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		LogFormat:    getEnv("LOG_FORMAT", "json"),
		JWTSecret:    getEnv("JWT_SECRET", "mysecretkey"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
