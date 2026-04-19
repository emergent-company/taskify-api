package handler

import (
	"net/http"
	"strings"
	"taskify-api/service"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Path: /tasks/{id}/complete
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[2] != "complete" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	id := parts[1]

	err := h.Service.CompleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
