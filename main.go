package main

import (
	"log"
	"net/http"

	task_management "SkillsForge-2025-Assessment/task-management"
)

func main() {
	store := task_management.NewTaskStore()
	server := &task_management.Server{Store: store}

	// Register routes
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			task_management.ListTasks(server, w, r)
		case http.MethodPost:
			task_management.CreateTask(server, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			task_management.GetTask(server, w, r)
		case http.MethodPut:
			task_management.UpdateTask(server, w, r)
		case http.MethodDelete:
			task_management.DeleteTask(server, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
