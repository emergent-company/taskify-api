package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Task represents a task in the database
type Task struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description sql.NullString `db:"description"`
	Completed   bool      `db:"completed"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Queries provides database query methods
type Queries struct {
	db *sql.DB
}

// NewQueries creates a new Queries instance
func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

// ListTasks returns all tasks from the database
func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, `
		SELECT 
			id, 
			title, 
			description, 
			completed, 
			created_at, 
			updated_at
		FROM tasks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
