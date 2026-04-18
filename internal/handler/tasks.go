package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/e2e-test/taskify-api/internal/db"
	"github.com/e2e-test/taskify-api/internal/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(db *db.DB) *TaskHandler {
	return &TaskHandler{
		taskService: service.NewTaskService(db),
	}
}

type CompleteTaskRequest struct {
	ID string `json:"id"`
}

func (h *TaskHandler) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path: /tasks/{id}/complete
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	path = strings.TrimSuffix(path, "/complete")
	id := strings.TrimPrefix(path, "/")

	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	// Complete the task
	updatedTask, err := h.taskService.CompleteTask(r.Context(), service.CompleteTaskInput{ID: id})
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return updated task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}
