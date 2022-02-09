package main

import (
	"fmt"
	"math/rand"
	"time"
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

		for _, occupied_cell := range occupied_cells {
			if default_cells[index].x == occupied_cell.x && default_cells[index].y == occupied_cell.y {
				default_cells[index].colour = occupied_cell.colour
			}

		}
	}

	return default_cells

}

func get_random_colour() string {
	available_colours := [7]string{"blue", "yellow", "orange", "green", "brown", "white", "black"}

	rand.Seed(time.Now().UnixNano())
	return available_colours[rand.Intn(len(available_colours))]
}

func CPUCreateChallengeRow() []cell {
	cpu_colour_choice := []cell{
		{0, 0, get_random_colour()},
		{1, 0, get_random_colour()},
		{2, 0, get_random_colour()},
		{3, 0, get_random_colour()},
		{4, 0, get_random_colour()},
	}

	return cpu_colour_choice
}

func main() {
	occupied_cells := CPUCreateChallengeRow()
	board := generate_board(occupied_cells)
	renderedBoard := render(board)

	fmt.Println(renderedBoard)
}
