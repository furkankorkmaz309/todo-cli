package internal

import (
	"log"
	"os"

	"github.com/furkankorkmaz309/todo-cli/handlers"
	"github.com/furkankorkmaz309/todo-cli/models"
)

func NewApp() (*models.TodoApp, error) {
	filenameTodo := "todos.json"
	filenameCategory := "categories.json"
	filenameArchivedTodo := "archivedTodo.json"

	valuesTodo, valuesCategory, valuesArchivedTodo, err := loadFiles(filenameTodo, filenameCategory, filenameArchivedTodo)

	if err != nil {
		return nil, err
	}

	manager := &models.Manager{
		Todos:         valuesTodo,
		Categories:    valuesCategory,
		ArchivedTodos: valuesArchivedTodo,
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &models.Logger{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	todoApp := &models.TodoApp{
		Manager:      manager,
		Logger:       app,
		TodoFile:     filenameTodo,
		CategoryFile: filenameCategory,
		ArchivedFile: filenameArchivedTodo,
	}

	return todoApp, nil
}

func Run(app *models.TodoApp) {

	for {
		option := menu()
		switch option {
		case 0:
			filename, err := saveFiles(app.TodoFile, app.CategoryFile, app.ArchivedFile, app.Manager)
			if err != nil {
				app.Logger.ErrorLog.Println("An error occurred while saving the file with filename :", filename)
			}
			return
		case 1:
			handlers.AddTodo(app.Manager, app.Logger)
		case 2:
			handlers.ListTodos(app.Manager, app.Logger)
		case 3:
		sortLoop:
			for {
				optionTodo := sortMenu()
				switch optionTodo {
				case 0:
					app.Logger.InfoLog.Println("Turning to main menu...")
					break sortLoop
				case 1:
					handlers.SortByID(app.Manager, app.Logger)
				case 2:
					handlers.SortByPriority(app.Manager, app.Logger)
				case 3:
					handlers.SortByCategory(app.Manager, app.Logger)
				case 4:
					handlers.SortByCreateDate(app.Manager, app.Logger)
				case 5:
					handlers.SortByDueDate(app.Manager, app.Logger)
				case 6:
					handlers.SortByIsDone(app.Manager, app.Logger)
				default:
					app.Logger.ErrorLog.Println("Invalid choice! Please choose between 1 and 6")
				}

				handlers.Reverse(app.Manager, app.Logger, "descending")
				break

			}
			handlers.ListTodos(app.Manager, app.Logger)
		case 4:
			handlers.UpdateTodo(app.Manager, app.Logger)
		case 5:
			handlers.CompleteTodo(app.Manager, app.Logger)
		case 6:
			handlers.DeleteTodo(app.Manager, app.Logger)
		case 7:
			handlers.DeleteFinishedTodo(app.Manager, app.Logger)
		default:
			app.Logger.ErrorLog.Println("Invalid choice! Please choose between 1 and 7")
		}
	}
}
