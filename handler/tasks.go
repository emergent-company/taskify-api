package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/e2e-test/taskify-api/service"
)

type TaskHandler struct {
	tasksService *service.TasksService
}

func NewTaskHandler(tasksService *service.TasksService) *TaskHandler {
	return &TaskHandler{
		tasksService: tasksService,
	}
}

// HandleCompleteTask handles PATCH /tasks/{id}/complete
func (h *TaskHandler) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract task ID from path: /tasks/{id}/complete
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id := strings.TrimSuffix(path, "/complete")
	if id == "" || id == path {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Complete the task
	err := h.tasksService.CompleteTask(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to complete task", http.StatusInternalServerError)
		return
	}

	// Get the updated task
	task, err := h.tasksService.GetTask(r.Context(), id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Return the updated task as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
