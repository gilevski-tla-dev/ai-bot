# 🗄️ pgAdmin - Управление базой данных

## 🚀 Доступ к pgAdmin

**URL**: http://localhost:5050

**Логин**: admin@admin.com  
**Пароль**: admin

## 📋 Настройка подключения к PostgreSQL

### 1. Войдите в pgAdmin

- Откройте http://localhost:5050
- Введите email: `admin@admin.com`
- Введите пароль: `admin`

### 2. Добавьте сервер PostgreSQL

1. **Правый клик** на "Servers" в левой панели
2. Выберите **"Register" → "Server..."**

### 3. Заполните данные подключения:

**General Tab:**

- **Name**: `Telegram Bot DB`

**Connection Tab:**

- **Host name/address**: `postgres`
- **Port**: `5432`
- **Maintenance database**: `telegram_bot`
- **Username**: `postgres`
- **Password**: `password`

### 4. Сохраните подключение

- Нажмите **"Save"**

## 🔍 Что можно делать в pgAdmin

### 📊 Просмотр данных

- **Таблицы**: `telegram_bot` → `Schemas` → `public` → `Tables` → `users`
- **Данные**: Правый клик на таблицу → **"View/Edit Data"** → **"All Rows"**

### 📈 SQL запросы

1. Выберите базу данных `telegram_bot`
2. Нажмите **"Query Tool"** (иконка SQL)
3. Введите SQL запросы:

```sql
-- Все пользователи
SELECT * FROM users;

-- Количество пользователей
SELECT COUNT(*) FROM users;

-- Последние 10 пользователей
SELECT * FROM users ORDER BY created_at DESC LIMIT 10;

-- Пользователи за сегодня
SELECT * FROM users WHERE DATE(created_at) = CURRENT_DATE;
```

### 🛠️ Управление базой

- **Создание таблиц**
- **Редактирование данных**
- **Выполнение SQL скриптов**
- **Экспорт/Импорт данных**
- **Мониторинг производительности**

## 🌐 Доступные сервисы

| Сервис         | URL                   | Описание                |
| -------------- | --------------------- | ----------------------- |
| **pgAdmin**    | http://localhost:5050 | Веб-интерфейс для БД    |
| **Mini App**   | http://localhost:3000 | React приложение        |
| **PostgreSQL** | localhost:5433        | Прямое подключение к БД |

## 🔧 Полезные функции pgAdmin

### 📋 Dashboard

- Общая статистика базы данных
- Активные подключения
- Размер базы данных

### 🔍 Query Tool

- Выполнение SQL запросов
- Просмотр результатов
- Экспорт результатов в CSV/JSON

### 📊 Statistics

- Статистика по таблицам
- Индексы и их использование
- Производительность запросов

## 🚨 Важные заметки

- **Данные сохраняются** в Docker volume `pgadmin-data`
- **Настройки подключений** сохраняются между перезапусками
- **Безопасность**: pgAdmin доступен только локально (localhost:5050)
- **Пароль**: можно изменить в переменных окружения docker-compose.yml

## 🐛 Решение проблем

### pgAdmin не запускается

```bash
docker compose logs pgadmin
docker compose restart pgadmin
```

### Не могу подключиться к PostgreSQL

- Проверьте, что PostgreSQL запущен: `docker compose ps postgres`
- Убедитесь, что используете правильные данные подключения
- Host должен быть `postgres` (имя сервиса в Docker), а не `localhost`
