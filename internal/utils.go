package internal

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(str string) string {
	fmt.Print(str + " :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
