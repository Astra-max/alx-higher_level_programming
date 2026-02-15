/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 16:04:30
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 16:29:52
 */

/*
** Disclaimer:
** Student should use own implemented string.Split(...) and strconv.Atoi(...)
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // NOW ALLOWED
	"strings" // NOW ALLOWED

	pushswap "../pushswap"
	tools "../tools"
)

// PushSwap alias
type PushSwap = pushswap.PushSwap

func command(PS *PushSwap, line string) {
	switch cmd := line; cmd {
	case "sa":
		PS.SwapA()
	case "sb":
		PS.SwapB()
	case "ss":
		PS.SwapAB()
	case "pa":
		PS.PushA()
	case "pb":
		PS.PushB()
	case "ra":
		PS.RotateA()
	case "rb":
		PS.RotateB()
	case "rr":
		PS.RotateAB()
	case "rra":
		PS.RevRotateA()
	case "rrb":
		PS.RevRotateB()
	case "rrr":
		PS.RevRotateAB()
	default:
		tools.Exit()
	}
}

func main() {
	var PS PushSwap
	var numbers []string
	if len(os.Args) != 2 {
		return
	}
	numbers = strings.Split(os.Args[1], " ")
	for _, num := range numbers {
		if tools.CheckNumber(num) == false {
			tools.Exit()
		}
		tmpNum, _ := strconv.Atoi(num)
		PS.Array = append(PS.Array, tmpNum)
	}
	PS.Init(len(numbers))
	tools.Qsort(PS.Array)
	if tools.CheckDuplicates(PS.Array) == false {
		tools.Exit()
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		command(&PS, string(line))
	}
	if PS.A.IsSorted(len(numbers)) && PS.B.Length == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
