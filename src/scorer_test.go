package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreWhenCPUNotSet(t *testing.T) {
	var expectedResult []string

	occupiedCells := []cell{}
	board := generateBoard(occupiedCells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[4], "CPU row not set")
}

func TestWinningScore(t *testing.T) {
	var expectedResult []string

	occupiedCells := []cell{
		{0, 0, "🟢"},
		{1, 0, "⚪"},
		{2, 0, "🟤"},
		{3, 0, "🔵"},
		{4, 0, "⚫"},

		{0, 12, "🟢"},
		{1, 12, "⚪"},
		{2, 12, "🟤"},
		{3, 12, "🔵"},
		{4, 12, "⚫"},
	}
	board := generateBoard(occupiedCells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[0], "⬛")
	assert.Equal(t, expectedResult[1], "⬛")
	assert.Equal(t, expectedResult[2], "⬛")
	assert.Equal(t, expectedResult[3], "⬛")
	assert.Equal(t, expectedResult[4], "⬛")
}

func TestRightColoursWrongPlacementScore(t *testing.T) {
	var expectedResult []string

	occupiedCells := []cell{
		{0, 0, "🟢"},
		{1, 0, "⚪"},
		{2, 0, "🟤"},
		{3, 0, "🔵"},
		{4, 0, "⚫"},
		{0, 12, "⚫"},
		{1, 12, "🟢"},
		{2, 12, "🔵"},
		{3, 12, "⚪"},
		{4, 12, "🟤"},
	}
	board := generateBoard(occupiedCells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, "⬜", expectedResult[0])
	assert.Equal(t, "⬜", expectedResult[1])
	assert.Equal(t, "⬜", expectedResult[2])
	assert.Equal(t, "⬜", expectedResult[3])
	assert.Equal(t, "⬜", expectedResult[4])
}

func TestMixRightColourAndPositionScore(t *testing.T) {
	var score []string

	occupiedCells := []cell{
		{0, 0, "🟢"}, //black
		{1, 0, "⚪"}, //white
		{2, 0, "🟤"}, // white
		{3, 0, "🔵"}, // white
		{4, 0, "⚫"}, // nothing

		{0, 12, "🟢"},
		{1, 12, "🟡"},
		{2, 12, "🔵"},
		{3, 12, "⚪"},
		{4, 12, "🟤"},
	}
	board := generateBoard(occupiedCells)
	score = scoreRow(board, 12)

	assert.Equal(t, "", score[0])
	assert.Equal(t, "⬛", score[1])
	assert.Equal(t, "⬜", score[2])
	assert.Equal(t, "⬜", score[3])
	assert.Equal(t, "⬜", score[4])
}

func TestNoScore(t *testing.T) {
	var score []string

	occupiedCells := []cell{
		{0, 0, "🟢"}, // nothing
		{1, 0, "⚪"}, // nothing
		{2, 0, "🟤"}, // nothing
		{3, 0, "🔵"}, // nothing
		{4, 0, "⚫"}, // nothing

		{0, 12, "🟡"},
		{1, 12, "🟡"},
		{2, 12, "🟡"},
		{3, 12, "🟡"},
		{4, 12, "🟡"},
	}
	board := generateBoard(occupiedCells)
	score = scoreRow(board, 12)

	assert.Equal(t, "", score[0])
	assert.Equal(t, "", score[1])
	assert.Equal(t, "", score[2])
	assert.Equal(t, "", score[3])
	assert.Equal(t, "", score[4])
}

func TestCheckWinner(t *testing.T) {
	rowScores := []string{"⬛", "⬛", "⬛", "⬛", "⬛"}

	result := checkWinner(rowScores)
	assert.Equal(t, result, true)
}
