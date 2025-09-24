package database

import (
	"database/sql"
	"fmt"

	"telegram-bot/models"
)

// UserRepository реализует интерфейс models.UserRepository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository создает новый репозиторий пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Save сохраняет пользователя в базе данных
func (r *UserRepository) Save(user *models.User) error {
	query := `
		INSERT INTO users (user_id, username, first_name, last_name, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			username = EXCLUDED.username,
			first_name = EXCLUDED.first_name,
			last_name = EXCLUDED.last_name
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		user.UserID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
