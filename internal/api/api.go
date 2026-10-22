package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dadams2/go-shorter/internal/store"
	"github.com/go-chi/chi/v5"
)

type URLHandler struct {
	store *store.Store
}

func NewURLHandler(urlStore *store.Store) *URLHandler {
	return &URLHandler{store: urlStore}
}

func (h *URLHandler) RecoverURL(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	url, err := h.store.RetrieveURL(r.Context(), token)
	if err != nil {
		http.Error(w, "could not find url", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	token, err := h.store.StoreURL(r.Context(), originalURL)
	if err != nil {
		http.Error(w, "could not shorten URL", http.StatusInternalServerError)
		return
	}

	shortURL := "http://localhost:8080/sr/" + token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
	}); err != nil {
		log.Printf("write shorten response: %v", err)
	}
}
