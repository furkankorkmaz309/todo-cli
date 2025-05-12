package handlers

import (
	"time"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func DeleteTodo(manager *models.Manager, logger *models.Logger) {
	var id int
	var exTodo models.ArchivedTodo

	for {
		isHere := false
		id = ReadIntInput("ID of Todo")
		for i, v := range manager.Todos {
			if id == v.ID {
				isHere = true
				exTodo = models.ArchivedTodo{
					Todo: models.Todo{
						ID:         v.ID,
						Title:      v.Title,
						Content:    v.Content,
						Priority:   v.Priority,
						CategoryID: v.CategoryID,
						CreatedAt:  v.CreatedAt,
						DueDate:    v.DueDate,
						IsDone:     v.IsDone,
					},
					ArchivedAt: time.Now(),
				}

				manager.Todos = append(manager.Todos[:i], manager.Todos[i+1:]...)
				manager.ArchivedTodos = append(manager.ArchivedTodos, exTodo)

				logger.InfoLog.Printf("Todo with ID %v deleted successfully!\n", id)
				logger.InfoLog.Printf("Todo with ID %v archived successfully!\n", id)

				return
			}
		}

		if !isHere {
			logger.ErrorLog.Println("No todo with ID", id)
			return
		}
	}
}

func DeleteFinishedTodo(manager *models.Manager, logger *models.Logger) {
	var exTodo models.ArchivedTodo

	for i := len(manager.Todos) - 1; i >= 0; i-- {
		if manager.Todos[i].IsDone {
			exTodo = models.ArchivedTodo{
				Todo: models.Todo{
					ID:         manager.Todos[i].ID,
					Title:      manager.Todos[i].Title,
					Content:    manager.Todos[i].Content,
					Priority:   manager.Todos[i].Priority,
					CategoryID: manager.Todos[i].CategoryID,
					CreatedAt:  manager.Todos[i].CreatedAt,
					DueDate:    manager.Todos[i].DueDate,
					IsDone:     manager.Todos[i].IsDone,
				},
				ArchivedAt: time.Now(),
			}

			manager.ArchivedTodos = append(manager.ArchivedTodos, exTodo)
			manager.Todos = append(manager.Todos[:i], manager.Todos[i+1:]...)

			logger.InfoLog.Printf("Todo with ID %v archived successfully!\n", exTodo.Todo.ID)
			logger.InfoLog.Printf("Todo with ID %v deleted successfully!\n", exTodo.Todo.ID)

		}
	}
}
