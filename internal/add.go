package internal

import (
	"fmt"
	"time"

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

	// list category kısmı güzel bir arayüz ile yapılacak!
	var categoryID int
	listCategories(manager)
	for {
		fmt.Print("Category ID : ")
		fmt.Scan(&categoryID)

		if categoryID > len(manager.Categories) || categoryID < 1 {
			fmt.Printf("Category ID should be in range 1 to %v\n", len(manager.Categories))
		} else {
			newTodo.Category = categoryID
			break
		}
	}

	newTodo.CreatedAt = time.Now().Format(time.RFC822)

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
