package config

import (
	"os"
)

type Configuration struct {
	TelegramApiToken string
}

func Load() *Configuration {
	return &Configuration{
		TelegramApiToken: getEnv("TELEGRAM_API_TOKEN", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
