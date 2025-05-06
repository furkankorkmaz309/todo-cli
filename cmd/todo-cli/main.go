package main

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/internal"
	"github.com/furkankorkmaz309/todo-cli/models"
)

func main() {

	filename := "todos.json"
	fmt.Println("Hello World!")

	values, err := internal.LoadTodos(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	manager := &models.Manager{Todos: values}

	for {
		option := internal.Menu()
		switch option {
		case 0:
			err := internal.SaveTodos(filename, manager.Todos)

			if err != nil {
				fmt.Println(err)
				return
			}

			return
		case 1:
			fmt.Println("Adding")
			internal.AddTodo(manager)
		case 2:
			fmt.Println("Listing")
			fmt.Println(manager.Todos)
		case 3:
			fmt.Println("Updating")
		case 4:
			fmt.Println("Deleting")
		case 5:
			fmt.Println("Completing")
		default:
			fmt.Println("Wrong option! 1 to 5")
		}
	}
}
