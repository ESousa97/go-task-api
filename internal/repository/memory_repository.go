package repository

import (
	"sync"

	"github.com/ESousa97/apigotask/internal/domain"
)

type memoryRepository struct {
	mu     sync.RWMutex
	tasks  []domain.Task
	nextID int
}

// NewMemoryRepository creates a new in-memory task repository.
func NewMemoryRepository() TaskRepository {
	return &memoryRepository{
		tasks:  []domain.Task{},
		nextID: 1,
	}
}

func (r *memoryRepository) Create(t domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t, nil
}

func (r *memoryRepository) List() ([]domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.tasks, nil
}

func (r *memoryRepository) GetByID(id int) (domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, t := range r.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return domain.Task{}, nil // Simulando 'not found' vazio (melhor seria um erro map)
}

func (r *memoryRepository) Update(id int, task domain.Task) (domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, t := range r.tasks {
		if t.ID == id {
			task.ID = id
			r.tasks[i] = task
			return task, nil
		}
	}
	return domain.Task{}, nil
}

func (r *memoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
