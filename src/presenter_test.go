package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGenerateBoardWithColours(t *testing.T) {
	occupied_cells := []cell{
		cell{0, 0, "blue"},
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
