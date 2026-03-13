package task

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler handles HTTP requests for tasks.
type Handler struct {
	repo Repository
}

// NewHandler creates a new task handler.
func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// Routes manual routing for tasks
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL.Path)
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) list(w http.ResponseWriter, _ *http.Request) {
	tasks := h.repo.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created := h.repo.Create(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
