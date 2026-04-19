package service

import (
	"database/sql"
	"fmt"
	"time"
)

// Task represents a task record.
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskService provides task business logic.
type TaskService struct {
	db *sql.DB
}

// NewTaskService creates a new TaskService.
func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

// CompleteTask marks a task as completed and returns the updated task.
func (s *TaskService) CompleteTask(id string) (*Task, error) {
	row := s.db.QueryRow(
		`UPDATE tasks SET completed = TRUE, updated_at = CURRENT_TIMESTAMP WHERE id = ? RETURNING id, title, completed, created_at, updated_at`,
		id,
	)

	var t Task
	var createdAt, updatedAt string
	if err := row.Scan(&t.ID, &t.Title, &t.Completed, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found: %s", id)
		}
		return nil, fmt.Errorf("complete task: %w", err)
	}

	t.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	t.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &t, nil
}
