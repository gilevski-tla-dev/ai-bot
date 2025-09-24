package bot

import (
	"log"

	"telegram-bot/config"
	"telegram-bot/database"
	"telegram-bot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot представляет Telegram бота
type Bot struct {
	api        *tgbotapi.BotAPI
	dbConn     *database.Connection
	cmdHandler *handlers.CommandHandler
}

// New создает новый экземпляр бота
func New(cfg *config.Config) (*Bot, error) {
	// Инициализация Telegram API
	botAPI, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	// Инициализация базы данных
	dbConn, err := database.NewConnection(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		return nil, err
	}

	// Создание таблиц
	if err := dbConn.CreateTables(); err != nil {
		return nil, err
	}

	// Инициализация репозитория и обработчика
	userRepo := database.NewUserRepository(dbConn.GetDB())
	cmdHandler := handlers.NewCommandHandler(userRepo)

	return &Bot{
		api:        botAPI,
		dbConn:     dbConn,
		cmdHandler: cmdHandler,
	}, nil
}

// Start запускает бота
func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	log.Println("Bot started, waiting for updates...")
	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		}
	}
}

// handleMessage обрабатывает входящие сообщения
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		b.cmdHandler.HandleCommand(b.api, message)
	}
}

// Close закрывает соединения
func (b *Bot) Close() error {
	return b.dbConn.Close()
}
