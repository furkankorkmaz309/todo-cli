package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func ReadInput(prompt string, app *models.Logger) string {
	var text string
	for {
		fmt.Print(prompt + " : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text = scanner.Text()

		if strings.TrimSpace(text) == "" {
			app.ErrorLog.Println("Text can not be null")
		} else {
			return text
		}
	}
}

func ReadIntInput(prompt string) int {
	var myInt int
	fmt.Print(prompt + " : ")
	fmt.Scan(&myInt)
	return myInt
}

func takeCategory(question int, manager *models.Manager, app *models.Logger) int {

	listCategories(manager)
	for {
		var categoryID int
		var categoryIDstr string = ReadInput("Category ID", app)

		if question == 1 && categoryIDstr == "x" {
			return 30092002
		} else {
			categoryID, _ = strconv.Atoi(categoryIDstr)
		}

		if categoryID > len(manager.Categories) || categoryID < 1 {
			app.ErrorLog.Printf("Category ID should be in between 1 and %v\n", len(manager.Categories))
		} else {
			return categoryID
		}
	}
}

func takePriority(question int, app *models.Logger) int {
	for {
		var prioritystr string = ReadInput("Priority (1-3)", app)
		var priority int

		if question == 1 && prioritystr == "x" {
			return 30092002
		} else {
			priority, _ = strconv.Atoi(prioritystr)
		}

		if priority < 1 || priority > 3 {
			app.ErrorLog.Println("Priority should be in between 1 and 3!")
		} else {
			return priority
		}
	}
}

func takeDueDate(question int, app *models.Logger) time.Time {
	for {
		today := time.Now()
		nowDate := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

		dateInput := ReadInput("Date (YYYY-MM-DD)", app)
		if question == 1 && dateInput == "x" {
			return nowDate.AddDate(2000, 10, 30)
		}
		dueDate, err := time.Parse("2006-01-02", dateInput)

		if err != nil {
			app.ErrorLog.Println("Invalid date format!")
		} else if dueDate.Before(nowDate) {
			app.ErrorLog.Println("This day is before today!")
		} else {
			return dueDate
		}
	}
}
