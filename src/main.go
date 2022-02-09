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

func findCellsInRow(board []cell, rowNumber int) []int {
	cellsInRow := make([]int, 5)

	x := 0
	for i, cell := range board {
		if cell.y == rowNumber {
			cellsInRow[x] = i
			x = x + 1
		}

	}

	return cellsInRow
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

func difference(a [7]string, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}

	}
	return diff
}

func getRandomColour(excludeList []string) string {
	defined_colours := [7]string{"blue", "yellow", "orange", "green", "brown", "white", "black"}
	availableColours := difference(defined_colours, excludeList)

	rand.Seed(time.Now().UnixNano())
	return availableColours[rand.Intn(len(availableColours))]
}

func CPUCreateChallengeRow() []cell {
	var selectedColours []string

	colour1 := getRandomColour(selectedColours)
	selectedColours = append(selectedColours, colour1)
	colour2 := getRandomColour(selectedColours)
	selectedColours = append(selectedColours, colour2)
	colour3 := getRandomColour(selectedColours)
	selectedColours = append(selectedColours, colour3)
	colour4 := getRandomColour(selectedColours)
	selectedColours = append(selectedColours, colour4)
	colour5 := getRandomColour(selectedColours)

	cpu_colour_choice := []cell{
		{0, 0, colour1},
		{1, 0, colour2},
		{2, 0, colour3},
		{3, 0, colour4},
		{4, 0, colour5},
	}

	return cpu_colour_choice
}

func main() {
	occupied_cells := CPUCreateChallengeRow()
	board := generate_board(occupied_cells)
	renderedBoard := render(board)

	fmt.Println(renderedBoard)
}
