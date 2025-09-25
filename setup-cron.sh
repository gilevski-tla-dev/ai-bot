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

log "Настраиваем автоматические задачи cron..."

CURRENT_DIR=$(pwd)

mkdir -p /var/log/ai-bot

log "Удаляем старые задачи cron..."
crontab -l 2>/dev/null | grep -v "$CURRENT_DIR" | crontab - || true

log "Добавляем новые задачи cron..."

TEMP_CRON=$(mktemp)

crontab -l 2>/dev/null | grep -v "$CURRENT_DIR" > "$TEMP_CRON" || true

cat >> "$TEMP_CRON" << EOF

0 3 * * 0 cd $CURRENT_DIR && ./renew-ssl.sh >> /var/log/ai-bot/ssl-renewal.log 2>&1

0 2 * * * cd $CURRENT_DIR && ./update.sh >> /var/log/ai-bot/update.log 2>&1

0 4 * * 0 cd $CURRENT_DIR && docker system prune -f >> /var/log/ai-bot/docker-cleanup.log 2>&1

EOF

crontab "$TEMP_CRON"

rm "$TEMP_CRON"

log "Проверяем установленные задачи cron..."
crontab -l | grep "$CURRENT_DIR"

success "Автоматические задачи cron настроены!"
echo ""
echo "Установленные задачи:"
echo "  - Обновление SSL сертификатов: каждое воскресенье в 3:00"
echo "  - Обновление приложения: каждый день в 2:00"
echo "  - Очистка Docker: каждое воскресенье в 4:00"
echo ""
echo "Логи доступны в:"
echo "  - /var/log/ai-bot/ssl-renewal.log"
echo "  - /var/log/ai-bot/update.log"
echo "  - /var/log/ai-bot/docker-cleanup.log"
echo ""
echo "Для просмотра логов:"
echo "  tail -f /var/log/ai-bot/ssl-renewal.log"
echo "  tail -f /var/log/ai-bot/update.log"
echo "  tail -f /var/log/ai-bot/docker-cleanup.log"