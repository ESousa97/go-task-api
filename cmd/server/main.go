package main

import (
	"log"
	"net/http"
	"go-task-api/internal/task"
)

func main() {
	repo := task.NewMemoryRepository()
	handler := task.NewHandler(repo)

	mux := http.NewServeMux()
	
	// Manual routing using ServeMux
	// We point /tasks to our handler which manages methods
	mux.Handle("/tasks", handler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
