# 🚀 Быстрая настройка проекта

## ✅ Что уже готово:

- ✅ PostgreSQL база данных (порт 5433)
- ✅ Telegram бот (Go)
- ✅ Mini App (React) - http://localhost:3000
- ✅ Docker Compose конфигурация

## 🔧 Что нужно сделать:

### 1. Получить токен Telegram бота:

1. Найдите [@BotFather](https://t.me/botfather) в Telegram
2. Отправьте `/newbot`
3. Следуйте инструкциям
4. Скопируйте полученный токен

### 2. Настроить токен:

```bash
# Отредактируйте файл .env
nano .env

# Замените your_telegram_bot_token_here на ваш токен
BOT_TOKEN=1234567890:ABCdefGHIjklMNOpqrsTUVwxyz
```

### 3. Перезапустить бота:

```bash
docker compose restart telegram-bot
```

### 4. Проверить работу:

```bash
# Статус всех сервисов
docker compose ps

# Логи бота
docker compose logs telegram-bot
```

## 🌐 Доступные сервисы:

- **pgAdmin**: http://localhost:5050 (логин: admin@admin.com, пароль: admin)
- **Mini App**: http://localhost:3000
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
