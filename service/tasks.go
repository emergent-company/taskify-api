package service

import (
	"database/sql"
)

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (s *TaskService) CompleteTask(id int) error {
	_, err := s.db.Exec("UPDATE tasks SET completed = TRUE WHERE id = $1", id)
	return err
}
