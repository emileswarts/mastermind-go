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
	board := generateBoard(occupied_cells)
	renderedBoard := render(board)

	assert.Equal(t, expectedResult, renderedBoard)
}

func TestPresentBoardWithTwoRows(t *testing.T) {
	occupied_cells := []cell{
		{0, 12, "green"},
		{1, 12, "white"},
		{2, 12, "brown"},
		{3, 12, "blue"},
		{4, 12, "black"},
		{0, 11, "blue"},
		{1, 11, "black"},
		{2, 11, "brown"},
		{3, 11, "white"},
		{4, 11, "green"},
		{0, 0, "orange"},
		{1, 0, "yellow"},
		{2, 0, "brown"},
		{3, 0, "white"},
		{4, 0, "green"},
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
