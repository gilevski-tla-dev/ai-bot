package main

import (
	"log"

	"telegram-bot/bot"
	"telegram-bot/config"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Валидируем конфигурацию
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Создаем экземпляр бота
	telegramBot, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}
	defer telegramBot.Close()

	// Запускаем бота
	telegramBot.Start()
}
