#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log() {
    echo -e "${BLUE}[$(date +'%Y-%m-%d %H:%M:%S')]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1" >&2
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log "Начинаем обновление приложения..."

if [ ! -f ".env" ]; then
    error "Файл .env не найден. Создайте его на основе env.example"
    exit 1
fi

log "Получаем последние изменения из репозитория..."
git pull

log "Останавливаем сервисы..."
docker-compose down

log "Собираем и запускаем сервисы..."
docker-compose up -d --build

log "Ожидаем запуска сервисов..."
sleep 20

log "Проверяем статус сервисов..."
docker-compose ps

log "Проверяем доступность приложения..."

DOMAIN=$(grep "^DOMAIN=" .env | cut -d '=' -f2 | tr -d '"' | tr -d "'")

if [ -n "$DOMAIN" ]; then
    if curl -f -s "https://$DOMAIN/health" > /dev/null; then
        success "API доступен по адресу: https://$DOMAIN/health"
    else
        warning "API может быть недоступен. Проверьте логи: docker-compose logs api"
    fi

    if curl -f -s "https://$DOMAIN" > /dev/null; then
        success "Веб-приложение доступно по адресу: https://$DOMAIN"
    else
        warning "Веб-приложение может быть недоступно. Проверьте логи: docker-compose logs mini-app"
    fi
else
    warning "Домен не найден в .env файле. Проверьте доступность вручную."
fi

success "Обновление завершено!"
echo ""
echo "Полезные команды:"
echo "  docker-compose logs -f [service]  - просмотр логов сервиса"
echo "  docker-compose ps                 - статус всех сервисов"
echo "  docker-compose restart [service]  - перезапуск сервиса"