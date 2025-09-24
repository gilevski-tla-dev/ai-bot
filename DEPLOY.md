# 🚀 Развертывание на VPS

Простая инструкция по развертыванию Telegram Mini App с DeepSeek AI.

## 📋 Быстрый старт

### 1. Подключитесь к VPS

```bash
ssh username@your-vps-ip
```

### 2. Клонируйте репозиторий

```bash
git clone https://github.com/your-username/ai-bot.git
cd ai-bot
```

### 3. Настройте переменные окружения

```bash
# Создайте .env файл
nano .env
```

**Содержимое .env:**

```bash
BOT_TOKEN=your_telegram_bot_token_here
OPENROUTER_API_KEY=your_openrouter_api_key_here
POSTGRES_PASSWORD=your_strong_database_password_here
PGADMIN_EMAIL=admin@yourdomain.com
PGADMIN_PASSWORD=your_strong_pgadmin_password_here
DOMAIN=your-domain.com
```

### 4. Запустите систему

```bash
docker compose up -d
```

### 5. Проверьте статус

```bash
docker compose ps
```

## 🔧 Настройка SSL (опционально)

### 1. Установите certbot

```bash
sudo apt install -y certbot
```

### 2. Получите сертификаты

```bash
sudo certbot certonly --standalone -d your-domain.com
```

### 3. Скопируйте сертификаты

```bash
sudo cp /etc/letsencrypt/live/your-domain.com/fullchain.pem ssl/cert.pem
sudo cp /etc/letsencrypt/live/your-domain.com/privkey.pem ssl/key.pem
sudo chown $USER:$USER ssl/cert.pem ssl/key.pem
```

### 4. Перезапустите nginx

```bash
docker compose restart nginx
```

## 📊 Управление

```bash
# Статус сервисов
docker compose ps

# Логи
docker compose logs -f

# Перезапуск
docker compose restart

# Остановка
docker compose down

# Обновление
docker compose up -d --build
```

## 🌐 Доступные сервисы

- **Mini App**: http://your-domain.com/
- **API**: http://your-domain.com/api/
- **pgAdmin**: http://your-domain.com:5050/ (локальный доступ)

## 🔒 Безопасность

### Firewall

```bash
sudo ufw allow 22
sudo ufw allow 80
sudo ufw allow 443
sudo ufw enable
```

### Автообновление SSL

```bash
(crontab -l 2>/dev/null; echo "0 12 * * * /usr/bin/certbot renew --quiet && docker compose restart nginx") | crontab -
```

## 🚨 Устранение неполадок

```bash
# Проверка логов
docker compose logs api
docker compose logs telegram-bot

# Проверка API
curl http://localhost/health

# Проверка базы данных
docker compose exec postgres psql -U postgres -d telegram_bot -c "SELECT COUNT(*) FROM users;"
```
