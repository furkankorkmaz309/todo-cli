package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Priority  int    `json:"priority"`
	Category  int    `json:"category"`
	CreatedAt string `json:"created_at"`
	IsDone    bool   `json:"is_done"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Manager struct {
	Categories []Category
	Todos      []Todo
}
