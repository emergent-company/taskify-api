package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/e2e-test/taskify-api/internal/db"
)

type TaskService struct {
	db *db.DB
}

func NewTaskService(db *db.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

type CompleteTaskInput struct {
	ID string `json:"id"`
}

func (s *TaskService) CompleteTask(ctx context.Context, input CompleteTaskInput) (*db.Task, error) {
	// Verify task exists first
	existingTask, err := s.GetTask(ctx, input.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}

	// Mark as complete
	now := time.Now()
	_, err = s.db.ExecContext(ctx, `
		UPDATE tasks
		SET completed = $2,
		    updated_at = $3
		WHERE id = $1
	`, input.ID, true, now)
	if err != nil {
		return nil, err
	}

	// Return updated task
	existingTask.Completed = true
	existingTask.UpdatedAt = &now
	return existingTask, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*db.Task, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`
	var task db.Task
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
