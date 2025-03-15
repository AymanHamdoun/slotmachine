package config

import (
	"go-backend/internal/app/utils/maputils"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppEnv  string
	AppHost string
	AppPort string

	JWTSecret string

	DatabaseHost string
	DatabaseUser string
	DatabasePass string
	DatabasePort string
	DatabaseName string

	S3Key    string
	S3Secret string
	S3Host   string
	S3Region string
}

func (c *Config) IsProd() bool {
	return c.AppEnv == "production"
}

var config *Config

func Get() *Config {
	return config
}

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Info().Err(err).Msgf("Loaded without .env")
	}

	config = &Config{
		AppEnv:  getEnvString("APP_ENV", "local"),
		AppHost: getEnvString("APP_HOST", "localhost"),
		AppPort: getEnvString("APP_PORT", "8080"),

		JWTSecret: getEnvString("JWT_KEY", ""),

		DatabaseHost: getEnvString("DB_HOST", "localhost"),
		DatabaseUser: getEnvString("DB_USER", "user"),
		DatabasePass: getEnvString("DB_PASS", "password"),
		DatabasePort: getEnvString("DB_PORT", "3306"),
		DatabaseName: getEnvString("DB_NAME", "default"),

		S3Key:    getEnvString("SPACES_KEY", ""),
		S3Secret: getEnvString("SPACES_SECRET", ""),
		S3Host:   getEnvString("SPACES_HOST", ""),
		S3Region: getEnvString("SPACES_REGION", ""),
	}
}

func getEnvString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	trueValues := map[string]bool{"true": true, "1": true}
	if value, ok := os.LookupEnv(key); ok {
		return maputils.HasKey(trueValues, strings.ToLower(value))
	}
	return fallback
}
