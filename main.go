package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	mu      sync.Mutex
	tasks   []Task
	nextID  = 1
)

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskRouter)
	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		var t Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		t.ID = nextID
		nextID++
		tasks = append(tasks, t)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func taskRouter(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	if len(parts) < 1 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if len(parts) == 2 && parts[1] == "complete" {
		completeTaskHandler(w, r, parts[0])
		return
	}

	http.Error(w, "not found", http.StatusNotFound)
}

func completeTaskHandler(w http.ResponseWriter, r *http.Request, idStr string) {
	if r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	http.Error(w, "not found", http.StatusNotFound)
}
