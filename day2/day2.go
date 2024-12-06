package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateSafeReports() {
	file, _ := os.Open("day2/day2.txt")
	scanner := bufio.NewScanner(file)

	defer file.Close()

	safe := 0
	safeDamp := 0
	for scanner.Scan() {
		valueList := strings.Split(scanner.Text(), " ")
		if EvalLine(valueList) {
			safe += 1
		}

		works := EvalLine(valueList)
		if !works {
			for i := 0; i < len(valueList); i++ {
				newList := rmvElementFromArray(valueList, i)
				if EvalLine(newList) {
					safeDamp += 1
					break
				}
			}
		} else {
			safeDamp += 1
		}
	}

	fmt.Println(safe)
	fmt.Println(safeDamp)
}

func EvalLine(valueList []string) bool {
	side := ""

	for i := 1; i < len(valueList); i++ {
		l, _ := strconv.Atoi(valueList[i-1])
		r, _ := strconv.Atoi(valueList[i])

		if l-r > 3 || l-r < -3 {
			return false
		} else if l == r {
			return false
		}

		if side == "" && l-r > 0 {
			side = "up"
		} else if side == "" && l-r < 0 {
			side = "down"
		} else if side == "up" && l-r < 0 {
			return false
		} else if side == "down" && l-r > 0 {
			return false
		}
	}

	return true
}

func rmvElementFromArray(oldArray []string, el int) []string {
	var newArray []string
	for i := 0; i < len(oldArray); i++ {
		if i != el {
			newArray = append(newArray, oldArray[i])
		}
	}

	return newArray
}
