package main

import (
	"fmt"
	"time"
)

// Maps for localization
var daysInPortuguese = map[string]string{
	"Monday":    "Seg",
	"Tuesday":   "Ter",
	"Wednesday": "Quar",
	"Thursday":  "Quin",
	"Friday":    "Sex",
	"Saturday":  "Sab",
	"Sunday":    "Dom",
}

func main() {
	currentTime := time.Now()

	// Get English day and month
	day := currentTime.Format("Monday")

	// Map to Portuguese
	dayPT := daysInPortuguese[day]

	// Print in Portuguese
	fmt.Printf("%s, %02d/%02d/%02d, %02d:%02d\n", dayPT, currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
}
