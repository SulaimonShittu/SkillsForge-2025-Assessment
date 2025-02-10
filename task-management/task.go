package task_management

import (
	"sync"
	"time"
)

// Task represents a single task in the system
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
}

// TaskStore is our in-memory database
type TaskStore struct {
	sync.RWMutex
	Tasks  map[int]*Task
	NextID int
}

// NewTaskStore creates a new TaskStore
func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks:  make(map[int]*Task),
		NextID: 1,
	}
}

// Server wraps our dependencies
type Server struct {
	Store *TaskStore
}

// Response wrapper for consistent API responses
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginationParams handles pagination request parameters
type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
