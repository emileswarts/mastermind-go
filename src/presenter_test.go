package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresentBoard(t *testing.T) {
	occupied_cells := []cell{
		{0, 12, "green"},
		{1, 12, "white"},
		{2, 12, "brown"},
		{3, 12, "blue"},
		{4, 12, "black"},
	}

	expectedResult := `|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|__|__|__|__|__|
|🟢|⚪|🟤|🔵|⚫|
`
	board := generate_board(occupied_cells)
	renderedBoard := render(board)

	assert.Equal(t, expectedResult, renderedBoard)
}
