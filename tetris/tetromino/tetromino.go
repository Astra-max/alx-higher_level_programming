package tetromino

import "errors"

type Tetromino struct {
	Grid  [][]rune
	Label rune
}

func New(lines []string, label rune) (*Tetromino, error) {
	if len(lines) != 4 {
		return nil, errors.New("invalid")
	}

	grid := make([][]rune, 4)
	count := 0

	for i := range lines {
		if len(lines[i]) != 4 {
			return nil, errors.New("invalid")
		}

		grid[i] = []rune(lines[i])

		for _, c := range grid[i] {
			if c == '#' {
				count++
			} else if c != '.' {
				return nil, errors.New("invalid")
			}
		}
	}

	// must be exactly 4 blocks
	if count != 4 {
		return nil, errors.New("invalid")
	}

	t := &Tetromino{
		Grid:  grid,
		Label: label,
	}

	t.normalize() // ⭐⭐⭐ CRITICAL FIX

	return t, nil
}

func (t *Tetromino) normalize() {
	minRow := 4
	minCol := 4

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if t.Grid[r][c] == '#' {
				if r < minRow {
					minRow = r
				}
				if c < minCol {
					minCol = c
				}
			}
		}
	}

	newGrid := make([][]rune, 4)
	for i := range newGrid {
		newGrid[i] = []rune{'.', '.', '.', '.'}
	}

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if t.Grid[r][c] == '#' {
				newGrid[r-minRow][c-minCol] = '#'
			}
		}
	}

	t.Grid = newGrid
}
