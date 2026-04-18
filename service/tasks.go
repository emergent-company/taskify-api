package service

import (
	"context"
	"database/sql"

	"github.com/e2e-test/taskify-api/internal/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func (s *Service) CreateTask(ctx context.Context, input CreateTaskInput) (db.Task, error) {
	var desc sql.NullString
	if input.Description != "" {
		desc = sql.NullString{String: input.Description, Valid: true}
	}
	
	q := db.NewQueries(s.db.DB)
	task, err := q.InsertTask(ctx, db.InsertTaskParams{
		Title:       input.Title,
		Description: desc,
		Status:      "pending",
	})
	if err != nil {
		return db.Task{}, err
	}
	return task, nil
}
