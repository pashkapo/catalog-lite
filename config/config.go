package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppPort string
	DBHost  string
	DBName  string
	DBUser  string
	DBPass  string
}

func New() *Config {
	return &Config{
		AppPort: getEnv("PORT", "3000"),
		DBHost:  getEnv("DB_HOST", "0.0.0.0:5432"),
		DBName:  getEnv("DB_NAME", "catalog_lite"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASS", "postgres"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
