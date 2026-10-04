package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"greeter/internal/models"
)

func doGet(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", GetGreeting)
	mux.HandleFunc("GET /health", GetHealth)
	mux.ServeHTTP(rec, req)
	return rec
}

func decodeGreeting(t *testing.T, rec *httptest.ResponseRecorder) models.Greeting {
	t.Helper()
	var g models.Greeting
	if err := json.Unmarshal(rec.Body.Bytes(), &g); err != nil {
		t.Fatalf("decode greeting: %v", err)
	}
	return g
}

func TestGreetingWithName(t *testing.T) {
	rec := doGet(t, "/hello?name=Ada")
	g := decodeGreeting(t, rec)
	if g.Name != "Ada" || g.Message != "Hello, Ada!" {
		t.Fatalf("unexpected greeting: %+v", g)
	}
}

func TestGreetingWithoutName(t *testing.T) {
	rec := doGet(t, "/hello")
	g := decodeGreeting(t, rec)
	if g.Name != "World" {
		t.Fatalf("expected default World, got %+v", g)
	}
}

func TestGreetingWithEmptyName(t *testing.T) {
	rec := doGet(t, "/hello?name=")
	g := decodeGreeting(t, rec)
	if g.Name != "World" {
		t.Fatalf("expected default World, got %+v", g)
	}
}

func TestGreetingWithLongName(t *testing.T) {
	name := strings.Repeat("a", 2000)
	rec := doGet(t, "/hello?name="+name)
	g := decodeGreeting(t, rec)
	if g.Name != name {
		t.Fatalf("expected long name echoed in full, got len %d", len(g.Name))
	}
}

func TestGreetingWithUnicodeName(t *testing.T) {
	cases := []string{"José", "田中", "Ada🚀"}
	for _, name := range cases {
		rec := doGet(t, "/hello?name="+url.QueryEscape(name))
		g := decodeGreeting(t, rec)
		if g.Name != name {
			t.Fatalf("expected %q, got %q", name, g.Name)
		}
	}
}

func TestGreetingIgnoresExtraParams(t *testing.T) {
	rec := doGet(t, "/hello?name=Ada&foo=bar")
	g := decodeGreeting(t, rec)
	if g.Name != "Ada" {
		t.Fatalf("expected Ada, got %+v", g)
	}
}

func TestGreetingContentType(t *testing.T) {
	rec := doGet(t, "/hello?name=Ada")
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected application/json, got %q", ct)
	}
}

func TestHealth(t *testing.T) {
	rec := doGet(t, "/health")
	var h models.Health
	if err := json.Unmarshal(rec.Body.Bytes(), &h); err != nil {
		t.Fatalf("decode health: %v", err)
	}
	if h.Status == "" {
		t.Fatalf("expected non-empty status")
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected application/json, got %q", ct)
	}
}
