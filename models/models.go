package models

import "time"

type Todo struct {
	ID        int
	Title     string
	Content   string
	Priority  int
	Category  string
	CreatedAt time.Time
	IsDone    bool
}
