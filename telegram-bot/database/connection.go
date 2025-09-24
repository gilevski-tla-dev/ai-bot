package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection представляет подключение к базе данных
type Connection struct {
	db *sql.DB
}

// NewConnection создает новое подключение к базе данных
func NewConnection(host, port, user, password, dbname string) (*Connection, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Connection{db: db}, nil
}

// GetDB возвращает подключение к базе данных
func (c *Connection) GetDB() *sql.DB {
	return c.db
}

// Close закрывает подключение к базе данных
func (c *Connection) Close() error {
	return c.db.Close()
}

// CreateTables больше не нужен - таблицы создаются при инициализации PostgreSQL
// Этот метод оставлен для совместимости, но не выполняет никаких действий
func (c *Connection) CreateTables() error {
	// Таблицы создаются автоматически при инициализации PostgreSQL
	// через SQL скрипты в /docker-entrypoint-initdb.d/
	return nil
}
