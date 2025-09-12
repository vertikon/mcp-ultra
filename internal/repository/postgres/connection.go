package postgres

import (
	"database/sql"
	"fmt"

	"github.com/vertikon/mcp-ultra/internal/config"
	_ "github.com/lib/pq"
)

// Connect creates a PostgreSQL database connection
func Connect(cfg config.PostgreSQLConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return db, nil
}