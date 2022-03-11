package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTick(t *testing.T) {
	occupiedCells := []cell{}
	board := generateBoard(occupiedCells)
	codeSubmission := ""
	nextTickBoard := tick(board, codeSubmission)

	assert.NotEqual(t, nextTickBoard[0].colour, "")
	assert.NotEqual(t, nextTickBoard[1].colour, "")
	assert.NotEqual(t, nextTickBoard[2].colour, "")
	assert.NotEqual(t, nextTickBoard[3].colour, "")
	assert.NotEqual(t, nextTickBoard[4].colour, "")
	assert.Equal(t, nextTickBoard[5].colour, "")
}

func TestWhenThePlayerHasPlayedTick(t *testing.T) {
	occupiedCells := []cell{
		{0, 0, "🟢"},
		{1, 0, "⚪"},
		{2, 0, "🟤"},
		{3, 0, "🔵"},
		{4, 0, "⚫"},
	}
	board := generateBoard(occupiedCells)
	codeSubmission := "⚫⚪🟤🟡🟠"
	nextTickBoard := tick(board, codeSubmission)

	assert.NotEqual(t, nextTickBoard[0].colour, "")
	assert.NotEqual(t, nextTickBoard[1].colour, "")
	assert.NotEqual(t, nextTickBoard[2].colour, "")
	assert.NotEqual(t, nextTickBoard[3].colour, "")
	assert.NotEqual(t, nextTickBoard[4].colour, "")
	assert.Equal(t, nextTickBoard[5].colour, "")
}

func TestWhenThePlayerIsPlayingForTheSecondTimeTick(t *testing.T) {
	var nextTickBoard []cell

	occupiedCells := []cell{
		{0, 0, "🟢"},
		{1, 0, "⚪"},
		{2, 0, "🟤"},
		{3, 0, "🔵"},
		{4, 0, "⚫"},
	}

	board := generateBoard(occupiedCells)
	codeSubmission := "⚫⚪🟤🟡🟠"
	nextTickBoard = tick(board, codeSubmission)

	assert.Equal(t, nextTickBoard[0].colour, "🟢")
	assert.Equal(t, nextTickBoard[1].colour, "⚪")
	assert.Equal(t, nextTickBoard[2].colour, "🟤")
	assert.Equal(t, nextTickBoard[3].colour, "🔵")
	assert.Equal(t, nextTickBoard[4].colour, "⚫")

	assert.Equal(t, nextTickBoard[60].colour, "⚫")
	assert.Equal(t, nextTickBoard[61].colour, "⚪")
	assert.Equal(t, nextTickBoard[62].colour, "🟤")
	assert.Equal(t, nextTickBoard[63].colour, "🟡")
	assert.Equal(t, nextTickBoard[64].colour, "🟠")

	codeSubmission = "🔵🟢🟡🟤⚪"
	nextTickBoard = tick(nextTickBoard, codeSubmission)

	assert.Equal(t, nextTickBoard[0].colour, "🟢")
	assert.Equal(t, nextTickBoard[1].colour, "⚪")
	assert.Equal(t, nextTickBoard[2].colour, "🟤")
	assert.Equal(t, nextTickBoard[3].colour, "🔵")
	assert.Equal(t, nextTickBoard[4].colour, "⚫")

	assert.Equal(t, nextTickBoard[59].colour, "⚪")
	assert.Equal(t, nextTickBoard[58].colour, "🟤")
	assert.Equal(t, nextTickBoard[57].colour, "🟡")
	assert.Equal(t, nextTickBoard[56].colour, "🟢")
	assert.Equal(t, nextTickBoard[55].colour, "🔵")

	assert.Equal(t, nextTickBoard[60].colour, "⚫")
	assert.Equal(t, nextTickBoard[61].colour, "⚪")
	assert.Equal(t, nextTickBoard[62].colour, "🟤")
	assert.Equal(t, nextTickBoard[63].colour, "🟡")
	assert.Equal(t, nextTickBoard[64].colour, "🟠")

}
