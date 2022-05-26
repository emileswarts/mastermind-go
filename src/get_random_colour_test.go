package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomColours(t *testing.T) {
	excludeList := []string{
		"🟢",
		"🟡",
		"🔵",
		"🟠",
		"🟢",
		"🟤",
		"⚪",
	}

	assert.Equal(t, "⚫", getRandomColour(excludeList))
}
