package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-task-api/internal/task"
	_ "github.com/lib/pq"
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

	repo := task.NewPostgresRepository(db)
	handler := task.NewHandler(repo)

	mux := http.NewServeMux()
	
	// Manual routing using ServeMux
	// Use trailing slash to allow Handler to do its own sub-routing (e.g. /tasks/1)
	mux.Handle("/tasks/", http.StripPrefix("", handler))
	mux.Handle("/tasks", handler)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
