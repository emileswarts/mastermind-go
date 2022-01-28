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
