package handlers

import (
	"sort"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func SortByID(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].ID < manager.Todos[j].ID
	})
	logger.InfoLog.Println("Todos sorted by ID")
}

func SortByPriority(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].Priority < manager.Todos[j].Priority
	})
	logger.InfoLog.Println("Todos sorted by Priority")
}

func SortByCategory(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].CategoryID < manager.Todos[j].CategoryID
	})
	logger.InfoLog.Println("Todos sorted by Category")
}

func SortByCreateDate(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].CreatedAt.Before(manager.Todos[j].CreatedAt)
	})
	logger.InfoLog.Println("Todos sorted by Creating Time")
}

func SortByDueDate(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].DueDate.Before(manager.Todos[j].DueDate)
	})
	logger.InfoLog.Println("Todos sorted by Due Date")
}

func SortByIsDone(manager *models.Manager, logger *models.Logger) {
	sort.Slice(manager.Todos, func(i, j int) bool {
		return manager.Todos[i].IsDone
	})
	logger.InfoLog.Println("Todos sorted by Finished to unfinished")
}

func Reverse(manager *models.Manager, logger *models.Logger, s string) {
	optionSort := ReadInput("If you want "+s+", type '0'. Otherwise anything", logger)
	if optionSort == "0" {
		for i, j := 0, len(manager.Todos)-1; i < j; i, j = i+1, j-1 {
			manager.Todos[i], manager.Todos[j] = manager.Todos[j], manager.Todos[i]
		}
	}
	logger.InfoLog.Println("List succesfully turned to", s)
}
