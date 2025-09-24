# 🤖 API Documentation - Telegram Bot Chat

## 🎯 Обзор

API сервис для интеграции Telegram Mini App с DeepSeek AI через OpenRouter.

## 🏗️ Архитектура

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Telegram Bot   │    │   Gin API       │    │   OpenRouter    │
│                 │    │                 │    │                 │
│ - /start        │    │ - /api/chat     │    │ - DeepSeek AI   │
│ - WebApp link   │    │ - /api/history  │    │ - Chat API      │
│ - Notifications │    │ - /api/stats    │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   PostgreSQL    │
                    │                 │
                    │ - users         │
                    │ - messages      │
                    └─────────────────┘
```

## 🚀 Endpoints

### **POST /api/chat**

Отправляет сообщение в чат с ИИ.

**Headers:**

```
Content-Type: application/json
X-Telegram-Init-Data: <telegram_webapp_data>
```

**Request Body:**

```json
{
  "message": "Привет! Как дела?"
}
```

**Response:**

```json
{
  "message": "Привет! У меня все отлично, спасибо! Как дела у тебя?",
  "timestamp": "2025-09-24T12:00:00Z"
}
```

**Ошибки:**

- `400` - Неверный запрос
- `401` - Не авторизован
- `429` - Превышен лимит сообщений (50 в день)
- `500` - Внутренняя ошибка сервера

### **GET /api/history**

Получает историю сообщений пользователя.

**Headers:**

```
X-Telegram-Init-Data: <telegram_webapp_data>
```

**Response:**

```json
{
  "messages": [
    {
      "id": 1,
      "user_id": 123456789,
      "content": "Привет!",
      "role": "user",
      "created_at": "2025-09-24T12:00:00Z"
    },
    {
      "id": 2,
      "user_id": 123456789,
      "content": "Привет! Как дела?",
      "role": "assistant",
      "created_at": "2025-09-24T12:00:01Z"
    }
  ],
  "count": 2
}
```

### **GET /api/stats**

Получает статистику пользователя.

**Headers:**

```
X-Telegram-Init-Data: <telegram_webapp_data>
```

**Response:**

```json
{
  "daily_messages": 15,
  "daily_limit": 50,
  "remaining": 35
}
```

### **GET /health**

Проверка здоровья сервиса.

**Response:**

```json
{
  "status": "ok",
  "service": "telegram-api"
}
```

## 🔐 Аутентификация

### Telegram WebApp Authentication

API использует аутентификацию через Telegram WebApp API:

1. **Получение данных**: Mini App получает `initData` от Telegram
2. **Отправка в API**: Данные передаются в заголовке `X-Telegram-Init-Data`
3. **Валидация**: API проверяет подпись и время создания данных
4. **Извлечение пользователя**: Из данных извлекается информация о пользователе

### Пример использования в JavaScript:

```javascript
// Получение данных от Telegram WebApp
const tg = window.Telegram.WebApp;
const initData = tg.initData;

// Отправка запроса к API
const response = await fetch("/api/chat", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "X-Telegram-Init-Data": initData,
  },
  body: JSON.stringify({ message: "Привет!" }),
});
```

## 📊 Лимиты и ограничения

### Лимиты пользователей:

- **50 сообщений в день** на пользователя
- **300 символов** максимум в сообщении
- **10 сообщений** в контексте для ИИ

### Rate Limits:

- Пока не реализованы (планируется в будущем)

## 🗄️ База данных

### Таблица `messages`:

```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('user', 'assistant')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Индексы:

- `idx_messages_user_id` - по user_id
- `idx_messages_created_at` - по времени создания
- `idx_messages_user_role` - по user_id и role
- `idx_messages_user_date` - по user_id и дате

## 🔧 Конфигурация

### Переменные окружения:

```bash
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=telegram_bot

# OpenRouter
OPENROUTER_API_KEY=your_openrouter_api_key
OPENROUTER_URL=https://openrouter.ai/api/v1

# Telegram
TELEGRAM_BOT_TOKEN=your_telegram_bot_token

# API
API_PORT=8080
```

## 🚀 Развертывание

### Docker Compose:

```bash
# Запуск всех сервисов
docker compose up -d

# Проверка статуса
docker compose ps

# Логи API
docker compose logs api
```

### Доступные сервисы:

- **API**: http://localhost/api/
- **Mini App**: http://localhost/
- **pgAdmin**: http://localhost:5050
- **PostgreSQL**: localhost:5433

## 🧪 Тестирование

### Тест API:

```bash
# Health check
curl http://localhost/health

# Тест чата (требует Telegram WebApp данные)
curl -X POST http://localhost/api/chat \
  -H "Content-Type: application/json" \
  -H "X-Telegram-Init-Data: <telegram_data>" \
  -d '{"message": "Привет!"}'
```

### Тест через pgAdmin:

1. Открыть http://localhost:5050
2. Войти: admin@admin.com / admin
3. Подключиться к PostgreSQL
4. Проверить таблицы и данные

## 📝 Логирование

### Структурированные логи:

- Все запросы к API
- Ошибки OpenRouter
- Статистика использования
- Производительность

### Формат логов:

```
[2025/09/24 - 12:00:00] 192.168.1.1 POST 200 1.5s /api/chat
```

## 🔍 Мониторинг

### Health Checks:

- **API**: `/health` endpoint
- **Database**: Connection ping
- **OpenRouter**: API availability

### Метрики:

- Количество запросов
- Время ответа
- Ошибки
- Использование лимитов

## 🚨 Обработка ошибок

### Типы ошибок:

1. **Валидация**: Неверные данные запроса
2. **Аутентификация**: Неверные Telegram данные
3. **Лимиты**: Превышение дневного лимита
4. **OpenRouter**: Ошибки AI API
5. **База данных**: Ошибки подключения/запросов

### Коды ответов:

- `200` - Успех
- `400` - Неверный запрос
- `401` - Не авторизован
- `429` - Превышен лимит
- `500` - Внутренняя ошибка

## 🔄 Интеграция с Telegram Bot

### Команда /start:

```go
// В Telegram боте
msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Откройте Mini App для чата с ИИ.")
webApp := tgbotapi.WebAppInfo{URL: "https://yourdomain.com"}
keyboard := tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonWebApp("Открыть чат", webApp),
    ),
)
msg.ReplyMarkup = keyboard
```

## 🎯 Планы развития

### Ближайшие улучшения:

- [ ] Rate limiting
- [ ] Кэширование ответов
- [ ] Аналитика использования
- [ ] Экспорт диалогов
- [ ] Групповые чаты

### Долгосрочные планы:

- [ ] Множественные AI модели
- [ ] Персонализация
- [ ] Платные подписки
- [ ] API для разработчиков
