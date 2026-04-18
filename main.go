package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTaskByID)
	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// TODO: implement POST /tasks (create task) and GET /tasks (list tasks)
func handleTasks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// TODO: implement PATCH /tasks/{id}/complete
func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
