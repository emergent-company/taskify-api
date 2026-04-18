package handler

import (
	"encoding/json"
	"net/http"

	"github.com/e2e-test/taskify-api/internal/service"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	store *service.TaskStore
}

// NewTaskHandler creates a new TaskHandler
func NewTaskHandler(store *service.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}

// HandleListTasks handles GET /api/v1/tasks requests
func (h *TaskHandler) HandleListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks := h.store.ListTasks()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
