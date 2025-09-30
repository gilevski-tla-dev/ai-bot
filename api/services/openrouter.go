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
	apiKey     string
	url        string
	model      string
	timeout    int
	maxRetries int
	client     *http.Client
}

// NewOpenRouterService создает новый сервис OpenRouter
func NewOpenRouterService(apiKey, url, model string) *OpenRouterService {
	return &OpenRouterService{
		apiKey:     apiKey,
		url:        url,
		model:      model,
		timeout:    30, // 30 секунд таймаут
		maxRetries: 3,  // 3 попытки
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second, // Увеличиваем таймаут подключения
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          20, // Увеличиваем пул соединений
				MaxIdleConnsPerHost:   10,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second, // Увеличиваем таймаут TLS
				ResponseHeaderTimeout: 15 * time.Second, // Добавляем таймаут заголовков
			},
		},
	}
}

// SendMessage отправляет сообщение в OpenRouter и получает ответ с retry логикой
func (s *OpenRouterService) SendMessage(messages []*models.Message) (*models.Message, error) {
	const baseDelay = 1 * time.Second

	for attempt := 0; attempt < s.maxRetries; attempt++ {
		message, err := s.sendMessageAttempt(messages)
		if err == nil {
			return message, nil
		}

		// Логируем попытку
		fmt.Printf("OpenRouter attempt %d/%d failed: %v\n", attempt+1, s.maxRetries, err)

		// Если это последняя попытка, возвращаем ошибку
		if attempt == s.maxRetries-1 {
			return nil, fmt.Errorf("failed after %d attempts: %w", s.maxRetries, err)
		}

		// Экспоненциальная задержка с jitter
		delay := baseDelay * time.Duration(1<<attempt) // 1s, 2s, 4s
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("unexpected error in retry loop")
}

// sendMessageAttempt выполняет одну попытку отправки сообщения
func (s *OpenRouterService) sendMessageAttempt(messages []*models.Message) (*models.Message, error) {
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
