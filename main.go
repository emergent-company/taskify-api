package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/e2e-test/taskify-api/internal/db"
	"github.com/e2e-test/taskify-api/internal/handler"
	"github.com/e2e-test/taskify-api/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	// Get database URL from environment or use default
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://localhost/taskify"
	}

	// Connect to database
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	defer dbConn.Close()

	// Initialize services and handlers
	queries := db.NewQueries(dbConn)
	taskService := service.NewTaskService(queries)
	taskHandler := handler.NewTaskHandler(taskService)
	taskByIDHandler := handler.NewTaskByIDHandler(taskService)

	// Register routes
	http.HandleFunc("/tasks", taskHandler.HandleListTasks)
	http.HandleFunc("/tasks/", taskByIDHandler.HandleTaskByID)

	fmt.Println("taskify-api listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
