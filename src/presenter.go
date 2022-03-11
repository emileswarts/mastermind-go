package main

import "fmt"

func render(board []cell) string {
	var presented_board string

	for i, cell := range board {
		if i%5 == 0 && i != 0 {
			presented_board += "|\n"
		}

		if cell.colour != "" {
			presented_board += fmt.Sprintf("|%s", cell.colour)
		} else {
			presented_board += "|__"
		}
	}

	presented_board += "|\n"
	return presented_board
}
