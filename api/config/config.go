package config

import (
	"os"
)

// Config содержит все настройки API сервиса
type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// OpenRouter
	OpenRouterAPIKey string
	OpenRouterURL    string
	AIModel          string

	// Telegram
	TelegramBotToken string

	// API
	APIPort string
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	return &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "postgres"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "telegram_bot"),

		// OpenRouter
		OpenRouterAPIKey: getEnv("OPENROUTER_API_KEY", ""),
		OpenRouterURL:    getEnv("OPENROUTER_URL", "https://openrouter.ai/api/v1"),
		AIModel:          getEnv("AI_MODEL", "deepseek/deepseek-chat-v3.1:free"),

		// Telegram
		TelegramBotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),

		// API
		APIPort: getEnv("API_PORT", "8080"),
	}
}

// Validate проверяет обязательные поля конфигурации
func (c *Config) Validate() error {
	if c.OpenRouterAPIKey == "" {
		return &ConfigError{Field: "OPENROUTER_API_KEY", Message: "OpenRouter API key is required"}
	}
	if c.TelegramBotToken == "" {
		return &ConfigError{Field: "TELEGRAM_BOT_TOKEN", Message: "Telegram bot token is required"}
	}
	return nil
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ConfigError представляет ошибку конфигурации
type ConfigError struct {
	Field   string
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}
