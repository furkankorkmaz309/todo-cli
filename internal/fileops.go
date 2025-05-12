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

func loadFile[T any](filename string) ([]T, error) {
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

func saveFiles(filenameTodo, filenameCategory, filenameArchivedTodo string, manager *models.Manager) (string, error) {
	err := saveFile(filenameTodo, manager.Todos)

	if err != nil {
		return filenameTodo, err
	}

	err = saveFile(filenameCategory, manager.Categories)

	if err != nil {
		return filenameCategory, err
	}

	err = saveFile(filenameArchivedTodo, manager.ArchivedTodos)

	if err != nil {
		return filenameArchivedTodo, err
	}

	return "", nil
}

func loadFiles(filenameTodo, filenameCategory, filenameArchivedTodo string) ([]models.Todo, []models.Category, []models.ArchivedTodo, error) {
	valuesTodo, err := loadFile[models.Todo](filenameTodo)

	if err != nil {
		return nil, nil, nil, err
	}

	valuesCategory, err := loadFile[models.Category](filenameCategory)

	if err != nil {
		return nil, nil, nil, err
	}
	valuesarchivedTodo, err := loadFile[models.ArchivedTodo](filenameArchivedTodo)

	if err != nil {
		return nil, nil, nil, err
	}

	return valuesTodo, valuesCategory, valuesarchivedTodo, nil
}
