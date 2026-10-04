package handlers

import (
	"encoding/json"
	"net/http"

	"greeter/internal/models"
)

const defaultName = "World"

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// GetGreeting handles GET /hello.
func GetGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = defaultName
	}
	writeJSON(w, http.StatusOK, models.Greeting{
		Name:    name,
		Message: "Hello, " + name + "!",
	})
}

// GetHealth handles GET /health.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, models.Health{Status: "healthy"})
}

// NotFound answers any unmatched route with the error shape.
func NotFound(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotFound, models.Error{Error: "not found"})
}
