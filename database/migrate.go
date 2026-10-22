package database

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	"github.com/pressly/goose/v3"
)

// Migrate applies pending migrations from a directory on disk.
func Migrate(ctx context.Context, db *sql.DB, dir string) error {
	return MigrateFS(ctx, db, os.DirFS(dir), ".")
}

// MigrateFS applies pending migrations from dir within fsys.
func MigrateFS(ctx context.Context, db *sql.DB, fsys fs.FS, dir string) error {
	migrations, err := fs.Sub(fsys, dir)
	if err != nil {
		return fmt.Errorf("open migrations: %w", err)
	}
	provider, err := goose.NewProvider(goose.DialectSQLite3, db, migrations)
	if err != nil {
		return fmt.Errorf("create migration provider: %w", err)
	}
	if _, err := provider.Up(ctx); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}
	return nil
}
