package task_management

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreateTask(s *Server, w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Store.Lock()
	t.ID = s.Store.NextID
	t.CreatedAt = time.Now()
	s.Store.Tasks[t.ID] = &t
	s.Store.NextID++
	s.Store.Unlock()

	WriteJSON(w, Response{Data: t, Message: "Task created successfully"})
}

func ListTasks(s *Server, w http.ResponseWriter, r *http.Request) {
	page := GetIntParam(r, "page", 1)
	pageSize := GetIntParam(r, "pageSize", 10)
	category := r.URL.Query().Get("category")

	s.Store.RLock()
	defer s.Store.RUnlock()

	var filteredTasks []*Task
	for _, t := range s.Store.Tasks {
		if category == "" || t.Category == category {
			filteredTasks = append(filteredTasks, t)
		}
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(filteredTasks) {
		WriteJSON(w, Response{Data: []Task{}, Message: "No tasks found"})
		return
	}
	if end > len(filteredTasks) {
		end = len(filteredTasks)
	}

	WriteJSON(w, Response{
		Data:    filteredTasks[start:end],
		Message: fmt.Sprintf("Showing page %d of tasks", page),
	})
}

func UpdateTask(s *Server, w http.ResponseWriter, r *http.Request) {
	taskID := GetTaskIDFromURL(r)

	var updates Task
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Store.Lock()
	defer s.Store.Unlock()

	if t, exists := s.Store.Tasks[taskID]; exists {
		if updates.Title != "" {
			t.Title = updates.Title
		}
		if updates.Description != "" {
			t.Description = updates.Description
		}
		if updates.Category != "" {
			t.Category = updates.Category
		}
		if !updates.DueDate.IsZero() {
			t.DueDate = updates.DueDate
		}
		t.Completed = updates.Completed

		WriteJSON(w, Response{Data: t, Message: "Task updated successfully"})
		return
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteTask(s *Server, w http.ResponseWriter, r *http.Request) {
	taskID := GetTaskIDFromURL(r)

	s.Store.Lock()
	defer s.Store.Unlock()

	if _, exists := s.Store.Tasks[taskID]; exists {
		delete(s.Store.Tasks, taskID)
		WriteJSON(w, Response{Message: "Task deleted successfully"})
		return
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func GetTask(s *Server, w http.ResponseWriter, r *http.Request) {
	taskID := GetTaskIDFromURL(r)

	s.Store.RLock()
	defer s.Store.RUnlock()

	if t, exists := s.Store.Tasks[taskID]; exists {
		WriteJSON(w, Response{Data: t})
		return
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
