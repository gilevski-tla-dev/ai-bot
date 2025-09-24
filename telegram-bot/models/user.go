package models

import (
	"time"
)

// User представляет пользователя Telegram бота
type User struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// UserRepository интерфейс для работы с пользователями в БД
type UserRepository interface {
	Save(user *User) error
}
