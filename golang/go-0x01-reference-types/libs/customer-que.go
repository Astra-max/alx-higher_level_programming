package libs

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func CustomerQue(que []string) {
	Instructions()
	for {
		comm := Input("Enter something")
		switch comm {
		case "add":
			customer := Input("Enter customer name")
			que = AddToQue(customer, que)
			break;
		case "show":
			ShowQue(que)
			break
		case "remove":
			que = Remove(que)
			break
		case "remove middle":
			message := ""
			position := Input("Choose customer")
			message, que = RemoveMiddle(position,que)
			fmt.Println(message)
			break
		case "clear que":
			que = ClearQue(que)
			break
		case "add many":
			customers := Input("Enter customer names")
			que = AddManyToQue(customers, que)
			break
		case "size":
			fmt.Printf("Number of waiting customers %v\n",QueSize(que))
			break
		case "exit":
			return
		default:
			fmt.Println("Invalid instruction!!")
			break
		}
	}
}

func Input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("[ %s ]: ", message)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func QueSize(que []string) int {
	return len(que)
}

func IsEmpty(que []string) bool {
	return len(que) == 0
}

func AddToQue(customer string, que []string) []string {
	que = append(que, customer)
	return que
}

func AddManyToQue(customers string, que []string) []string {
	tempQue := strings.Fields(customers)

	for _, customer := range tempQue {
		que = append(que, customer)
	}
	return que
}

func Remove(que []string) []string {
	if !IsEmpty(que) {
		fmt.Printf("Removing %s..\n", que[0])
		que = que[1:]
		return que
	}
	fmt.Println("Ooops!, no cutomer to serve today!")
	return que
}

func RemoveMiddle(index string, que []string) (string, []string) {
	position, err := strconv.Atoi(index)

	if err != nil {
		return "Invalid customer que position", nil
	}

	position--

	if position < 0 || position >= QueSize(que) {
		return "Failed to remove customer", nil
	}

	fmt.Printf("Removing %v\n", que[position])

	que = append(que[:position], que[position+1:]...)
	return "removed successfully", que
}

func ShowQue(que []string) {
	if IsEmpty(que) {
		fmt.Println("No customer que")
		return
	}

	for i, customer := range que {
		if i < len(que)-1 {
			fmt.Printf("[%v. %v] -> ", i+1, customer)
		} else {
			fmt.Printf("[%v. %v]", i+1, customer)
		}
	}
	fmt.Println()
}

func ClearQue(que []string) []string {
	return que[0:0]
}

func Instructions() {
	fmt.Println("Hello and welcome to customer Que!!")
	fmt.Printf("1. %s\n2. %s\n3. %v\n", "Add customer.", "Remove customer", "check que size")
	fmt.Printf("4. %s\n5. %s\n", "Show customers.", "Exit")
	fmt.Printf("6. %s\n7. %s\n", "Show que.", "Remove customer")

}