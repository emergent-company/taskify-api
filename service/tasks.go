package service

import (
	"database/sql"
)

type TaskService struct {
	DB *sql.DB
}

func (s *TaskService) CompleteTask(id string) error {
	_, err := s.DB.Exec("UPDATE tasks SET completed = true WHERE id = ?", id)
	return err
}
