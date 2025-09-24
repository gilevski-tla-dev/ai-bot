package models

import (
	"time"
)

// Message представляет сообщение в чате
type Message struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	Role      string    `json:"role" db:"role"` // "user" или "assistant"
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ChatRequest представляет запрос на отправку сообщения
type ChatRequest struct {
	Message string `json:"message" binding:"required,max=300"`
}

// ChatResponse представляет ответ от API
type ChatResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// TelegramWebAppData представляет данные от Telegram WebApp
type TelegramWebAppData struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

// OpenRouterRequest представляет запрос к OpenRouter API
type OpenRouterRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

// OpenRouterResponse представляет ответ от OpenRouter API
type OpenRouterResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// MessageRepository интерфейс для работы с сообщениями
type MessageRepository interface {
	Save(message *Message) error
	GetByUserID(userID int64, limit int) ([]*Message, error)
	GetUserMessageCount(userID int64) (int, error)
}
