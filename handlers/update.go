package handlers

import (
	"time"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func CompleteTodo(manager *models.Manager, logger *models.Logger) {
	var id int
	isHere := false

	for {
		id = ReadIntInput("ID of Todo")
		for i, v := range manager.Todos {
			if id == v.ID {
				manager.Todos[i].IsDone = true
				logger.InfoLog.Printf("Todo with ID %v completed successfully!\n", id)
				isHere = true
				return
			}
		}

		if !isHere {
			logger.ErrorLog.Println("No todo with ID", id)
			return
		}
	}
}

func UpdateTodo(manager *models.Manager, logger *models.Logger) {
	var id int
	isHere := false
	var newTitle, newContent string
	var newPriority, newCategory int
	var newDueDate time.Time

	for {
		id = ReadIntInput("ID of Todo")
		logger.InfoLog.Println("If you don't want to change type 'x'")
		for i, v := range manager.Todos {
			if id == v.ID {
				isHere = true

				newTitle = ReadInput("Title", logger)
				newContent = ReadInput("Content", logger)

				newPriority = takePriority(1, logger)
				newCategory = takeCategory(1, manager, logger)

				newDueDate = takeDueDate(1, logger)

				if newTitle != "x" {
					manager.Todos[i].Title = newTitle
				}

				if newContent != "x" {
					manager.Todos[i].Content = newContent
				}

				if newPriority != 30092002 {
					manager.Todos[i].Priority = newPriority
				}

				if newCategory != 30092002 {
					manager.Todos[i].CategoryID = newCategory
				}

				today := time.Now()
				nowDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

				if newDueDate != nowDate.AddDate(2000, 10, 30) {
					manager.Todos[i].DueDate = newDueDate
				}

				logger.InfoLog.Printf("Todo with ID %v updated successfully!\n", id)
				return
			}
		}

		if !isHere {
			logger.ErrorLog.Println("No Todo with ID", id)
			return
		}
	}

}
