package handlers

import (
	"fmt"
	"time"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func AddTodo(manager *models.Manager, app *models.Logger) {
	var newTodo models.Todo

	newTodo.ID = autoID(manager)
	fmt.Println("ID :", newTodo.ID)

	newTodo.Title = ReadInput("Title", app)

	newTodo.Content = ReadInput("Content", app)

	var priority int = takePriority(0, app)
	newTodo.Priority = priority

	newTodo.Category = takeCategory(0, manager, app)

	newTodo.CreatedAt = time.Now()

	newTodo.DueDate = takeDueDate(0, app)

	manager.Todos = append(manager.Todos, newTodo)
	app.InfoLog.Println("Todo added successfully!")
}

func autoID(manager *models.Manager) int {
	myMap := make(map[int]bool)

	for _, v := range manager.Todos {
		myMap[v.ID] = true
	}

	for i := 1; ; i++ {
		if !myMap[i] {
			return i
		}
	}
}
