package task

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	path := strings.TrimPrefix(r.URL.Path, "/tasks")

	// Tratar list e creation
	if path == "" || path == "/" {
		switch r.Method {
		case http.MethodGet:
			h.list(w, r)
		case http.MethodPost:
			h.create(w, r)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// Tratar GetByID, Update, Delete
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 1 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.get(w, r, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, r, id)
	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) list(w http.ResponseWriter, _ *http.Request) {
	tasks, err := h.repo.List()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.repo.Create(t)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request, id int) {
	task, err := h.repo.GetByID(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request, id int) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updated, err := h.repo.Update(id, t)
	if err != nil {
		if err.Error() == "task not found" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.Delete(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
