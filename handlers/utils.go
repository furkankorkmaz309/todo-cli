package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/furkankorkmaz309/todo-cli/models"
)

func ReadInput(prompt string, logger *models.Logger) string {
	var text string
	for {
		fmt.Print(prompt + " : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text = scanner.Text()

		if strings.TrimSpace(text) == "" {
			logger.ErrorLog.Println("Text can not be null")
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
