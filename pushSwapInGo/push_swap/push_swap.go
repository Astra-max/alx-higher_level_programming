package main

/*
** Disclaimer:
** Student should use own implemented string.Split(...) and strconv.Atoi(...)
 */

import (
	"fmt"
	"os"
	"strconv" // NOT ALLOWED
	"strings" // NOT ALLOWED

	pushswap "../pushswap"
	stack "../stack"
	tools "../tools"
)

// Stack alias
type Stack = stack.Stack

// PushSwap alias
type PushSwap = pushswap.PushSwap

func sortSecondStack(PS *PushSwap, size int) int {
	if size == 2 {
		PS.SortFirstTwoB()
	} else if size == 3 {
		PS.SortFirstThreeB()
	}
	if size < 4 || PS.B.IsRevSorted(size) == true {
		return 0
	}
	topHalf := 0
	median := PS.B.Median(size)
	for i := 0; i < size; i++ {
		if PS.B.List.Number > median && topHalf >= 0 {
			PS.PushA()
		} else {
			PS.RotateB()
		}
		topHalf++
	}

	for i := 0; i < size-topHalf && PS.B.Length != size-topHalf; i++ {
		PS.RevRotateB()
	}
	pushed := topHalf + sortSecondStack(PS, size-topHalf)
	if PS.SyncedWithArray() == false {
		for i := 0; i < topHalf; i++ {
			PS.PushB()
		}
		pushed -= topHalf
	}
	return pushed + sortSecondStack(PS, topHalf)
}

func bestMove(stack *Stack, median int) int {
	i := 0
	best := -1
	mid := stack.Length / 2
	for tmpNode := stack.List; tmpNode != nil; tmpNode = tmpNode.Next {
		if tmpNode.Number < median && (tools.Abs(mid-best) < tools.Abs(mid-i) || best == -1) {
			best = i
		}
		i++
	}
	if best > mid {
		return 0
	}
	return 1
}

func insertionSort(PS *PushSwap, size int) {
	if size == 2 {
		PS.SortFirstTwoB()
	} else if size == 3 {
		PS.SortFirstThreeB()
	}
	if size < 4 && PS.B.IsRevSorted(size) == true {
		for i := 0; i < size; i++ {
			PS.PushA()
		}
		return
	}
	start := 0
	for i := 0; i < size; i++ {
		for PS.B.List.Number != PS.B.Max(start, size) {
			cmd := PS.B.GetBig()
			if cmd == 0 {
				PS.RotateB()
			} else {
				PS.RevRotateB()
			}
			if cmd == 0 {
				start--
			} else {
				start++
			}
			PS.PushA()
		}
	}
}

func sort(PS *PushSwap, size int) {
	if size == 2 {
		PS.SortFirstTwoA()
	} else if size == 3 {
		PS.SortFirstThreeA()
	}
	if size < 4 || PS.A.IsSorted(size) == true {
		return
	}
	topHalf := 0
	median := PS.A.Median(size)
	for PS.A.IsLess(PS.A.Length, median) {
		if PS.A.List.Number < median {
			PS.PushB()
			topHalf++
		} else {
			if bestMove(PS.A, median) == 1 {
				PS.RotateA()
			} else {
				PS.RevRotateA()
			}
		}
	}
	sort(PS, size-topHalf)
	if PS.Length < 150 {
		insertionSort(PS, topHalf)
		return
	}
	topHalf -= sortSecondStack(PS, topHalf)
	for topHalf > 0 {
		PS.PushA()
		topHalf--
	}
}

func main() {
	var PS PushSwap
	var numbers []string

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [args]\n", os.Args[0])
		fmt.Println("\targs - space separated numbers")
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
	PS.Print = 1
	tools.Qsort(PS.Array)
	if tools.CheckDuplicates(PS.Array) == false {
		tools.Exit()
	}
	sort(&PS, len(numbers))
}
