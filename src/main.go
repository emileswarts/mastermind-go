package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type cell struct {
	x      int
	y      int
	colour string
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

func playUser() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	occupiedCells := CPUCreateChallengeRow()
	board := generateBoard(occupiedCells)

	renderedBoard := render(board)

	fmt.Println(renderedBoard)
	i := 0

	for {
		i = i + 1
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		board = tick(board, text)
		renderedBoard := render(board)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println(renderedBoard)
	}
}

func main() {
	playUser()
}
