package main

import (
	"fmt"
	"math/rand"
)

type cell struct {
	x      int
	y      int
	colour string
}

func generate_board(occupied_cells []cell) []cell {
	default_cells := make([]cell, 65)
	x := 0
	y := 0

	for index := range default_cells {
		if index == 0 {
			x = 0
			y = 0
		} else if index%5 == 0 {
			x = 0
			y = y + 1
		} else {
			x = x + 1
		}

		default_cells[index].y = y
		default_cells[index].x = x
		default_cells[index].colour = ""
	}

	// for row_index := 0; row_index < game_board_max_rows; row_index++ {
	// 	for cell_index := 0; cell_index < game_board_max_cells_per_rows; cell_index++ {

	// 	}
	// for occupied_cell := range occupied_cells {
	// 	if default_cell.x == occupied_cell.x && default_cell.y == occupied_cell.y {
	// 		default_cells[default_cell_index] = occupied_cell
	// 	}

	// }
	// }

	return default_cells

}

func get_random_colour() string {
	available_colours := [7]string{"blue", "yellow", "orange", "green", "brown", "white", "black"}

	return available_colours[rand.Intn(len(available_colours))]
}

func CPUCreateChallengeRow() [5]cell {
	var cpu_colour_choice [5]cell
	const game_board_max_rows int = 11

	cpu_colour_choice[0] = cell{x: 0, y: game_board_max_rows, colour: get_random_colour()}
	cpu_colour_choice[1] = cell{x: 1, y: game_board_max_rows, colour: get_random_colour()}
	cpu_colour_choice[2] = cell{x: 2, y: game_board_max_rows, colour: get_random_colour()}
	cpu_colour_choice[3] = cell{x: 3, y: game_board_max_rows, colour: get_random_colour()}
	cpu_colour_choice[4] = cell{x: 4, y: game_board_max_rows, colour: get_random_colour()}

	return cpu_colour_choice
}

func main() {
	fmt.Println("Hello mastermind")
}
