package board

import "fmt"

type Board struct {
	Grid [][]rune
	Size int
}

func New(size int) *Board {
	grid := make([][]rune, size)

	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	return &Board{
		Grid: grid,
		Size: size,
	}
}

func (b *Board) Print() {
	for _, row := range b.Grid {
		fmt.Println(string(row))
	}
}
