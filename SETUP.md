# 🚀 Быстрая настройка проекта

## ✅ Что уже готово:

- ✅ PostgreSQL база данных (порт 5433)
- ✅ Telegram бот (Go)
- ✅ API сервис (Gin) - http://localhost/api/ (DeepSeek V3.1 Free)
- ✅ Mini App (React) - http://localhost/
- ✅ Nginx reverse proxy
- ✅ Docker Compose конфигурация

## 🔧 Что нужно сделать:

### 1. Получить токены:

**Telegram Bot Token:**

1. Найдите [@BotFather](https://t.me/botfather) в Telegram
2. Отправьте `/newbot`
3. Следуйте инструкциям
4. Скопируйте полученный токен

**OpenRouter API Key:**

1. Зарегистрируйтесь на [OpenRouter](https://openrouter.ai/)
2. Получите API ключ
3. Скопируйте ключ

### 2. Настроить токены:

```bash
# Создайте файл .env
touch .env

# Добавьте ваши токены
echo "BOT_TOKEN=your_telegram_bot_token_here" >> .env
echo "OPENROUTER_API_KEY=your_openrouter_api_key_here" >> .env

# Опционально: настройте модель ИИ (по умолчанию используется бесплатная)
echo "AI_MODEL=deepseek/deepseek-chat-v3.1:free" >> .env
```

**Пример .env файла:**

```bash
BOT_TOKEN=1234567890:ABCdefGHIjklMNOpqrsTUVwxyz
OPENROUTER_API_KEY=sk-or-v1-your_openrouter_api_key_here
AI_MODEL=deepseek/deepseek-chat-v3.1:free
```

### 3. Запустить все сервисы:

```bash
# Запуск всех сервисов
docker compose up -d

# Проверка статуса
docker compose ps
```

### 4. Проверить работу:

```bash
# Логи API
docker compose logs api

# Логи бота
docker compose logs telegram-bot

# Проверка API
curl http://localhost/health
```

## 🌐 Доступные сервисы:

- **Mini App**: http://localhost/ (чат с ИИ)
- **API**: http://localhost/api/ (REST API)
- **pgAdmin**: http://localhost:5050 (логин: admin@admin.com, пароль: admin)
- **PostgreSQL**: localhost:5433
- **Telegram Bot**: найдите вашего бота в Telegram и отправьте `/start`

## 🐛 Если что-то не работает:

```bash
# Пересобрать все сервисы
docker compose down
docker compose up -d --build

# Посмотреть логи
docker compose logs -f
```
