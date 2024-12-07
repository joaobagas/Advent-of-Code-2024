package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CalculateWords() {
	file, _ := os.Open("day4/day4.txt")
	scanner := bufio.NewScanner(file)

	defer file.Close()

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	countXMAS := 0
	countMAS := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "X" {
				countXMAS += CheckHorizontalXMAS(lines, i, j)
				countXMAS += CheckVerticalXMAS(lines, i, j)
				countXMAS += CheckDiagonalXMAS(lines, i, j)
			}
			if lines[i][j] == "A" {
				countMAS += CalculateMas(lines, i, j)
			}
		}
	}

	fmt.Println(countXMAS)
	fmt.Println(countMAS)
}

func CheckHorizontalXMAS(lines [][]string, x int, y int) int {
	count := 0
	maxSize := len(lines) - 1

	if x-3 >= 0 {
		if lines[x-1][y] == "M" && lines[x-2][y] == "A" && lines[x-3][y] == "S" {
			count += 1
		}
	}

	if x+3 <= maxSize {
		if lines[x+1][y] == "M" && lines[x+2][y] == "A" && lines[x+3][y] == "S" {
			count += 1
		}
	}

	return count
}

func CheckVerticalXMAS(lines [][]string, x int, y int) int {
	count := 0
	maxSize := len(lines[x]) - 1

	if y-3 >= 0 {
		if lines[x][y-1] == "M" && lines[x][y-2] == "A" && lines[x][y-3] == "S" {
			count += 1
		}
	}

	if y+3 <= maxSize {
		if lines[x][y+1] == "M" && lines[x][y+2] == "A" && lines[x][y+3] == "S" {
			count += 1
		}
	}

	return count
}

func CheckDiagonalXMAS(lines [][]string, x int, y int) int {
	count := 0
	maxSizeX := len(lines) - 1
	maxSizeY := len(lines[x]) - 1

	if x-3 >= 0 && y-3 >= 0 {
		if lines[x-1][y-1] == "M" && lines[x-2][y-2] == "A" && lines[x-3][y-3] == "S" {
			count += 1
		}
	}

	if x-3 >= 0 && y+3 <= maxSizeY {
		if lines[x-1][y+1] == "M" && lines[x-2][y+2] == "A" && lines[x-3][y+3] == "S" {
			count += 1
		}
	}

	if x+3 <= maxSizeX && y-3 >= 0 {
		if lines[x+1][y-1] == "M" && lines[x+2][y-2] == "A" && lines[x+3][y-3] == "S" {
			count += 1
		}
	}

	if x+3 <= maxSizeX && y+3 <= maxSizeY {
		if lines[x+1][y+1] == "M" && lines[x+2][y+2] == "A" && lines[x+3][y+3] == "S" {
			count += 1
		}
	}

	return count
}

func CalculateMas(lines [][]string, x int, y int) int {
	count := 0
	maxSizeX := len(lines) - 1
	maxSizeY := len(lines[x]) - 1

	if x-1 >= 0 && x+1 <= maxSizeX && y-1 >= 0 && y+1 <= maxSizeY {
		firstLine := (lines[x-1][y-1] == "M" && lines[x+1][y+1] == "S") || (lines[x-1][y-1] == "S" && lines[x+1][y+1] == "M")
		secondLine := (lines[x-1][y+1] == "M" && lines[x+1][y-1] == "S") || (lines[x-1][y+1] == "S" && lines[x+1][y-1] == "M")

		if firstLine && secondLine {
			count += 1
		}
	}

	return count
}
