package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dadams2/go-shorter/internal/storage"
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter();

	r.Get("/healthcheck/", healthCheck)
	r.Get("/shorten/", shorten)
	r.Get("/sr/{token}", recoverURL)
	return r;

}

func healthCheck (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"msg": "test"})
	if err != nil {
		log.Println("there was an error")
	}
}

func recoverURL(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	url, err := storage.RetrieveURL(token);
    if err != nil {
        http.Error(w, "could not find url", http.StatusInternalServerError)
        return
    }
	http.Redirect(w, r, *url, http.StatusFound);
}


func shorten(w http.ResponseWriter, r *http.Request) {
    originalURL := r.URL.Query().Get("url")
    if originalURL == "" {
        http.Error(w, "url is required", http.StatusBadRequest)
        return
    }

    token, err := storage.StoreURL(originalURL)
    if err != nil {
        http.Error(w, "could not shorten URL", http.StatusInternalServerError)
        return
    }

    shortURL := "http://localhost:8080/sr/" + *token

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    if err := json.NewEncoder(w).Encode(map[string]string{
        "short_url": shortURL,
    }); err != nil {
        log.Printf("write shorten response: %v", err)
    }
}
