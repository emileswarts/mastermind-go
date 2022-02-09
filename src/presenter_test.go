package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresentBoard(t *testing.T) {
	occupied_cells := []cell{
		{0, 0, "green"},
		{1, 0, "white"},
		{2, 0, "brown"},
	}

	expectedResult := ""
	board := generate_board(occupied_cells)
	renderedBoard := render(board)

	assert.Equal(t, expectedResult, renderedBoard)
}
