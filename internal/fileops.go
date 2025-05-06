package internal

import (
	"encoding/json"
	"os"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func SaveTodos(filename string, values []models.Todo) error {
	file, err := os.Create("../../data/" + filename)

	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(values)

	if err != nil {
		return err
	}

	return nil
}

func LoadTodos(filename string) ([]models.Todo, error) {
	file, err := os.Open("../../data/" + filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []models.Todo
	err = json.NewDecoder(file).Decode(&values)

	if err != nil {
		return nil, err
	}

	return values, nil
}
