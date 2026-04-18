package service

import (
	"sync"
	"time"
)

// Task represents a task entity
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskStore is an in-memory task store
type TaskStore struct {
	mu   sync.RWMutex
	tasks map[int]*Task
	nextID int
}

// NewTaskStore creates a new in-memory task store
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// ListTasks retrieves all tasks ordered by creation date descending
func (s *TaskStore) ListTasks() []*Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	// Sort by created_at descending
	for i := 0; i < len(tasks)-1; i++ {
		for j := i + 1; j < len(tasks); j++ {
			if tasks[j].CreatedAt.After(tasks[i].CreatedAt) {
				tasks[i], tasks[j] = tasks[j], tasks[i]
			}
		}
	}

	return tasks
}

// AddTask adds a new task to the store
func (s *TaskStore) AddTask(title, description string) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &Task{
		ID:          s.nextID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.tasks[s.nextID] = task
	s.nextID++

	return task
}
