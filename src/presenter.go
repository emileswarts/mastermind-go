package main

import "fmt"

func render(board []cell, scores []string) string {
	var presented_board string

	for i, cell := range board {
		if i%5 == 0 && i != 0 {
			for _, score := range scores {
				presented_board += "| " + score
			}

			presented_board += "|\n"
		}

		if cell.colour != "" {
			presented_board += fmt.Sprintf("|%s", cell.colour)
		} else {
			presented_board += "|__"
		}
	}

	presented_board += "| | | | | |\n"
	return presented_board
}
