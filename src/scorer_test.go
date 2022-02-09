package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreWhenCPUNotSet(t *testing.T) {
	var expectedResult []string

	occupied_cells := []cell{
		{0, 12, "green"},
		{1, 12, "white"},
		{2, 12, "brown"},
		{3, 12, "blue"},
		{4, 12, "black"},
	}
	board := generate_board(occupied_cells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[0], "CPU row not set")
}
