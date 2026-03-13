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

func (r *memoryRepository) Create(t Task) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t
}

func (r *memoryRepository) List() []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.tasks
}
