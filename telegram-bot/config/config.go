package config

import (
	"os"
)

// Config содержит все настройки приложения
type Config struct {
	BotToken   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// Load загружает конфигурацию из переменных окружения
func Load() *Config {
	return &Config{
		BotToken:   getEnv("TELEGRAM_BOT_TOKEN", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "telegram_bot"),
	}
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Validate проверяет обязательные поля конфигурации
func (c *Config) Validate() error {
	if c.BotToken == "" {
		return &ConfigError{Field: "TELEGRAM_BOT_TOKEN", Message: "Bot token is required"}
	}
	return nil
}

// ConfigError представляет ошибку конфигурации
type ConfigError struct {
	Field   string
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}
