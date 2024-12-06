package main

import (
	"AoC2024/day1"
	"AoC2024/day2"
	"AoC2024/day3"
	"fmt"
)

func main() {
	fmt.Println("What day do you want to run!")
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch input {
	case "1":
		day1.CalculateDistance()
		day1.CalculateOcurences()
	case "2":
		day2.CalculateSafeReports()
	case "3":
		day3.CalculateMultiplications()
	}
}
