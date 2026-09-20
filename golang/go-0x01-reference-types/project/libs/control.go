package libs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ControlPanel() {
	for {
		input := Input("Enter todo: ")
		switch input {
		case "add grades":
			fmt.Println("will implement add grades")
			break
		case "done":
			return
		case "class report":
			fmt.Println("will be implemented")
			break
		case "school report":
			fmt.Println("will implement school report")
			break
		case "individual report":
			fmt.Println("will implement individual student report")
			break
		default:
			fmt.Println("Invalid instruction")
		}
	}
}

func Input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s: ", message)
	scanner.Scan()
	data := strings.Trim(scanner.Text(), " ")
	return data
}
