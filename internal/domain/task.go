// Package domain holds the core business entities and models for the go-task-api.
package domain

// Task represents a core business entity for a task in the system.
//
// It contains all the necessary information to track its lifecycle,
// such as Title, Description, and Status.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
