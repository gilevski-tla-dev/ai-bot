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

DOMAIN=${1:-"nikitagilevski.ru"}
EMAIL=${2:-"admin@nikitagilevski.ru"}

log "Начинаем развертывание для домена: $DOMAIN"
log "Email для SSL сертификата: $EMAIL"

if ! command -v docker &> /dev/null; then
    error "Docker не установлен. Установите Docker и попробуйте снова."
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    error "Docker Compose не установлен. Установите Docker Compose и попробуйте снова."
    exit 1
fi

if [ ! -f ".env" ]; then
    error "Файл .env не найден. Создайте его на основе env.example"
    exit 1
fi

log "Останавливаем существующие контейнеры..."
docker-compose down --remove-orphans || true

log "Создаем необходимые директории..."
mkdir -p ssl
mkdir -p certbot/conf
mkdir -p certbot/www

log "Обновляем конфигурацию nginx для домена: $DOMAIN"
sed -i.bak "s/nikitagilevski.ru/$DOMAIN/g" nginx.conf

cat > nginx-temp.conf << EOF
worker_processes auto;

events {
    worker_connections 1024;
}

http {
    include mime.types;
    default_type application/octet-stream;
    sendfile on;
    keepalive_timeout 65;

    server {
        listen 80;
        server_name $DOMAIN;

        location /.well-known/acme-challenge/ {
            root /var/www/certbot;
        }

        location / {
            return 301 https://\$server_name\$request_uri;
        }
    }
}
EOF

log "Запускаем временный nginx для получения SSL сертификата..."
docker run --rm -d \
    --name nginx-temp \
    -p 80:80 \
    -v $(pwd)/nginx-temp.conf:/etc/nginx/nginx.conf:ro \
    -v $(pwd)/certbot/www:/var/www/certbot \
    nginx:alpine

sleep 5

log "Получаем SSL сертификат через certbot..."
docker run --rm \
    -v $(pwd)/certbot/conf:/etc/letsencrypt \
    -v $(pwd)/certbot/www:/var/www/certbot \
    certbot/certbot certonly \
    --webroot \
    --webroot-path=/var/www/certbot \
    --email $EMAIL \
    --agree-tos \
    --no-eff-email \
    --force-renewal \
    -d $DOMAIN

log "Останавливаем временный nginx..."
docker stop nginx-temp || true

log "Копируем сертификаты..."
if [ -f "certbot/conf/live/$DOMAIN/fullchain.pem" ] && [ -f "certbot/conf/live/$DOMAIN/privkey.pem" ]; then
    cp certbot/conf/live/$DOMAIN/fullchain.pem ssl/cert.pem
    cp certbot/conf/live/$DOMAIN/privkey.pem ssl/key.pem
    success "SSL сертификаты успешно скопированы"
else
    error "Не удалось найти SSL сертификаты"
    exit 1
fi

log "Обновляем nginx.conf с путями к сертификатам..."
sed -i.bak "s|/etc/nginx/ssl/cert.pem|/etc/nginx/ssl/cert.pem|g" nginx.conf
sed -i.bak "s|/etc/nginx/ssl/key.pem|/etc/nginx/ssl/key.pem|g" nginx.conf

log "Собираем и запускаем все сервисы..."
docker-compose build --no-cache
docker-compose up -d

log "Ожидаем запуска сервисов..."
sleep 30

log "Проверяем статус сервисов..."
docker-compose ps

log "Проверяем доступность API..."
if curl -f -s "https://$DOMAIN/health" > /dev/null; then
    success "API доступен по адресу: https://$DOMAIN/health"
else
    warning "API может быть недоступен. Проверьте логи: docker-compose logs api"
fi

log "Проверяем доступность веб-приложения..."
if curl -f -s "https://$DOMAIN" > /dev/null; then
    success "Веб-приложение доступно по адресу: https://$DOMAIN"
else
    warning "Веб-приложение может быть недоступно. Проверьте логи: docker-compose logs mini-app"
fi

cat > renew-ssl.sh << 'EOF'
#!/bin/bash
set -e

DOMAIN=${1:-"nikitagilevski.ru"}

echo "Обновляем SSL сертификат для домена: $DOMAIN"

docker run --rm \
    -v $(pwd)/certbot/conf:/etc/letsencrypt \
    -v $(pwd)/certbot/www:/var/www/certbot \
    certbot/certbot renew

if [ -f "certbot/conf/live/$DOMAIN/fullchain.pem" ] && [ -f "certbot/conf/live/$DOMAIN/privkey.pem" ]; then
    cp certbot/conf/live/$DOMAIN/fullchain.pem ssl/cert.pem
    cp certbot/conf/live/$DOMAIN/privkey.pem ssl/key.pem
    
    docker-compose exec nginx nginx -s reload
    
    echo "SSL сертификат успешно обновлен и nginx перезагружен"
else
    echo "Ошибка: не удалось найти обновленные сертификаты"
    exit 1
fi
EOF

chmod +x renew-ssl.sh

log "Настраиваем автоматическое обновление сертификатов..."
(crontab -l 2>/dev/null; echo "0 3 * * 0 cd $(pwd) && ./renew-ssl.sh $DOMAIN >> /var/log/ssl-renewal.log 2>&1") | crontab -

success "Развертывание завершено успешно!"
echo ""
echo "Ваше приложение доступно по адресу: https://$DOMAIN"
echo "API доступен по адресу: https://$DOMAIN/api/"
echo "Health check: https://$DOMAIN/health"
echo ""
echo "Полезные команды:"
echo "  docker-compose logs -f [service]  - просмотр логов сервиса"
echo "  docker-compose ps                 - статус всех сервисов"
echo "  docker-compose restart [service]  - перезапуск сервиса"
echo "  ./renew-ssl.sh                    - обновление SSL сертификата"
echo ""
echo "SSL сертификаты будут автоматически обновляться каждое воскресенье в 3:00"