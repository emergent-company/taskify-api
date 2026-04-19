package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"taskify-api/handler"
	"taskify-api/service"
	_ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "postgres://user:pass@localhost/db?sslmode=disable")
	svc := service.NewTaskService(db)
	h := handler.NewTaskHandler(svc)

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/complete") {
			h.HandleCompleteTask(w, r)
			return
		}
		http.Error(w, "not implemented", http.StatusNotImplemented)
	})
	fmt.Println("taskify-api listening on :8080")
	http.ListenAndServe(":8080", nil)
}
