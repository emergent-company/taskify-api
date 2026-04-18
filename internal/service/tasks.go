package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/e2e-test/taskify-api/internal/db"
)

type TaskService struct {
	db *db.TasksDB
}

func NewTaskService(db *db.TasksDB) *TaskService {
	return &TaskService{db: db}
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func (s *TaskService) CreateTask(ctx context.Context, input CreateTaskInput) (db.Task, error) {
	now := time.Now()
	id := uuid.New()

	task, err := s.db.InsertTask(ctx, db.InsertTaskParams{
		ID:          id,
		Title:       input.Title,
		Description: sql.NullString{String: input.Description, Valid: input.Description != ""},
		Status:      "pending",
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return db.Task{}, err
	}

	return task, nil
}
