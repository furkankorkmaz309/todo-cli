package handlers

import (
	"fmt"
	"os"
	"strconv"

	clr "github.com/fatih/color"
	"github.com/furkankorkmaz309/todo-cli/models"
	"github.com/olekukonko/tablewriter"
)

func listCategories(manager *models.Manager) {
	fmt.Println("")

	data := [][]string{}

	for _, v := range manager.Categories {
		data = append(data, []string{strconv.Itoa(v.ID), v.Icon, v.Name, v.Description})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Description", "Icon"})

	for _, v := range data {
		table.Append(v)
	}

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor}, // id
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor}, // name
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor}, // description
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor}) // icon

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgBlueColor}, // id
		tablewriter.Colors{tablewriter.FgBlueColor}, // name
		tablewriter.Colors{tablewriter.FgBlueColor}, // description
		tablewriter.Colors{tablewriter.FgBlueColor}) // icon

	table.SetRowLine(true)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("╪")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.Render()
}

func ListTodos(manager *models.Manager, logger *models.Logger) {
	fmt.Println("")

	data := [][]string{}

	for _, v := range manager.Todos {
		var isDone string

		if v.IsDone {
			isDone = "✓"
			isDone = clr.New(clr.FgHiGreen).Sprint(isDone)
		} else {
			isDone = "✗"
			isDone = clr.New(clr.FgHiRed).Sprint(isDone)
		}

		var priority string

		if v.Priority == 1 {
			priority = "1"
			priority = clr.New(clr.FgGreen).Sprint(priority)
		} else if v.Priority == 2 {
			priority = "2"
			priority = clr.New(clr.FgBlue).Sprint(priority)
		} else if v.Priority == 3 {
			priority = "3"
			priority = clr.New(clr.FgRed).Sprint(priority)
		} else {
			logger.ErrorLog.Println("something is wrong but how")
		}

		var categoryName string
		var icon string

		for _, vCategory := range manager.Categories {
			if v.CategoryID == vCategory.ID {
				categoryName = vCategory.Name
				icon = vCategory.Icon
				break
			}
		}

		category := fmt.Sprintf("%v ( %v )", categoryName, icon)

		data = append(data, []string{strconv.Itoa(v.ID), v.Title, v.Content, priority, category, v.DueDate.Format("2006-01-02"), isDone, v.CreatedAt.Format("2006-01-02 15:04")})

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Content", "Priority", "Category", "Due Date", "Done?", "Created At"})

	for _, v := range data {
		table.Append(v)
	}

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // id
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // title
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // content
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // priority
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // category
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // due date
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}, // is done
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgHiBlueColor, tablewriter.FgBlueColor}) // created at

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},    // id
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor}, // title
		tablewriter.Colors{}, // content
		tablewriter.Colors{}, // priority
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}, // category
		tablewriter.Colors{tablewriter.FgYellowColor},                  // due date
		tablewriter.Colors{tablewriter.Bold},                           // is done
		tablewriter.Colors{tablewriter.FgHiBlueColor})                  // created at

	table.Render()
}
