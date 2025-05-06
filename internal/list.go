package internal

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func listCategories(manager *models.Manager) {
	fmt.Println(manager.Categories)
}

func ListTodos(T any) {
	fmt.Println(T)
}
