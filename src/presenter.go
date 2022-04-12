package main

func render(board []cell) string {
	var presentedBoard string

	for _, cell := range board {
		if cell.colour != "" {
			presentedBoard += "|" + cell.colour
		}

		if cell.colour == "" {
			presentedBoard += "|__"
		}

		if cell.x == 4 {
			if cell.y != 0 {
				rowScore := scoreRow(board, cell.y)

				if checkWinner(rowScore) {
					return "You are a winner!"
				}

				for _, scoreItem := range rowScore {
					presentedBoard += "|" + scoreItem
				}
			}
			presentedBoard += "|\n"
		}
	}

	presentedBoard += "Please enter your 5 colours\n"
	presentedBoard += "🔵 🟡 🟠 🟢 🟤 ⚪ ⚫"

	return presentedBoard
}
