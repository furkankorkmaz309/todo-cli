package main

import (
	"fmt"

	"github.com/furkankorkmaz309/todo-cli/internal"
	"github.com/furkankorkmaz309/todo-cli/models"
)

func main() {
	filename := "deneme.json"
	fmt.Println("Hello World!")

	var denemeslice []models.Todo

	fmt.Println("denemeslice :", denemeslice)

	values, err := internal.LoadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		denemeslice = values
	}

	err = internal.SaveFile(filename, denemeslice)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("denemeslice :", denemeslice)
}
