package repository

import "go-task-api/internal/domain"

// TaskRepository defines the contract for task persistence.
type TaskRepository interface {
	Create(task domain.Task) (domain.Task, error)
	List() ([]domain.Task, error)
	GetByID(id int) (domain.Task, error)
	Update(id int, task domain.Task) (domain.Task, error)
	Delete(id int) error
}
