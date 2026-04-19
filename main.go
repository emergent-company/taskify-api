package main

import (
	"fmt"
	"net/http"

	"github.com/e2e-test/taskify-api/handler"
	"github.com/e2e-test/taskify-api/service"
)

func main() {
	// Initialize service and handler (nil db for now — wire real DB when available)
	taskSvc := service.NewTaskService(nil)
	taskHandler := handler.NewTaskHandler(taskSvc)

	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", taskHandler.HandleCompleteTask)

	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// TODO: implement POST /tasks (create task) and GET /tasks (list tasks)
func handleTasks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
