package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port                  string
	GinMode               string
	DBUrl                 string
	JWT_SECRET            string
	TURSO_DATABASE_URL    string
	TURSO_AUTH_TOKEN      string
	CLOUDINARY_CLOUD_NAME string
	CLOUDINARY_API_KEY    string
	CLOUDINARY_API_SECRET string
}

// Global variable to access config anywhere
var Envs *EnvConfig

func LoadConfig() *EnvConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using system env vars")
	}

	Envs = &EnvConfig{
		Port:                  getEnv("PORT", "8080"),
		GinMode:               getEnv("GIN_MODE", "debug"),
		DBUrl:                 getEnv("DB_URL", ""),
		TURSO_DATABASE_URL:    getEnv("TURSO_DATABASE_URL", ""),
		TURSO_AUTH_TOKEN:      getEnv("TURSO_AUTH_TOKEN", ""),
		JWT_SECRET:            getEnv("JWT_SECRET", ""),
		CLOUDINARY_CLOUD_NAME: getEnv("CLOUDINARY_CLOUD_NAME", ""),
		CLOUDINARY_API_KEY:    getEnv("CLOUDINARY_API_KEY", ""),
		CLOUDINARY_API_SECRET: getEnv("CLOUDINARY_API_SECRET", ""),
	}
	return Envs
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
