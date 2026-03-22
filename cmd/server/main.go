package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"go-task-api/internal/handler"
	"go-task-api/internal/middleware"
	"go-task-api/internal/repository"
)

func main() {
	connStr := "postgres://postgres:password@localhost:5433/taskdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not responsive: %v", err)
	}

	log.Println("Database connection established successfully!")

	repo := repository.NewPostgresRepository(db)
	taskHandler := handler.NewTaskHandler(repo)

	mux := http.NewServeMux()

	// Manual routing using ServeMux
	// Use trailing slash to allow Handler to do its own sub-routing (e.g. /tasks/1)
	mux.Handle("/tasks/", http.StripPrefix("", taskHandler))
	mux.Handle("/tasks", taskHandler)

	// Envolver o mux nos middlewares:
	// A ordem da execução de fora pra dentro será: Recovery -> Logger -> Auth
	handlerPipeline := middleware.Recovery(
		middleware.Logger(
			middleware.Auth(mux),
		),
	)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", handlerPipeline); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
