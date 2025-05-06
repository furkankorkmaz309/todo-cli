package internal

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func CompleteTodo(manager *models.Manager) {
	var id int
	isHere := false

	for {
		fmt.Print("ID of Todo : ")
		fmt.Scan(&id)
		for i, v := range manager.Todos {
			if id == v.ID {
				manager.Todos[i].IsDone = true
				fmt.Printf("Todo with ID %v completed!\n", id)
				isHere = true
				return
			}
		}

		if !isHere {
			fmt.Println("There is no todo with ID", id)
		}
	}

}

// buraya tamamlandıktan sonra arşive alınması eklenebilir
