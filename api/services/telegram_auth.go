package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"telegram-api/models"
)

// TelegramAuthService сервис для аутентификации через Telegram WebApp
type TelegramAuthService struct {
	botToken string
}

// NewTelegramAuthService создает новый сервис аутентификации
func NewTelegramAuthService(botToken string) *TelegramAuthService {
	return &TelegramAuthService{
		botToken: botToken,
	}
}

// ValidateWebAppData проверяет подлинность данных от Telegram WebApp
func (s *TelegramAuthService) ValidateWebAppData(initData string) (*models.TelegramWebAppData, error) {
	// Парсим данные
	params, err := url.ParseQuery(initData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse init data: %w", err)
	}

	// Извлекаем hash
	hash := params.Get("hash")
	if hash == "" {
		return nil, fmt.Errorf("hash not found in init data")
	}

	// Удаляем hash из параметров для проверки
	params.Del("hash")

	// Создаем строку для проверки
	dataCheckString := s.createDataCheckString(params)

	// Вычисляем секретный ключ
	secretKey := s.getSecretKey()

	// Проверяем подпись
	if !s.checkSignature(dataCheckString, hash, secretKey) {
		return nil, fmt.Errorf("invalid signature")
	}

	// Проверяем время (данные не старше 24 часов)
	authDateStr := params.Get("auth_date")
	if authDateStr == "" {
		return nil, fmt.Errorf("auth_date not found")
	}

	authDate, err := strconv.ParseInt(authDateStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid auth_date: %w", err)
	}

	if time.Now().Unix()-authDate > 86400 { // 24 часа
		return nil, fmt.Errorf("data is too old")
	}

	// Создаем структуру данных
	webAppData := &models.TelegramWebAppData{
		Hash: hash,
	}

	// Парсим остальные поля
	if userIDStr := params.Get("user"); userIDStr != "" {
		if userID, err := strconv.ParseInt(userIDStr, 10, 64); err == nil {
			webAppData.UserID = userID
		}
	}

	webAppData.Username = params.Get("username")
	webAppData.FirstName = params.Get("first_name")
	webAppData.LastName = params.Get("last_name")
	webAppData.AuthDate = authDate

	return webAppData, nil
}

// createDataCheckString создает строку для проверки подписи
func (s *TelegramAuthService) createDataCheckString(params url.Values) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var parts []string
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", key, params.Get(key)))
	}

	return strings.Join(parts, "\n")
}

// getSecretKey получает секретный ключ для проверки подписи
func (s *TelegramAuthService) getSecretKey() []byte {
	h := hmac.New(sha256.New, []byte("WebAppData"))
	h.Write([]byte(s.botToken))
	return h.Sum(nil)
}

// checkSignature проверяет подпись данных
func (s *TelegramAuthService) checkSignature(dataCheckString, hash string, secretKey []byte) bool {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataCheckString))
	expectedHash := hex.EncodeToString(h.Sum(nil))
	return hash == expectedHash
}
