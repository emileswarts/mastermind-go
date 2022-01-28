package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPUCCreatehallengeRow(t *testing.T) {
	challengeRow := CPUCreateChallengeRow()
	assert.Equal(t, len(challengeRow), 5)
	assert.Equal(t, challengeRow[0].x, 0)
	assert.Equal(t, challengeRow[0].y, 11)
}
