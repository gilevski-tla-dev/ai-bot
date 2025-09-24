package handlers

import (
	"log"
	"net/http"
	"time"

	"telegram-api/models"
	"telegram-api/services"

	"github.com/gin-gonic/gin"
)

// ChatHandler обработчик для чат API
type ChatHandler struct {
	messageRepo     models.MessageRepository
	openRouterSvc   *services.OpenRouterService
	telegramAuthSvc *services.TelegramAuthService
}

// NewChatHandler создает новый обработчик чата
func NewChatHandler(
	messageRepo models.MessageRepository,
	openRouterSvc *services.OpenRouterService,
	telegramAuthSvc *services.TelegramAuthService,
) *ChatHandler {
	return &ChatHandler{
		messageRepo:     messageRepo,
		openRouterSvc:   openRouterSvc,
		telegramAuthSvc: telegramAuthSvc,
	}
}

// SendMessage обрабатывает отправку сообщения
func (h *ChatHandler) SendMessage(c *gin.Context) {
	// Получаем данные пользователя из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDInt64 := userID.(int64)

	// Проверяем лимит сообщений (50 в день)
	messageCount, err := h.messageRepo.GetUserMessageCount(userIDInt64)
	if err != nil {
		log.Printf("Error getting message count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if messageCount >= 50 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Daily message limit reached (50 messages)",
			"limit": 50,
			"used":  messageCount,
		})
		return
	}

	// Парсим запрос
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем сообщение пользователя
	userMessage := &models.Message{
		UserID:    userIDInt64,
		Content:   req.Message,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	// Сохраняем сообщение пользователя
	if err := h.messageRepo.Save(userMessage); err != nil {
		log.Printf("Error saving user message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	// Получаем историю сообщений для контекста (последние 10)
	history, err := h.messageRepo.GetByUserID(userIDInt64, 10)
	if err != nil {
		log.Printf("Error getting message history: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get message history"})
		return
	}

	// Отправляем в OpenRouter
	assistantMessage, err := h.openRouterSvc.SendMessage(history)
	if err != nil {
		log.Printf("Error sending message to OpenRouter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI response"})
		return
	}

	// Устанавливаем UserID для ответа
	assistantMessage.UserID = userIDInt64

	// Сохраняем ответ ассистента
	if err := h.messageRepo.Save(assistantMessage); err != nil {
		log.Printf("Error saving assistant message: %v", err)
		// Не возвращаем ошибку, так как ответ уже получен
	}

	// Логируем успешный запрос
	log.Printf("Chat request processed: UserID=%d, MessageCount=%d, ResponseLength=%d",
		userIDInt64, messageCount+1, len(assistantMessage.Content))

	// Возвращаем ответ
	c.JSON(http.StatusOK, models.ChatResponse{
		Message:   assistantMessage.Content,
		Timestamp: assistantMessage.CreatedAt,
	})
}

// GetHistory получает историю сообщений пользователя
func (h *ChatHandler) GetHistory(c *gin.Context) {
	// Получаем данные пользователя из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDInt64 := userID.(int64)

	// Получаем историю сообщений
	messages, err := h.messageRepo.GetByUserID(userIDInt64, 50)
	if err != nil {
		log.Printf("Error getting message history: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get message history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
	})
}

// GetStats получает статистику пользователя
func (h *ChatHandler) GetStats(c *gin.Context) {
	// Получаем данные пользователя из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDInt64 := userID.(int64)

	// Получаем количество сообщений за сегодня
	messageCount, err := h.messageRepo.GetUserMessageCount(userIDInt64)
	if err != nil {
		log.Printf("Error getting message count: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"daily_messages": messageCount,
		"daily_limit":    50,
		"remaining":      50 - messageCount,
	})
}
