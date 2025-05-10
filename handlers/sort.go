package handlers

import (
	"sort"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func SortByID(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].ID < manager.Todos[j].ID
	})
	app.InfoLog.Println("Todos sorted by ID")
}

func SortByPriority(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].Priority < manager.Todos[j].Priority
	})
	app.InfoLog.Println("Todos sorted by Priority")
}

func SortByCategory(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].Category < manager.Todos[j].Category
	})
	app.InfoLog.Println("Todos sorted by Category")
}

func SortByCreateDate(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].CreatedAt.Before(manager.Todos[j].CreatedAt)
	})
	app.InfoLog.Println("Todos sorted by Creating Time")
}

func SortByDueDate(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].DueDate.Before(manager.Todos[j].DueDate)
	})
	app.InfoLog.Println("Todos sorted by Due Date")
}

func SortByIsDone(manager *models.Manager, app *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].IsDone
	})
	app.InfoLog.Println("Todos sorted by Finished to unfinished")
}

func Reverse(manager *models.Manager, app *models.Logger, s string) {
	optionSort := ReadInput("If you want "+s+", type '0'. Otherwise anything", app)
	if optionSort == "0" {
		for i, j := 0, len(manager.Todos)-1; i < j; i, j = i+1, j-1 {
			manager.Todos[i], manager.Todos[j] = manager.Todos[j], manager.Todos[i]
		}
	}
	app.InfoLog.Println("List succesfully turned to", s)
}
