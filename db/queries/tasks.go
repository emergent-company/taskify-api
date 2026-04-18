package queries

import (
	"context"
	"database/sql"
)

type Task struct {
	ID          sql.NullString `json:"id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	Completed   sql.NullBool   `json:"completed"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type Queries struct {
	db *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

// GetTask retrieves a task by ID
func (q *Queries) GetTask(ctx context.Context, id string) (*Task, error) {
	row := q.db.QueryRowContext(ctx, `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks WHERE id = $1
	`, id)

	var t Task
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// CompleteTask updates a task to mark it as complete
func (q *Queries) CompleteTask(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, `
		UPDATE tasks
		SET completed = true, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`, id)
	return err
}
