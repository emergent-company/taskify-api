package service

import (
	"context"
	"time"

	"github.com/e2e-test/taskify-api/internal/db"
	"github.com/e2e-test/taskify-api/internal/db/queries"
)

type CreateTaskInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status,omitempty"`
	CreatedBy   string  `json:"created_by"`
	OrgID       string  `json:"org_id"`
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	OrgID       string     `json:"org_id"`
}

type TaskService struct {
	queries *queries.Queries
}

func NewTaskService(queries *queries.Queries) *TaskService {
	return &TaskService{queries: queries}
}

func (s *TaskService) CreateTask(ctx context.Context, input CreateTaskInput) (*Task, error) {
	task, err := s.queries.InsertTask(ctx, queries.InsertTaskParams{
		ID:        input.Title, // Placeholder - should be UUID
		Title:     input.Title,
		OrgId:     input.OrgID,
		Status:    input.Status,
		CreatedBy: db.NullString{String: input.CreatedBy, Valid: input.CreatedBy != ""},
	})
	if err != nil {
		return nil, err
	}
	return &Task{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description.Ptr(),
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.Ptr(),
		UpdatedAt:   task.UpdatedAt.Ptr(),
		CreatedBy:   task.CreatedBy.Ptr(),
		OrgID:       task.OrgID.String(),
	}, nil
}
