package internal

import "github.com/furkankorkmaz309/todo-cli/models"

func NewManager() *models.Manager {
	return &models.Manager{
		Todos: []models.Todo{},
	}
}
