package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomColours(t *testing.T) {
	excludeList := []string{
		"green",
		"yellow",
		"blue",
		"orange",
		"green",
		"brown",
		"white",
	}

	assert.Equal(t, "black", getRandomColour(excludeList))
}

func TestCPUCCreatehallengeRow(t *testing.T) {
	challengeRow := CPUCreateChallengeRow()
	assert.Equal(t, 5, len(challengeRow))
	assert.Equal(t, 0, challengeRow[0].x)
	assert.Equal(t, 0, challengeRow[0].y)
}

func TestGenerateEmptyBoard(t *testing.T) {
	occupied_cells := make([]cell, 0)
	empty_board := generate_board(occupied_cells)

	assert.Equal(t, 0, empty_board[0].x)
	assert.Equal(t, 0, empty_board[0].y)
	assert.Equal(t, "", empty_board[0].colour)

	assert.Equal(t, 1, empty_board[21].x)
	assert.Equal(t, 4, empty_board[21].y)
	assert.Equal(t, "", empty_board[21].colour)

	assert.Equal(t, 4, empty_board[64].x)
	assert.Equal(t, 12, empty_board[64].y)
	assert.Equal(t, "", empty_board[64].colour)
}

func TestGenerateBoardWithOneColour(t *testing.T) {
	occupied_cells := []cell{
		{0, 0, "blue"},
	}

	empty_board := generate_board(occupied_cells)

	assert.Equal(t, 0, empty_board[0].x)
	assert.Equal(t, 0, empty_board[0].y)
	assert.Equal(t, "blue", empty_board[0].colour)

	assert.Equal(t, 1, empty_board[21].x)
	assert.Equal(t, 4, empty_board[21].y)
	assert.Equal(t, "", empty_board[21].colour)

	assert.Equal(t, 4, empty_board[64].x)
	assert.Equal(t, 12, empty_board[64].y)
	assert.Equal(t, "", empty_board[64].colour)
}

func TestGenerateBoardWithMultipleColours(t *testing.T) {
	occupied_cells := []cell{
		{0, 0, "green"},
		{1, 0, "white"},
		{2, 0, "brown"},
	}

	empty_board := generate_board(occupied_cells)

	assert.Equal(t, "green", empty_board[0].colour)
	assert.Equal(t, "white", empty_board[1].colour)
	assert.Equal(t, "brown", empty_board[2].colour)

	assert.Equal(t, 4, empty_board[64].x)
	assert.Equal(t, 12, empty_board[64].y)
	assert.Equal(t, "", empty_board[64].colour)
}

func TestFindCellsInRow(t *testing.T) {
	occupied_cells := []cell{}
	emptyBoard := generate_board(occupied_cells)
	expectedResult := []int{0, 1, 2, 3, 4}

	assert.Equal(t, expectedResult, findCellsInRow(emptyBoard, 0))
}
