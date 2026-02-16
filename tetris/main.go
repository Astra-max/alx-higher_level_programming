package main

import (
	"fmt"
	"os"

	"tetris/parser"
	"tetris/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	pieces, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	board := solver.Solve(pieces)

	if board == nil {
		fmt.Println("ERROR")
		return
	}

	board.Print()
}
