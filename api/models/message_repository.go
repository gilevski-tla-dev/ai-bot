package models

import (
	"database/sql"
	"fmt"
)

// MessageRepositoryImpl реализует интерфейс MessageRepository
type MessageRepositoryImpl struct {
	db *sql.DB
}

// NewMessageRepository создает новый репозиторий сообщений
func NewMessageRepository(db *sql.DB) MessageRepository {
	return &MessageRepositoryImpl{db: db}
}

// Save сохраняет сообщение в базе данных
func (r *MessageRepositoryImpl) Save(message *Message) error {
	query := `
		INSERT INTO messages (user_id, content, role, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		message.UserID,
		message.Content,
		message.Role,
		message.CreatedAt,
	).Scan(&message.ID)

	if err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}

	return nil
}

// GetByUserID получает последние сообщения пользователя
func (r *MessageRepositoryImpl) GetByUserID(userID int64, limit int) ([]*Message, error) {
	query := `
		SELECT id, user_id, content, role, created_at
		FROM messages
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := r.db.Query(query, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}
	defer rows.Close()

	var messages []*Message
	for rows.Next() {
		message := &Message{}
		err := rows.Scan(
			&message.ID,
			&message.UserID,
			&message.Content,
			&message.Role,
			&message.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	// Разворачиваем порядок (от старых к новым)
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

// GetUserMessageCount возвращает количество сообщений пользователя за сегодня
func (r *MessageRepositoryImpl) GetUserMessageCount(userID int64) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM messages
		WHERE user_id = $1 
		AND role = 'user'
		AND DATE(created_at) = CURRENT_DATE
	`

	var count int
	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get message count: %w", err)
	}

	return count, nil
}
