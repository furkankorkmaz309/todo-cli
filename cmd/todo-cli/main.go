package main

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/internal"
	"github.com/furkankorkmaz309/todo-cli/models"
)

func main() {

	filenameTodo := "todos.json"
	filenameCategory := "categories.json"

	valuesTodo, err := internal.LoadFile[models.Todo](filenameTodo)

	if err != nil {
		fmt.Println(err)
		return
	}

	valuesCategory, err := internal.LoadFile[models.Category](filenameCategory)

	if err != nil {
		fmt.Println(err)
		return
	}

	manager := &models.Manager{
		Todos:      valuesTodo,
		Categories: valuesCategory,
	}

	for {
		option := internal.Menu()
		switch option {
		case 0:
			err = internal.SaveFiles(filenameTodo, filenameCategory, manager)
			if err != nil {
				fmt.Println(err)
			}
			return
		case 1:
			fmt.Println("Adding")
			internal.AddTodo(manager)
		case 2:
			fmt.Println("Listing")
			internal.ListTodos(manager.Todos)
		case 3:
			fmt.Println("Updating")
		case 4:
			fmt.Println("Deleting")
			internal.DeleteTodo(manager)
		case 5:
			fmt.Println("Completing")
			internal.CompleteTodo(manager)
		default:
			fmt.Println("Wrong option! 1 to 5")
		}
	}
}
