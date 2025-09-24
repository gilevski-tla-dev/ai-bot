-- Создание таблицы сообщений для чата с ИИ
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('user', 'assistant')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для оптимизации
CREATE INDEX IF NOT EXISTS idx_messages_user_id ON messages(user_id);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at);
CREATE INDEX IF NOT EXISTS idx_messages_user_role ON messages(user_id, role);
CREATE INDEX IF NOT EXISTS idx_messages_user_date ON messages(user_id, DATE(created_at));

-- Добавляем внешний ключ на таблицу users (если нужно)
-- ALTER TABLE messages ADD CONSTRAINT fk_messages_user_id 
-- FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE;
