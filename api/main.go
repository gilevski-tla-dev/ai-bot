package main

import (
	"database/sql"
	"fmt"
	"log"

	"telegram-api/config"
	"telegram-api/handlers"
	"telegram-api/middleware"
	"telegram-api/models"
	"telegram-api/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Валидируем конфигурацию
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Инициализируем базу данных
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализируем сервисы
	messageRepo := models.NewMessageRepository(db)
	openRouterSvc := services.NewOpenRouterService(cfg.OpenRouterAPIKey, cfg.OpenRouterURL, cfg.AIModel)
	telegramAuthSvc := services.NewTelegramAuthService(cfg.TelegramBotToken)

	// Инициализируем обработчики
	chatHandler := handlers.NewChatHandler(messageRepo, openRouterSvc, telegramAuthSvc)

	// Настраиваем Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())
	r.Use(gin.Recovery())

	// Публичные маршруты
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "telegram-api",
		})
	})

	// Защищенные маршруты
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(telegramAuthSvc))
	{
		api.POST("/chat", chatHandler.SendMessage)
		api.GET("/history", chatHandler.GetHistory)
		api.GET("/stats", chatHandler.GetStats)
	}

	// Запускаем сервер
	log.Printf("Starting API server on port %s", cfg.APIPort)
	if err := r.Run(":" + cfg.APIPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// initDatabase инициализирует подключение к базе данных
func initDatabase(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Тестируем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
