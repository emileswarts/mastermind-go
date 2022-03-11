package main

func rightColourWrongPosition(cpuCell cell, cellsForRow []cell) bool {
	for _, cell := range cellsForRow {
		if cell.colour == cpuCell.colour {
			return true
		}
	}

	return false
}

func scoreRow(currentBoard []cell, rowToTest int) []string {
	var rowResult = make([]string, 5)
	cellsForRow := findCellsInRow(currentBoard, rowToTest)
	cpuCells := findCPUCells(currentBoard)

	for i, cpuCell := range cpuCells {
		if cpuCell.colour == "" {
			rowResult[0] = "CPU row not set"
		} else if cellsForRow[i].x == cpuCell.x && cellsForRow[i].colour == cpuCell.colour {
			rowResult[i] = "⬛"
		} else if rightColourWrongPosition(cpuCell, cellsForRow) {
			rowResult[i] = "⬜"
		}
	}

	return rowResult
}
