     1	package handler
     2	
     3	import (
     4		"net/http"
     5	
     6		"github.com/e2e-test/taskify-api/internal/service"
     7	)
     8	
     9	type TaskHandler struct {
    10		taskService *service.TaskService
    11	}
    12	
    13	func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
    14		return &TaskHandler{taskService: taskService}
    15	}
    16	
    17	func (h *TaskHandler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
    18		var input service.CreateTaskInput
    19		if err := parseJSON(r, &input); err != nil {
    20			http.Error(w, err.Error(), http.StatusBadRequest)
    21			return
    22		}
    23	
    24		task, err := h.taskService.CreateTask(r.Context(), input)
    25		if err != nil {
    26			http.Error(w, err.Error(), http.StatusInternalServerError)
    27			return
    28		}
    29	
    30		respondJSON(w, task, http.StatusCreated)
    31	}
