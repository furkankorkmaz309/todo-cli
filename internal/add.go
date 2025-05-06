package internal

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func AddTodo(manager *models.Manager) {
	var newTodo models.Todo

	newTodo.ID = autoID(manager)
	fmt.Println("ID :", newTodo.ID)

	newTodo.Title = readInput("Title")

	newTodo.Content = readInput("Content")

	var priority int
	for {
		fmt.Print("Priority (1-5) : ")
		fmt.Scan(&priority)

		if priority < 1 || priority > 5 {
			fmt.Println("Priority should be in range 1 to 5!")
		} else {
			newTodo.Priority = priority
			break
		}
	}

	manager.Todos = append(manager.Todos, newTodo)
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
