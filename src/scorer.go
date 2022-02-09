package main

func scoreRow(currentBoard []cell, rowToTest int) []string {
	var rowResult = make([]string, 5)

	for i := 0; i < 5; i++ {
		if currentBoard[i].colour == "" {
			rowResult[0] = "CPU row not set"
		} else {
			rowResult[0] = "white"
			rowResult[1] = "white"
			rowResult[2] = "white"
			rowResult[3] = ""
			rowResult[4] = ""
		}
	}

	return rowResult
}
