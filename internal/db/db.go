package db

import (
	"database/sql"
	"time"

	"github.com/e2e-test/taskify-api/internal/db/queries"
)

type Queries struct {
	db *sql.DB
}

func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

// Getters for null types
type NullString struct {
	String string
	Valid  bool
}

func (n NullString) Ptr() *string {
	if n.Valid {
		return &n.String
	}
	return nil
}

// Task status types
type TaskStatus string

const (
	TaskStatusPending  TaskStatus = "pending"
	TaskStatusActive   TaskStatus = "active"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusArchived TaskStatus = "archived"
)
