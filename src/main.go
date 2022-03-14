package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
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

func getRandomColour(excludeList []string) string {
	colourSet := [7]string{"🔵", "🟡", "🟠", "🟢", "🟤", "⚪", "⚫"}
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
	var scores []string
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	occupiedCells := CPUCreateChallengeRow()
	board := generateBoard(occupiedCells)
	scores = scoreRow(board, 12)

	renderedBoard := render(board, scores)

	fmt.Println(renderedBoard)
	fmt.Println("Please enter your 5 colours")
	fmt.Println("🔵 🟡 🟠 🟢 🟤 ⚪ ⚫")
	for i := 12; i > 0; i-- {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		board = tick(board, text)

		scores = scoreRow(board, i)
		renderedBoard := render(board, scores)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println(renderedBoard)
		fmt.Println("Please enter your 5 colours")
		fmt.Println("🔵 🟡 🟠 🟢 🟤 ⚪ ⚫")
	}
}
