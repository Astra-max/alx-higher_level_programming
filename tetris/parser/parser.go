package parser

import (
	"bufio"
	"errors"
	"os"
	//"strings"

	"tetris/tetromino"
)

func ParseFile(path string) ([]*tetromino.Tetromino, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pieces []*tetromino.Tetromino
	scanner := bufio.NewScanner(f)

	block := []string{}
	label := 'A'

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if len(block) > 0 {
				t, err := tetromino.New(block, label)
				if err != nil {
					return nil, err
				}
				pieces = append(pieces, t)
				block = nil
				label++
			}
			continue
		}

		block = append(block, line)
	}

	if len(block) > 0 {
		t, err := tetromino.New(block, label)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, t)
	}

	if len(pieces) == 0 {
		return nil, errors.New("no pieces")
	}

	return pieces, nil
}
