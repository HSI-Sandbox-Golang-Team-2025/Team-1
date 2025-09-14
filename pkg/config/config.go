package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL      string
	JWTAccessSecret  string
	JWTRefreshSecret string
	JWTAccessTTL     string
	JWTRefreshTTL    string
	CORSAllowOrigin  string
	AppPort          string
	AppEnv           string
}

func LoadEnv() *Config {
	// Load data from .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://postgres:postgrs@localhost/sandbox_lms?sslmode=false"),
		JWTAccessSecret:  getEnv("JWT_ACCESS_SECRET", "SUPER_SECRET_KEY"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "SUPER_SECRET_KEY_REFRESH"),
		JWTAccessTTL:     getEnv("JWT_ACCESS_TTL", "15m"),
		JWTRefreshTTL:    getEnv("JWT_REFRESH_TTL", "7d"),
		CORSAllowOrigin:  getEnv("CORS_ALLOW_ORIGIN", "*"),
		AppPort:          getEnv("APP_PORT", "8080"),
		AppEnv:           getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
