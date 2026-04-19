package service

import (
	"context"
	"database/sql"
)

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) CompleteTask(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, "UPDATE tasks SET completed = TRUE WHERE id = $1", id)
	return err
}
