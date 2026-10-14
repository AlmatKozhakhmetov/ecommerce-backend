package http

import (
	"encoding/json"
	"net/http"
)

// PingHandler проверяет, живой ли сервер и база данных
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "pong! e-commerce backend is up and running",
	})
}
