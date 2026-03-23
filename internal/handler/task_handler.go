// Package handler provides the HTTP entry points and routing for the application.
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-task-api/internal/domain"
	"go-task-api/internal/middleware"
	"go-task-api/internal/repository"
)

// TaskHandler implements the http.Handler interface to process task-related HTTP requests.
// It uses a [repository.TaskRepository] to perform data operations.
type TaskHandler struct {
	repo repository.TaskRepository
}

// NewTaskHandler initializes and returns a new pointer to [TaskHandler].
// It expects a valid [repository.TaskRepository] implementation.
func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// ServeHTTP manual routing for tasks
func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *TaskHandler) list(w http.ResponseWriter, r *http.Request) {
	if userID := r.Context().Value(middleware.UserIDKey); userID != nil {
		log.Printf("[TaskHandler/List] Fetching tasks for authorized user: %v", userID)
	}

	tasks, err := h.repo.List()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) create(w http.ResponseWriter, r *http.Request) {
	var t domain.Task
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

func (h *TaskHandler) get(w http.ResponseWriter, r *http.Request, id int) {
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

func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var t domain.Task
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

func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request, id int) {
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
