package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresentBoard(t *testing.T) {
	occupied_cells := []cell{
		{0, 12, "🟢"},
		{1, 12, "⚪"},
		{2, 12, "🟤"},
		{3, 12, "🔵"},
		{4, 12, "⚫"},
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
	board := generateBoard(occupied_cells)
	renderedBoard := render(board)

	assert.Equal(t, expectedResult, renderedBoard)
}

func TestPresentBoardWithTwoRows(t *testing.T) {
	occupied_cells := []cell{
		{0, 12, "🟢"},
		{1, 12, "⚪"},
		{2, 12, "🟤"},
		{3, 12, "🔵"},
		{4, 12, "⚫"},
		{0, 11, "🔵"},
		{1, 11, "⚫"},
		{2, 11, "🟤"},
		{3, 11, "⚪"},
		{4, 11, "🟢"},
		{0, 0, "🟠"},
		{1, 0, "🟡"},
		{2, 0, "🟤"},
		{3, 0, "⚪"},
		{4, 0, "🟢"},
	}

	expectedResult := `|🟠|🟡|🟤|⚪|🟢|
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
|🔵|⚫|🟤|⚪|🟢|
|🟢|⚪|🟤|🔵|⚫|
`
	board := generateBoard(occupied_cells)
	renderedBoard := render(board)

	assert.Equal(t, expectedResult, renderedBoard)
}
