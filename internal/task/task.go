package task

// Task represents a task in the system.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// Repository defines the contract for task persistence.
type Repository interface {
	Create(task Task) (Task, error)
	List() ([]Task, error)
	GetByID(id int) (Task, error)
	Update(id int, task Task) (Task, error)
	Delete(id int) error
}
