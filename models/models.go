package models

import (
	"log"
	"time"
)

type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Priority   int       `json:"priority"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	DueDate    time.Time `json:"due_date"`
	IsDone     bool      `json:"is_done"`
}

type ArchivedTodo struct {
	Todo       Todo      `json:"todo"`
	ArchivedAt time.Time `json:"archived_at"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Manager struct {
	Categories    []Category
	Todos         []Todo
	ArchivedTodos []ArchivedTodo
}

type Logger struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

type TodoApp struct {
	Manager      *Manager
	Logger       *Logger
	TodoFile     string
	CategoryFile string
	ArchivedFile string
}
