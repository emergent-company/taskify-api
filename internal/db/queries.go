package db

import (
	"context"
	"database/sql"
)

const insertTask = `-- name: InsertTask :one
INSERT INTO tasks (title, description, status)
VALUES $1, $2, $3
RETURNING id, title, description, status, created_at, updated_at`

type InsertTaskParams struct {
	Title       string
	Description sql.NullString
	Status      string
}

type Queries struct {
	db *sql.DB
}

func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (Task, error) {
	var task Task
	err := q.db.QueryRowContext(ctx, insertTask, arg.Title, arg.Description, arg.Status).
		Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
	return task, err
}
