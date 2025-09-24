package middleware

import (
	"fmt"
	"net/http"

	"telegram-api/services"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware для аутентификации через Telegram WebApp
func AuthMiddleware(telegramAuth *services.TelegramAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем данные аутентификации из заголовка
		initData := c.GetHeader("X-Telegram-Init-Data")
		if initData == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Telegram WebApp data",
			})
			c.Abort()
			return
		}

		// Валидируем данные
		webAppData, err := telegramAuth.ValidateWebAppData(initData)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid Telegram WebApp data",
				"details": err.Error(),
			})
			c.Abort()
			return
		}

		// Сохраняем данные пользователя в контекст
		c.Set("user_id", webAppData.UserID)
		c.Set("username", webAppData.Username)
		c.Set("first_name", webAppData.FirstName)
		c.Set("last_name", webAppData.LastName)

		c.Next()
	}
}

// CORSMiddleware middleware для CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Telegram-Init-Data, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// LoggingMiddleware middleware для логирования
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s %s %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.ClientIP,
			param.Method,
			param.StatusCode,
			param.Latency,
			param.Path,
			param.ErrorMessage,
		)
	})
}
