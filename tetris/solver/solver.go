package solver

import (
	"math"

	"tetris/board"
	"tetris/tetromino"
)

func Solve(pieces []*tetromino.Tetromino) *board.Board {
	size := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))

	for {
		b := board.New(size)

		if backtrack(b, pieces, 0) {
			return b
		}

		size++
	}
}

func backtrack(b *board.Board, pieces []*tetromino.Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}

	t := pieces[index]

	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {

			if canPlace(b, t, r, c) {
				place(b, t, r, c)

				if backtrack(b, pieces, index+1) {
					return true
				}

				remove(b, t, r, c)
			}
		}
	}

	return false
}

func canPlace(b *board.Board, t *tetromino.Tetromino, row, col int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Grid[i][j] == '#' {

				r := row + i
				c := col + j

				if r >= b.Size || c >= b.Size {
					return false
				}

				if b.Grid[r][c] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func place(b *board.Board, t *tetromino.Tetromino, row, col int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Grid[i][j] == '#' {
				b.Grid[row+i][col+j] = t.Label
			}
		}
	}
}

func remove(b *board.Board, t *tetromino.Tetromino, row, col int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if t.Grid[i][j] == '#' {
				b.Grid[row+i][col+j] = '.'
			}
		}			
	}
}