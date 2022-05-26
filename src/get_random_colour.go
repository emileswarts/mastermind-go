package main

import (
	"math/rand"
	"time"
)

func getRandomColour(excludeList []string) string {
	colourSet := [7]string{"🔵", "🟡", "🟠", "🟢", "🟤", "⚪", "⚫"}
	availableColours := difference(colourSet, excludeList)

	rand.Seed(time.Now().UnixNano())
	return availableColours[rand.Intn(len(availableColours))]
}
