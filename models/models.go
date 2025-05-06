package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Priority  int       `json:"priority"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	IsDone    bool      `json:""is_done`
}

type Manager struct {
	Todos []Todo
}

var Todos []Todo
