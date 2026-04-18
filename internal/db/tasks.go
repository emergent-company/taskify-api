package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TasksDB struct {
	db *sql.DB
}

func NewTasksDB(db *sql.DB) *TasksDB {
	return &TasksDB{db: db}
}

const insertTask = `-- name: InsertTask :one
INSERT INTO tasks (id, title, description, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *`

type InsertTaskParams struct {
	ID          uuid.UUID
	Title       string
	Description sql.NullString
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Task struct {
	ID          uuid.UUID
	Title       string
	Description sql.NullString
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *TasksDB) InsertTask(ctx context.Context, arg InsertTaskParams) (Task, error) {
	var task Task
	err := q.db.QueryRowContext(ctx, insertTask,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	return task, err
}
