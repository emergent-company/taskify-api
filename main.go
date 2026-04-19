package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"taskify-api/handler"
	"taskify-api/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "tasks.db")
	taskService := &service.TaskService{DB: db}
	taskHandler := &handler.TaskHandler{Service: taskService}

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/complete") {
			taskHandler.HandleCompleteTask(w, r)
			return
		}
		http.Error(w, "not implemented", http.StatusNotImplemented)
	})
	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}
