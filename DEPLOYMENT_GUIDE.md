# 🚀 Руководство по развертыванию

## ✅ **Система полностью готова к работе!**

### 🏗️ **Архитектура системы:**

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Telegram Bot   │    │   Nginx Proxy   │    │   DeepSeek AI   │
│                 │    │                 │    │   (OpenRouter)  │
│ - /start        │    │ - /api/* → API  │    │                 │
│ - WebApp link   │    │ - /* → Mini App │    │                 │
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

## 🎯 **Что реализовано:**

### ✅ **Telegram Bot (Go)**

- Команда `/start` для регистрации пользователей
- Сохранение пользователей в PostgreSQL
- Приветственные сообщения

### ✅ **API Сервис (Gin)**

- **POST /api/chat** - чат с DeepSeek V3.1 (бесплатная модель)
- **GET /api/history** - история сообщений
- **GET /api/stats** - статистика пользователя
- **GET /health** - проверка здоровья
- Аутентификация через Telegram WebApp
- Лимит: 50 сообщений в день на пользователя

### ✅ **Mini App (React)**

- Современный чат интерфейс
- Интеграция с Telegram WebApp API
- Адаптивный дизайн для мобильных устройств
- Автопрокрутка и индикаторы загрузки

### ✅ **База данных (PostgreSQL)**

- Таблица `users` - пользователи бота
- Таблица `messages` - история чатов
- Автоматическое создание при запуске
- Оптимизированные индексы

### ✅ **Nginx Reverse Proxy**

- Маршрутизация API и Mini App
- CORS настройки для Telegram WebApp
- Единая точка входа на порту 80

### ✅ **pgAdmin**

- Веб-интерфейс для управления БД
- Просмотр таблиц и данных
- Выполнение SQL запросов

## 🔧 **Быстрый старт:**

### 1. **Получить токены:**

**Telegram Bot Token:**

1. Найдите [@BotFather](https://t.me/botfather) в Telegram
2. Отправьте `/newbot`
3. Следуйте инструкциям
4. Скопируйте токен

**OpenRouter API Key:**

1. Зарегистрируйтесь на [OpenRouter](https://openrouter.ai/)
2. Получите API ключ
3. Скопируйте ключ

### 2. **Настроить переменные окружения:**

```bash
# Отредактируйте файл .env
nano .env

# Добавьте ваши токены
BOT_TOKEN=1234567890:ABCdefGHIjklMNOpqrsTUVwxyz
OPENROUTER_API_KEY=sk-or-v1-your_openrouter_api_key_here
```

### 3. **Запустить систему:**

```bash
# Запуск всех сервисов
docker compose up -d

# Проверка статуса
docker compose ps
```

### 4. **Проверить работу:**

```bash
# Проверка API
curl http://localhost/health

# Проверка Mini App
open http://localhost/

# Проверка pgAdmin
open http://localhost:5050
```

## 🌐 **Доступные сервисы:**

| Сервис         | URL                   | Описание           |
| -------------- | --------------------- | ------------------ |
| **Mini App**   | http://localhost/     | Чат с ИИ           |
| **API**        | http://localhost/api/ | REST API           |
| **pgAdmin**    | http://localhost:5050 | Управление БД      |
| **PostgreSQL** | localhost:5433        | Прямое подключение |

### **Данные для pgAdmin:**

- **Email**: admin@admin.com
- **Пароль**: admin
- **Host**: postgres
- **Port**: 5432
- **Database**: telegram_bot
- **Username**: postgres
- **Password**: password

## 📊 **Мониторинг:**

### **Проверка статуса:**

```bash
# Статус всех сервисов
docker compose ps

# Логи API
docker compose logs api

# Логи бота
docker compose logs telegram-bot

# Логи PostgreSQL
docker compose logs postgres
```

### **Проверка базы данных:**

```bash
# Подключение к PostgreSQL
docker compose exec postgres psql -U postgres -d telegram_bot

# Просмотр таблиц
\dt

# Просмотр сообщений
SELECT * FROM messages ORDER BY created_at DESC LIMIT 10;
```

## 🔧 **Управление:**

### **Перезапуск сервисов:**

```bash
# Перезапуск конкретного сервиса
docker compose restart api

# Перезапуск всех сервисов
docker compose restart
```

### **Обновление кода:**

```bash
# Пересборка и перезапуск
docker compose up -d --build

# Пересборка конкретного сервиса
docker compose build api
docker compose up -d api
```

### **Просмотр логов:**

```bash
# Все логи
docker compose logs -f

# Логи конкретного сервиса
docker compose logs -f api
```

## 🚨 **Устранение неполадок:**

### **Проблема: API не отвечает**

```bash
# Проверить статус
docker compose ps api

# Проверить логи
docker compose logs api

# Проверить переменные окружения
docker compose exec api env | grep -E "(OPENROUTER|TELEGRAM)"
```

### **Проблема: База данных недоступна**

```bash
# Проверить статус PostgreSQL
docker compose ps postgres

# Проверить логи
docker compose logs postgres

# Проверить подключение
docker compose exec postgres pg_isready -U postgres
```

### **Проблема: Mini App не загружается**

```bash
# Проверить статус
docker compose ps mini-app

# Проверить nginx
docker compose ps nginx

# Проверить логи
docker compose logs nginx
```

## 📈 **Масштабирование:**

### **Горизонтальное масштабирование:**

```yaml
# В docker-compose.yml
api:
  deploy:
    replicas: 3
```

### **Добавление кэша Redis:**

```yaml
redis:
  image: redis:alpine
  networks:
    - app-network
```

### **Мониторинг с Prometheus:**

```yaml
prometheus:
  image: prom/prometheus
  ports:
    - "9090:9090"
```

## 🔒 **Безопасность:**

### **Рекомендации:**

1. **Изменить пароли** по умолчанию
2. **Настроить SSL** для production
3. **Ограничить доступ** к pgAdmin
4. **Настроить firewall**
5. **Регулярно обновлять** зависимости

### **Production настройки:**

```bash
# Изменить пароли в .env
POSTGRES_PASSWORD=strong_password_here
PGADMIN_DEFAULT_PASSWORD=strong_password_here

# Настроить SSL
# Добавить Let's Encrypt сертификаты
```

## 🎯 **Готово к использованию!**

Система полностью настроена и готова к работе:

1. ✅ **Telegram Bot** - регистрирует пользователей
2. ✅ **API** - обрабатывает чат с ИИ
3. ✅ **Mini App** - предоставляет интерфейс
4. ✅ **База данных** - хранит данные
5. ✅ **Nginx** - маршрутизирует запросы
6. ✅ **pgAdmin** - управляет БД

## 🤖 **Используемая модель ИИ:**

**DeepSeek V3.1 (Free)** - [deepseek/deepseek-chat-v3.1:free](https://openrouter.ai/deepseek/deepseek-chat-v3.1:free)

### **Характеристики модели:**

- ✅ **Полностью бесплатная** - $0 за входные и выходные токены
- 🧠 **671B параметров** - мощная модель с 37B активными параметрами
- 📝 **163,840 токенов контекста** - большой объем памяти для диалогов
- 🚀 **Быстрые ответы** - оптимизированная для скорости
- 🛠️ **Поддержка инструментов** - может использовать функции и API
- 💻 **Отличное программирование** - специализируется на коде

### **Возможности:**

- Гибридное мышление (thinking и non-thinking режимы)
- Структурированные вызовы инструментов
- Агенты для кодирования и поиска
- Подходит для исследований и сложных задач

**Следующий шаг**: Настроить токены в `.env` файле и начать использовать! 🚀
