package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreWhenCPUNotSet(t *testing.T) {
	var expectedResult []string

	occupied_cells := []cell{}
	board := generate_board(occupied_cells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[0], "CPU row not set")
}

func TestScoreWhenCPUSet(t *testing.T) {
	var expectedResult []string

	occupied_cells := []cell{
		{0, 0, "green"},
		{1, 0, "white"},
		{2, 0, "brown"},
		{3, 0, "blue"},
		{4, 0, "black"},
		{0, 12, "green"},
		{1, 12, "white"},
		{2, 12, "brown"},
		{3, 12, "blue"},
		{4, 12, "black"},
	}
	board := generate_board(occupied_cells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[0], "black")
	assert.Equal(t, expectedResult[1], "black")
	assert.Equal(t, expectedResult[2], "black")
	assert.Equal(t, expectedResult[3], "black")
	assert.Equal(t, expectedResult[4], "black")
}
