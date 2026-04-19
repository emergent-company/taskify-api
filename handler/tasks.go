package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/e2e-test/taskify-api/service"
)

// TaskHandler handles HTTP requests for tasks.
type TaskHandler struct {
	svc *service.TaskService
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

// HandleCompleteTask handles PATCH /tasks/{id}/complete requests.
// It marks the specified task as completed and returns the updated task.
func (h *TaskHandler) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract task ID from path: /tasks/{id}/complete
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 || parts[0] != "tasks" || parts[2] != "complete" {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}
	id := parts[1]
	if id == "" {
		http.Error(w, "missing task id", http.StatusBadRequest)
		return
	}

	task, err := h.svc.CompleteTask(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
