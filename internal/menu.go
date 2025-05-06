package internal

import "fmt"

func Menu() int {
	fmt.Println("\n1 => Add Todo")
	fmt.Println("2 => List Todos")
	fmt.Println("3 => Update Todo")
	fmt.Println("4 => Delete Todo")
	fmt.Println("5 => Complete Todo")
	fmt.Println("0 => QUIT")
	fmt.Print("Choice : ")

	var option int
	fmt.Scan(&option)
	return option
}
