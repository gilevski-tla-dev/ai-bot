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

DOMAIN=${1:-""}

if [ -z "$DOMAIN" ]; then
    if [ -f ".env" ]; then
        DOMAIN=$(grep "^DOMAIN=" .env | cut -d '=' -f2 | tr -d '"' | tr -d "'")
    fi
fi

if [ -z "$DOMAIN" ]; then
    error "Домен не указан. Используйте: ./renew-ssl.sh your-domain.com"
    exit 1
fi

log "Обновляем SSL сертификат для домена: $DOMAIN"

if [ ! -d "certbot/conf" ]; then
    error "Директория certbot/conf не найдена. Запустите сначала deploy.sh"
    exit 1
fi

log "Запускаем обновление сертификата через certbot..."
docker run --rm \
    -v $(pwd)/certbot/conf:/etc/letsencrypt \
    -v $(pwd)/certbot/www:/var/www/certbot \
    certbot/certbot renew --quiet

if [ -f "certbot/conf/live/$DOMAIN/fullchain.pem" ] && [ -f "certbot/conf/live/$DOMAIN/privkey.pem" ]; then
    log "Копируем обновленные сертификаты..."
    cp certbot/conf/live/$DOMAIN/fullchain.pem ssl/cert.pem
    cp certbot/conf/live/$DOMAIN/privkey.pem ssl/key.pem
    
    log "Перезагружаем nginx..."
    if docker-compose ps nginx | grep -q "Up"; then
        docker-compose exec nginx nginx -s reload
        success "SSL сертификат успешно обновлен и nginx перезагружен"
    else
        warning "Nginx не запущен. Сертификаты обновлены, но nginx не перезагружен"
    fi
else
    error "Не удалось найти обновленные сертификаты"
    exit 1
fi

if command -v openssl &> /dev/null; then
    EXPIRY_DATE=$(openssl x509 -in ssl/cert.pem -noout -dates | grep notAfter | cut -d= -f2)
    log "Сертификат действителен до: $EXPIRY_DATE"
fi

success "Обновление SSL сертификата завершено успешно!"