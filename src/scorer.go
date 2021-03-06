package main

import (
	"sort"
)

func checkWinner(score []string) bool {
	winningCombo := []string{"⬛", "⬛", "⬛", "⬛", "⬛"}

	if len(score) != len(winningCombo) {
		return false
	}

	for i := range score {
		if score[i] != winningCombo[i] {
			return false
		}
	}
	return true
}

func rightColourWrongPosition(cpuCell cell, cellsForRow []cell) bool {
	for _, cell := range cellsForRow {
		if cell.colour == cpuCell.colour {
			return true
		}
	}

	return false
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

func scoreRow(currentBoard []cell, rowToTest int) []string {
	var rowResult = make([]string, 5)
	cellsForRow := findCellsInRow(currentBoard, rowToTest)
	cpuCells := findCPUCells(currentBoard)

	for i, cpuCell := range cpuCells {
		if cpuCell.colour == "" {
			rowResult[0] = "CPU row not set"
		} else if cellsForRow[i].x == cpuCell.x && cellsForRow[i].colour == cpuCell.colour {
			rowResult[cellsForRow[i].x] = "⬛"
		} else if rightColourWrongPosition(cpuCell, cellsForRow) {
			rowResult[cellsForRow[i].x] = "⬜"
		} else {
			rowResult[cellsForRow[i].x] = ""
		}
	}

	sort.Strings(rowResult)
	return rowResult
}
