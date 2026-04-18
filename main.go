package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/e2e-test/taskify-api/handler"
	"github.com/e2e-test/taskify-api/internal/db"
	"github.com/e2e-test/taskify-api/service"
)

func main() {
	// Initialize database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://localhost/taskify"
	}

	dbConn, err := db.NewDB(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// Initialize service
	svc := service.NewService(dbConn)

	// Initialize handler
	h := handler.NewHandler(svc)

	// Register routes
	http.HandleFunc("/tasks", h.HandleCreateTask)
	http.HandleFunc("/tasks/", h.HandleTaskByID)

	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
