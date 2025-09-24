-- Инициализация базы данных для Telegram Bot
-- Создание таблицы пользователей

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для оптимизации
CREATE INDEX IF NOT EXISTS idx_users_user_id ON users(user_id);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);

-- Вставка тестовых данных (опционально)
-- INSERT INTO users (user_id, username, first_name, last_name) 
-- VALUES (123456789, 'test_user', 'Test', 'User')
-- ON CONFLICT (user_id) DO NOTHING;
