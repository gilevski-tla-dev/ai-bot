# Dockerfile для инициализации базы данных
FROM postgres:15-alpine

# Копируем SQL скрипты
COPY *.sql /docker-entrypoint-initdb.d/

# Устанавливаем права доступа
RUN chmod 644 /docker-entrypoint-initdb.d/*.sql

# PostgreSQL автоматически выполнит все .sql файлы из /docker-entrypoint-initdb.d/
# при первом запуске контейнера
