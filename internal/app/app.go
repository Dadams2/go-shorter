package app

import (
	"context"
	"database/sql"
	"os"

	"github.com/dadams2/go-shorter/database"
	"github.com/dadams2/go-shorter/internal/api"
	"github.com/dadams2/go-shorter/internal/store"
	"github.com/dadams2/go-shorter/migrations"
)

type Application struct {
	DB            *sql.DB
	HealthHandler *api.HealthHandler
	URLHandler    *api.URLHandler
}

func NewApplication() (*Application, error) {
	path := os.Getenv("DATABASE_PATH")
	if path == "" {
		path = "shorter.db"
	}

	// DATABASE OPEN AND MIGRATION
	db, err := database.Open(path)
	if err != nil {
		return nil, err
	}
	if err := database.MigrateFS(context.Background(), db, migrations.Migrations, "."); err != nil {
		db.Close()
		return nil, err
	}

	// SERVICE LAYER
	// not needed for this project

	// API HANDLER LAYER
	healthHandler := api.NewHealthHandler()
	urlHandler := api.NewURLHandler(store.New(db))

	app := &Application{
		DB:            db,
		HealthHandler: healthHandler,
		URLHandler:    urlHandler,
	}
	return app, nil
}
