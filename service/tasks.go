package service

import (
	"context"
	"database/sql"

	"github.com/e2e-test/taskify-api/db/queries"
)

type TasksService struct {
	queries *queries.Queries
}

func NewTasksService(db *sql.DB) *TasksService {
	return &TasksService{
		queries: queries.New(db),
	}
}

func (s *TasksService) CompleteTask(ctx context.Context, id string) error {
	return s.queries.CompleteTask(ctx, id)
}

// GetTask retrieves a task by ID
func (s *TasksService) GetTask(ctx context.Context, id string) (*queries.Task, error) {
	return s.queries.GetTask(ctx, id)
}
