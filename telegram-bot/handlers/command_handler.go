package handlers

import (
	"fmt"
	"log"
	"time"

	"telegram-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CommandHandler обрабатывает команды Telegram бота
type CommandHandler struct {
	userRepo models.UserRepository
}

// NewCommandHandler создает новый обработчик команд
func NewCommandHandler(userRepo models.UserRepository) *CommandHandler {
	return &CommandHandler{
		userRepo: userRepo,
	}
}

// HandleCommand обрабатывает входящую команду
func (h *CommandHandler) HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Command() == "start" {
		h.handleStartCommand(bot, message)
	}
}

// handleStartCommand обрабатывает команду /start
func (h *CommandHandler) handleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	user := &models.User{
		UserID:    int64(message.From.ID),
		Username:  message.From.UserName,
		FirstName: message.From.FirstName,
		LastName:  message.From.LastName,
		CreatedAt: time.Now(),
	}

	// Сохраняем пользователя в базу данных
	if err := h.userRepo.Save(user); err != nil {
		log.Printf("Error saving user: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка при сохранении данных")
		bot.Send(msg)
		return
	}

	// Отправляем приветственное сообщение
	welcomeText := fmt.Sprintf(
		"Привет, %s! 👋\n\n"+
			"Добро пожаловать в бота! Я успешно сохранил ваши данные в базу данных.\n\n"+
			"Ваш ID: %d\n"+
			"Имя: %s",
		user.FirstName,
		user.UserID,
		user.FirstName,
	)

	if user.LastName != "" {
		welcomeText += fmt.Sprintf(" %s", user.LastName)
	}

	if user.Username != "" {
		welcomeText += fmt.Sprintf("\nUsername: @%s", user.Username)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, welcomeText)
	bot.Send(msg)
}
