package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	TelegramAPIKey           string
	TelegramSpaceGettoChatID int
	RedisPass                string
	Log                      *log.Logger
}

// New returns a new Config struct
func New(log *log.Logger) *Config {
	chatID, _ := strconv.Atoi(getEnv("TELEGRAM_SPACE_GETTO_CHAT_ID", ""))

	return &Config{
		TelegramAPIKey:           getEnv("TELEGRAM_API_KEY", ""),
		TelegramSpaceGettoChatID: chatID,
		RedisPass:                getEnv("REDIS_PASS", ""),
		Log:                      log,
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
