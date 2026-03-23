// Package repository provides abstractions and implementations for data persistence.
package repository

import "github.com/ESousa97/go-task-api/internal/domain"

// TaskRepository defines the interface for task data access and persistence operations.
// It follows the Dependency Inversion principle, allowing the handlers to remain
// agnostic of the underlying database technology.
type TaskRepository interface {
	// Create persists a new task in the chosen data store.
	Create(task domain.Task) (domain.Task, error)
	// List retrieves all available tasks from the data store.
	List() ([]domain.Task, error)
	// GetByID finds a single task by its unique identifier.
	GetByID(id int) (domain.Task, error)
	// Update modifies an existing task identified by its ID.
	Update(id int, task domain.Task) (domain.Task, error)
	// Delete removes a task from the data store by its ID.
	Delete(id int) error
}
