package main

import (
	"log"
	"net/http"

	"greeter/internal/handlers"
)

const addr = ":9090"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.GetGreeting)
	mux.HandleFunc("GET /health", handlers.GetHealth)
	mux.HandleFunc("/", handlers.NotFound)

	log.Printf("greeter listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
