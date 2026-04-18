package service

import (
	"context"

	"github.com/e2e-test/taskify-api/db"
)

// TaskService provides business logic for tasks
type TaskService struct {
	queries *db.Queries
}

// NewTaskService creates a new TaskService instance
func NewTaskService(queries *db.Queries) *TaskService {
	return &TaskService{queries: queries}
}

// ListTasks returns all tasks
func (s *TaskService) ListTasks(ctx context.Context) ([]db.Task, error) {
	return s.queries.ListTasks(ctx)
}
