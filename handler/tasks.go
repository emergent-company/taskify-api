package handler

import (
	"encoding/json"
	"net/http"

	"github.com/e2e-test/taskify-api/service"
	"github.com/google/uuid"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	taskService *service.TaskService
}

// NewTaskHandler creates a new TaskHandler instance
func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// HandleListTasks handles GET /tasks requests
func (h *TaskHandler) HandleListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.taskService.ListTasks(r.Context())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// HandleTaskByID handles requests for individual tasks
type TaskByIDHandler struct {
	taskService *service.TaskService
}

// NewTaskByIDHandler creates a new TaskByIDHandler instance
func NewTaskByIDHandler(taskService *service.TaskService) *TaskByIDHandler {
	return &TaskByIDHandler{taskService: taskService}
}

// HandleTaskByID handles GET/PATCH /tasks/{id} requests
func (h *TaskByIDHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	path := r.URL.Path
	idStr := path[len("/tasks/"):]
	
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetTask(w, r, id)
	case http.MethodPatch:
		h.handleCompleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskByIDHandler) handleGetTask(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	// TODO: implement get single task
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *TaskByIDHandler) handleCompleteTask(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	// TODO: implement complete task
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
