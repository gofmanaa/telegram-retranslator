package config

import (
	"os"
	"strconv"
)

type Config struct {
	TelegramApiKey           string
	TelegramSpaceGettoChatId int
	RedisPass                string
}

// New returns a new Config struct
func New() *Config {
	chatId, _ := strconv.Atoi(getEnv("TELEGRAM_SPACE_GETTO_CHAT_ID", ""))
	return &Config{
		TelegramApiKey:           getEnv("TELEGRAM_API_KEY", ""),
		TelegramSpaceGettoChatId: chatId,
		RedisPass:                getEnv("REDIS_PASS", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
