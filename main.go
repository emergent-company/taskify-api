package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/e2e-test/taskify-api/handler"
	"github.com/e2e-test/taskify-api/service"
)

func main() {
	// Initialize database connection
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/taskify?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize services
	tasksService := service.NewTasksService(db)

	// Initialize handlers
	taskHandler := handler.NewTaskHandler(tasksService)

	// Register routes
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		handleTaskByID(w, r, taskHandler)
	})

	fmt.Println("taskify-api listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// TODO: implement POST /tasks (create task)
		http.Error(w, "not implemented", http.StatusNotImplemented)
	} else if r.Method == http.MethodGet {
		// TODO: implement GET /tasks (list tasks)
		http.Error(w, "not implemented", http.StatusNotImplemented)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTaskByID(w http.ResponseWriter, r *http.Request, taskHandler *handler.TaskHandler) {
	if r.Method == http.MethodPatch {
		taskHandler.HandleCompleteTask(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
