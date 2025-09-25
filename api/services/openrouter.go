package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"telegram-api/models"
)

// OpenRouterService сервис для работы с OpenRouter API
type OpenRouterService struct {
	apiKey string
	url    string
	model  string
	client *http.Client
}

// NewOpenRouterService создает новый сервис OpenRouter
func NewOpenRouterService(apiKey, url, model string) *OpenRouterService {
	return &OpenRouterService{
		apiKey: apiKey,
		url:    url,
		model:  model,
		client: &http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        10,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 5 * time.Second,
			},
		},
	}
}

// SendMessage отправляет сообщение в OpenRouter и получает ответ
func (s *OpenRouterService) SendMessage(messages []*models.Message) (*models.Message, error) {
	// Подготавливаем запрос
	request := models.OpenRouterRequest{
		Model:       s.model,
		Messages:    make([]models.Message, len(messages)),
		MaxTokens:   500,
		Temperature: 0.7,
	}

	// Копируем сообщения
	for i, msg := range messages {
		request.Messages[i] = *msg
	}

	// Сериализуем в JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Создаем HTTP запрос
	req, err := http.NewRequest("POST", s.url+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("HTTP-Referer", "https://telegram-bot.local")
	req.Header.Set("X-Title", "Telegram Bot")

	// Отправляем запрос
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("openrouter API error: %d - %s", resp.StatusCode, string(body))
	}

	// Парсим ответ
	var response models.OpenRouterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Проверяем на ошибки
	if response.Error != nil {
		return nil, fmt.Errorf("openrouter error: %s", response.Error.Message)
	}

	// Проверяем наличие ответа
	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no response from openrouter")
	}

	// Создаем сообщение-ответ
	assistantMessage := &models.Message{
		Content:   response.Choices[0].Message.Content,
		Role:      "assistant",
		CreatedAt: time.Now(),
	}

	return assistantMessage, nil
}
