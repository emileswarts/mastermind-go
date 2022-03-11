package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cell struct {
	x      int
	y      int
	colour string
}

func findCPUCells(board []cell) []cell {
	return findCellsInRow(board, 0)
}

func findCellsInRow(board []cell, rowNumber int) []cell {
	cellsInRow := make([]cell, 5)

	x := 0
	for _, cell := range board {
		if cell.y == rowNumber {
			cellsInRow[x] = cell
			x = x + 1
		}

	}

	return cellsInRow
}

func generateBoard(occupiedCells []cell) []cell {
	board := make([]cell, 65)
	x := 0
	y := 0

	for i := range board {
		if i == 0 {
			x = 0
			y = 0
		} else if i%5 == 0 {
			x = 0
			y = y + 1
		} else {
			x = x + 1
		}

		board[i].y = y
		board[i].x = x

		for _, cell := range occupiedCells {
			if board[i].x == cell.x && board[i].y == cell.y {
				board[i].colour = cell.colour
			}

		}
	}

	return board
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
	colourSet := [7]string{"blue", "yellow", "orange", "green", "brown", "white", "black"}
	availableColours := difference(colourSet, excludeList)

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

	challengeRow := []cell{
		{0, 0, colour1},
		{1, 0, colour2},
		{2, 0, colour3},
		{3, 0, colour4},
		{4, 0, colour5},
	}

	return challengeRow
}

func main() {
	occupiedCells := CPUCreateChallengeRow()
	board := generateBoard(occupiedCells)
	renderedBoard := render(board)

	fmt.Println(renderedBoard)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your 5 colours")
	fmt.Println("1 = 🔵 2 = 🟡 3 = 🟠 4 = 🟢 5 = 🟤 6 = ⚪ 7 = ⚫")
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}
