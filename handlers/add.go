package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func AddTodo(manager *models.Manager, logger *models.Logger) {
	var newTodo models.Todo

	newTodo.ID = autoID(manager)
	fmt.Println("ID :", newTodo.ID)

	newTodo.Title = ReadInput("Title", logger)

	newTodo.Content = ReadInput("Content", logger)

	var priority int = takePriority(0, logger)
	newTodo.Priority = priority

	newTodo.CategoryID = takeCategory(0, manager, logger)

	newTodo.CreatedAt = time.Now()

	newTodo.DueDate = takeDueDate(0, logger)

	manager.Todos = append(manager.Todos, newTodo)
	logger.InfoLog.Println("Todo added successfully!")
}

func autoID(manager *models.Manager) int {
	myMap := make(map[int]bool)

	for _, v := range manager.Todos {
		myMap[v.ID] = true
	}

	for i := 1; ; i++ {
		if !myMap[i] {
			return i
		}
	}
}

func takeCategory(question int, manager *models.Manager, logger *models.Logger) int {

	listCategories(manager)
	for {
		var categoryID int
		var categoryIDstr string = ReadInput("Category ID", logger)

		if question == 1 && categoryIDstr == "x" {
			return 30092002
		} else {
			categoryID, _ = strconv.Atoi(categoryIDstr)
		}

		if categoryID > len(manager.Categories) || categoryID < 1 {
			logger.ErrorLog.Printf("Category ID should be in between 1 and %v\n", len(manager.Categories))
		} else {
			return categoryID
		}
	}
}

func takePriority(question int, logger *models.Logger) int {
	for {
		var prioritystr string = ReadInput("Priority (1-3)", logger)
		var priority int

		if question == 1 && prioritystr == "x" {
			return 30092002
		} else {
			priority, _ = strconv.Atoi(prioritystr)
		}

		if priority < 1 || priority > 3 {
			logger.ErrorLog.Println("Priority should be in between 1 and 3!")
		} else {
			return priority
		}
	}
}

func takeDueDate(question int, logger *models.Logger) time.Time {
	for {
		today := time.Now()
		nowDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

		dateInput := ReadInput("Date (YYYY-MM-DD)", logger)
		if question == 1 && dateInput == "x" {
			return nowDate.AddDate(2000, 10, 30)
		}
		dueDate, err := time.Parse("2006-01-02", dateInput)

		if err != nil {
			logger.ErrorLog.Println("Invalid date format!")
		} else if dueDate.Before(nowDate) {
			logger.ErrorLog.Println("This day is before today!")
		} else {
			return dueDate
		}
	}
}
