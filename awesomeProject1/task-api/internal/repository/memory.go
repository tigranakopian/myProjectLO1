package repository

import (
	"sync"
	"task-api/internal/domain"
)

type InMemoryRepo struct {
	mu    sync.RWMutex
	tasks map[string]domain.Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: make(map[string]domain.Task)}
}

func (r *InMemoryRepo) Create(task domain.Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
}

func (r *InMemoryRepo) GetAll(status *domain.Status) []domain.Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []domain.Task
	for _, t := range r.tasks {
		if status == nil || t.Status == *status {
			result = append(result, t)
		}
	}
	return result
}

func (r *InMemoryRepo) GetByID(id string) (domain.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.tasks[id]
	return t, ok
}
