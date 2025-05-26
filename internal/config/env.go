package config

import (
	"backend/internal/logger"
	"fmt"
	"os"
	"sync"
)

type Env struct {
	AppName          string
	AppHost          string
	AppPort          string
	DBServiceUrl     string
	DBHost           string
	DBUsername       string
	DBPassword       string
	DBName           string
	DBPort           string
	LogLevel         string
	LogFormat        string
	StaticSecretKey  string
	DynamicSecretKey string
}

var (
	env  *Env
	once sync.Once
)

func Load() *Env {
	once.Do(func() {
		env = &Env{
			AppName:          getEnv("APP_NAME", "app"),
			AppHost:          getEnv("APP_HOST", "localhost"),
			AppPort:          getEnv("APP_PORT", "8080"),
			DBServiceUrl:     getEnv("DB_SERVICE_URL", "http://localhost:8080"),
			DBHost:           getEnv("DB_HOST", "localhost"),
			DBUsername:       getEnv("DB_USER", "admin"),
			DBPassword:       getEnv("DB_PASSWORD", "securepassword"),
			DBName:           getEnv("DB_NAME", "military_db"),
			DBPort:           getEnv("DB_PORT", "5432"),
			LogLevel:         getEnv("LOG_LEVEL", "info"),
			LogFormat:        getEnv("LOG_FORMAT", "text"),
			StaticSecretKey:  getEnv("STATIC_SECRET", "static-secret-key"),
			DynamicSecretKey: getEnv("DYNAMIC_SECRET", "dynamic-secret-key"),
		}
	})
	return env
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	logger.Warn(fmt.Sprintf("%s not found in env, using default: %s", key, defaultValue))
	return defaultValue
}
