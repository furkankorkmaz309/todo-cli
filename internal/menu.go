package internal

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/handlers"
)

func menu() int {
	fmt.Println("\n1 => Add Todo")
	fmt.Println("2 => List Todos")
	fmt.Println("3 => Sort Todos")
	fmt.Println("4 => Update Todo")
	fmt.Println("5 => Complete Todo")
	fmt.Println("6 => Delete Todo")
	fmt.Println("7 => Delete Finished Todos")
	fmt.Println("0 => QUIT")

	option := handlers.ReadIntInput("\nChoice")

	return option
}

func sortMenu() int {
	fmt.Println("\n1 => ID")
	fmt.Println("2 => Priority")
	fmt.Println("3 => Category")
	fmt.Println("4 => Create Date")
	fmt.Println("5 => Due Date")
	fmt.Println("6 => Finished")
	fmt.Println("0 => BACK")

	fmt.Println("\nDefault : Ascending")
	option := handlers.ReadIntInput("Choice")

	return option
}
