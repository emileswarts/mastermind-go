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

	assert.Equal(t, expectedResult[0], "CPU row not set")
}

func TestWinningScore(t *testing.T) {
	var expectedResult []string

	occupiedCells := []cell{
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
	board := generateBoard(occupiedCells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, expectedResult[0], "black")
	assert.Equal(t, expectedResult[1], "black")
	assert.Equal(t, expectedResult[2], "black")
	assert.Equal(t, expectedResult[3], "black")
	assert.Equal(t, expectedResult[4], "black")
}

func TestRightColoursWrongPlacementScore(t *testing.T) {
	var expectedResult []string

	occupiedCells := []cell{
		{0, 0, "green"},
		{1, 0, "white"},
		{2, 0, "brown"},
		{3, 0, "blue"},
		{4, 0, "black"},
		{0, 12, "black"},
		{1, 12, "green"},
		{2, 12, "blue"},
		{3, 12, "white"},
		{4, 12, "brown"},
	}
	board := generateBoard(occupiedCells)
	expectedResult = scoreRow(board, 12)

	assert.Equal(t, "white", expectedResult[0])
	assert.Equal(t, "white", expectedResult[1])
	assert.Equal(t, "white", expectedResult[2])
	assert.Equal(t, "white", expectedResult[3])
	assert.Equal(t, "white", expectedResult[4])
}

func TestMixRightColourAndPositionScore(t *testing.T) {
	var score []string

	occupiedCells := []cell{
		{0, 0, "green"}, //black
		{1, 0, "white"}, //white
		{2, 0, "brown"}, // white
		{3, 0, "blue"},  // white
		{4, 0, "black"}, // nothing

		{0, 12, "green"},
		{1, 12, "yellow"},
		{2, 12, "blue"},
		{3, 12, "white"},
		{4, 12, "brown"},
	}
	board := generateBoard(occupiedCells)
	score = scoreRow(board, 12)

	assert.Equal(t, "black", score[0])
	assert.Equal(t, "white", score[1])
	assert.Equal(t, "white", score[2])
	assert.Equal(t, "white", score[3])
	assert.Equal(t, "", score[4])
}

func TestNoScore(t *testing.T) {
	var score []string

	occupiedCells := []cell{
		{0, 0, "green"}, // nothing
		{1, 0, "white"}, // nothing
		{2, 0, "brown"}, // nothing
		{3, 0, "blue"},  // nothing
		{4, 0, "black"}, // nothing

		{0, 12, "yellow"},
		{1, 12, "yellow"},
		{2, 12, "yellow"},
		{3, 12, "yellow"},
		{4, 12, "yellow"},
	}
	board := generateBoard(occupiedCells)
	score = scoreRow(board, 12)

	assert.Equal(t, "", score[0])
	assert.Equal(t, "", score[1])
	assert.Equal(t, "", score[2])
	assert.Equal(t, "", score[3])
	assert.Equal(t, "", score[4])
}
