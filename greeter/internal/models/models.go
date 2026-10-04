package models

// Greeting is the response body for GET /hello.
type Greeting struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Health is the response body for GET /health.
type Health struct {
	Status string `json:"status"`
}

// Error is the single shape every error response uses.
type Error struct {
	Error string `json:"error"`
}
