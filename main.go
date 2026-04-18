package main

import (
	"fmt"
	"net/http"

	"github.com/e2e-test/taskify-api/internal/handler"
	"github.com/e2e-test/taskify-api/internal/service"
)

var taskStore *service.TaskStore

func main() {
	taskStore = service.NewTaskStore()
	taskHandler := handler.NewTaskHandler(taskStore)

	http.HandleFunc("/api/v1/tasks", taskHandler.HandleListTasks)
	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}
