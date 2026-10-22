package routes

import (
	"github.com/dadams2/go-shorter/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(application *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/healthcheck/", application.HealthHandler.HealthCheck)
	r.Get("/shorten/", application.URLHandler.Shorten)
	r.Get("/sr/{token}", application.URLHandler.RecoverURL)
	return r
}
