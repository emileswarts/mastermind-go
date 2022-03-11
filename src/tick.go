package main

import (
	"strings"
)

func nextPlayablePositions(board []cell) []cell {
	boardLength := len(board)
	for i := boardLength; i >= 0; i-- {
		if board[i-1].colour == "" {
			return board[i-5 : i]
		}
	}

	return board[boardLength-5 : boardLength]
}

func challengeHasBeenSet(challengeRow []cell) bool {
	if challengeRow[0].colour != "" {
		return true
	}

	return false
}

func findCPUChallengeRow(board []cell) []cell {
	return board[0:5]
}

func findOccupiedCells(board []cell) []cell {
	var result []cell

	for _, cell := range board {
		if cell.colour != "" && cell.y != 0 {
			result = append(result, cell)
		}
	}

	return result
}

func tick(board []cell, codeSubmission string) []cell {
	var occupiedCells []cell

	occupiedCells = findCPUChallengeRow(board)

	if challengeHasBeenSet(occupiedCells) == false {
		occupiedCells = CPUCreateChallengeRow()
	}

	for _, cell := range findOccupiedCells(board) {
		occupiedCells = append(occupiedCells, cell)
	}

	if codeSubmission != "" {
		codeSubmissionArray := strings.Split(codeSubmission, "")
		nextPlayablePositions := nextPlayablePositions(board)

		submission := []cell{
			{nextPlayablePositions[0].x, nextPlayablePositions[0].y, codeSubmissionArray[0]},
			{nextPlayablePositions[1].x, nextPlayablePositions[1].y, codeSubmissionArray[1]},
			{nextPlayablePositions[2].x, nextPlayablePositions[2].y, codeSubmissionArray[2]},
			{nextPlayablePositions[3].x, nextPlayablePositions[3].y, codeSubmissionArray[3]},
			{nextPlayablePositions[4].x, nextPlayablePositions[4].y, codeSubmissionArray[4]},
		}

		for _, cell := range submission {
			occupiedCells = append(occupiedCells, cell)
		}
	}

	// fmt.Println(render(generateBoard(occupiedCells)))
	return generateBoard(occupiedCells)
}
