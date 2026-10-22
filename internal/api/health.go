package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"msg": "test"})
	if err != nil {
		log.Println("there was an error")
	}
}
