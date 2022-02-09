package main

func scoreRow(currentBoard []cell, rowToTest int) []string {
	var rowResult = make([]string, 5)
	cellsForRow := findCellsInRow(currentBoard, rowToTest)

	for i := 0; i < 5; i++ {
		if currentBoard[i].colour == "" {
			rowResult[0] = "CPU row not set"
		} else {
			if cellsForRow[i].colour == findCPUCells(currentBoard)[i].colour {
				rowResult[i] = "black"
			}
		}
	}

	return rowResult
}
