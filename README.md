# AI Bot

## Развертывание

1. Настройте переменные:

```bash
cp env.example .env
nano .env
```

2. Запустите развертывание:

```bash
./deploy.sh your-domain.com your-email@example.com
```

3. Настройте автоматические задачи:

```bash
./setup-cron.sh
```

## Обновление

```bash
git pull
./update.sh
```

## Переменные окружения

- `DOMAIN` - ваш домен
- `SSL_EMAIL` - email для SSL
- `BOT_TOKEN` - токен Telegram бота
- `POSTGRES_PASSWORD` - пароль БД
- `PGADMIN_PASSWORD` - пароль pgAdmin
- `OPENROUTER_API_KEY` - API ключ OpenRouter

## Доступ

- Веб-приложение: `https://your-domain.com`
- API: `https://your-domain.com/api/`
- Health Check: `https://your-domain.com/health`

## Команды

```bash
docker-compose logs -f
docker-compose ps
docker-compose restart api
./renew-ssl.sh
```
