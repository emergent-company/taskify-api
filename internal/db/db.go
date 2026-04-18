package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{DB: db}
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type Queries struct {
	*sql.DB
}

func (q *Queries) CompleteTask(ctx context.Context, id string) error {
	query := `
		UPDATE tasks
		SET completed = $2,
		    updated_at = $3
		WHERE id = $1
	`
	_, err := q.ExecContext(ctx, query, id, true, time.Now())
	return err
}
