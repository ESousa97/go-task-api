package task

import "sync"

type memoryRepository struct {
	mu     sync.RWMutex
	tasks  []Task
	nextID int
}

// NewMemoryRepository creates a new in-memory task repository.
func NewMemoryRepository() Repository {
	return &memoryRepository{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (r *memoryRepository) Create(t Task) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t, nil
}

func (r *memoryRepository) List() ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.tasks, nil
}

func (r *memoryRepository) GetByID(id int) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, t := range r.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return Task{}, nil // Simulando 'not found' vazio (melhor seria um erro map)
}

func (r *memoryRepository) Update(id int, task Task) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, t := range r.tasks {
		if t.ID == id {
			task.ID = id
			r.tasks[i] = task
			return task, nil
		}
	}
	return Task{}, nil
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
