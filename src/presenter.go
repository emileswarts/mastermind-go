package main

import "fmt"

func render(board []cell) string {
	var presented_board string
	var colours map[string]string

	colours = map[string]string{
		"blue":   "🔵",
		"yellow": "🟡",
		"orange": "🟠",
		"green":  "🟢",
		"brown":  "🟤",
		"white":  "⚪",
		"black":  "⚫",
	}

	for i, cell := range board {
		if i%5 == 0 && i != 0 {
			presented_board += "|\n"
		}

		if cell.colour != "" {
			presented_board += fmt.Sprintf("|%s", colours[cell.colour])
		} else {
			presented_board += "|__"
		}
	}

	presented_board += "|\n"
	return presented_board
}
