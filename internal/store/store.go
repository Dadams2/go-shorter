package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dadams2/go-shorter/internal/token"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) StoreURL(ctx context.Context, url string) (string, error) {
	token, err := token.GenerateToken()
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	query := "INSERT INTO urls (token, original_url) VALUES (?, ?)"
	if _, err := s.db.ExecContext(ctx, query, *token, url); err != nil {
		return "", fmt.Errorf("store URL: %w", err)
	}
	return *token, nil
}

// RetrieveURL returns an error wrapping sql.ErrNoRows when the token is unknown.
func (s *Store) RetrieveURL(ctx context.Context, token string) (string, error) {
	var url string
	query := "SELECT original_url FROM urls WHERE token = ?"
	if err := s.db.QueryRowContext(ctx, query, token).Scan(&url); err != nil {
		return "", fmt.Errorf("retrieve URL: %w", err)
	}
	return url, nil
}
