package internal

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func DeleteTodo(manager *models.Manager) {
	var id int
	isHere := false

	for {
		fmt.Print("ID of Todo : ")
		fmt.Scan(&id)
		for i, v := range manager.Todos {
			if id == v.ID {
				manager.Todos = append(manager.Todos[:i], manager.Todos[i+1:]...)
				fmt.Printf("Todo with ID %v deleted!\n", id)
				isHere = true
				return
			}
		}

		if !isHere {
			fmt.Println("There is no todo with ID", id)
		}
	}

}

// buraya tamamlanmışların silinmesi fonksiyonu gelebilir
