package internal

import (
	"encoding/json"
	"os"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func saveFile[T any](filename string, values []T) error {
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

func LoadFile[T any](filename string) ([]T, error) {
	file, err := os.Open("../../data/" + filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []T
	err = json.NewDecoder(file).Decode(&values)

	if err != nil {
		return nil, err
	}

	return values, nil
}

func SaveFiles(filenameTodo, filenameCategory string, manager *models.Manager) error {
	err := saveFile(filenameTodo, manager.Todos)

	if err != nil {
		return err
	}

	err = saveFile(filenameCategory, manager.Categories)

	if err != nil {
		return err
	}
	return nil
}
