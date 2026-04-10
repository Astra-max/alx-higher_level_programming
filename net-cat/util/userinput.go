package util

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func UserInput(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(message)
	scanner.Scan()
	return strings.Trim(scanner.Text(), "")
}