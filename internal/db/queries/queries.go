package queries

import (
	"context"
	"database/sql"
	"time"

	"github.com/e2e-test/taskify-api/internal/db"
)

type InsertTaskParams struct {
	ID        interface{}
	Title     string
	Description sql.NullString
	Status    interface{}
	CreatedBy db.NullString
	OrgID     interface{}
}

type Task struct {
	ID          interface{}
	Title       string
	Description sql.NullString
	Status      interface{}
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	CreatedBy   db.NullString
	OrgID       interface{}
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (Task, error) {
	var task Task
	err := q.db.QueryRowContext(ctx, `
		INSERT INTO tasks (id, title, description, status, created_by, org_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, title, description, status, created_at, updated_at, created_by, org_id
	`, arg.ID, arg.Title, arg.Description, arg.Status, arg.CreatedBy, arg.OrgID).
		Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.CreatedBy,
			&task.OrgID,
		)
	return task, err
}
