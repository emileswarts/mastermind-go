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
	occupiedCells := make([]cell, 0)
	emptyBoard := generateBoard(occupiedCells)

	assert.Equal(t, 0, emptyBoard[0].x)
	assert.Equal(t, 0, emptyBoard[0].y)
	assert.Equal(t, "", emptyBoard[0].colour)

	assert.Equal(t, 1, emptyBoard[21].x)
	assert.Equal(t, 4, emptyBoard[21].y)
	assert.Equal(t, "", emptyBoard[21].colour)

	assert.Equal(t, 4, emptyBoard[64].x)
	assert.Equal(t, 12, emptyBoard[64].y)
	assert.Equal(t, "", emptyBoard[64].colour)
}

func TestGenerateBoardWithOneColour(t *testing.T) {
	occupiedCells := []cell{
		{0, 0, "blue"},
	}

	emptyBoard := generateBoard(occupiedCells)

	assert.Equal(t, 0, emptyBoard[0].x)
	assert.Equal(t, 0, emptyBoard[0].y)
	assert.Equal(t, "blue", emptyBoard[0].colour)

	assert.Equal(t, 1, emptyBoard[21].x)
	assert.Equal(t, 4, emptyBoard[21].y)
	assert.Equal(t, "", emptyBoard[21].colour)

	assert.Equal(t, 4, emptyBoard[64].x)
	assert.Equal(t, 12, emptyBoard[64].y)
	assert.Equal(t, "", emptyBoard[64].colour)
}

func TestGenerateBoardWithMultipleColours(t *testing.T) {
	occupiedCells := []cell{
		{0, 0, "green"},
		{1, 0, "white"},
		{2, 0, "brown"},
	}

	emptyBoard := generateBoard(occupiedCells)

	assert.Equal(t, "green", emptyBoard[0].colour)
	assert.Equal(t, "white", emptyBoard[1].colour)
	assert.Equal(t, "brown", emptyBoard[2].colour)

	assert.Equal(t, 4, emptyBoard[64].x)
	assert.Equal(t, 12, emptyBoard[64].y)
	assert.Equal(t, "", emptyBoard[64].colour)
}

func TestFindCellsInRow(t *testing.T) {
	occupiedCells := []cell{}
	emptyBoard := generateBoard(occupiedCells)
	expectedResult := []cell{
		{0, 0, ""},
		{1, 0, ""},
		{2, 0, ""},
		{3, 0, ""},
		{4, 0, ""},
	}

	assert.Equal(t, expectedResult, findCellsInRow(emptyBoard, 0))
}

func TestFindCPUCells(t *testing.T) {
	occupiedCells := []cell{}
	emptyBoard := generateBoard(occupiedCells)
	expectedResult := []cell{
		{0, 0, ""},
		{1, 0, ""},
		{2, 0, ""},
		{3, 0, ""},
		{4, 0, ""},
	}
	result := findCPUCells(emptyBoard)

	assert.Equal(t, expectedResult, result)
}
