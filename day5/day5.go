package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CalculateCorrectOrder() {
	file, _ := os.Open("day5/day5.txt")
	scanner := bufio.NewScanner(file)

	defer file.Close()

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	isLoading := true
	var hash map[string][]string
	hash = make(map[string][]string)

	countRight := 0
	countWrong := 0

	for i := 0; i < len(lines); i++ {
		if isLoading == true && len(lines[i]) > 1 {
			input := strings.Join(lines[i][0:2], "")
			output := strings.Join(lines[i][3:5], "")
			hash[input] = append(hash[input], output)
		} else if isLoading == true {
			isLoading = false
		} else if isLoading == false && len(lines[i]) > 1 {
			if !CheckIfWrong(lines, i, hash) {
				countRight += CalculateMiddle(lines, i)
			} else {
				//FixOrder()
				countWrong += CalculateMiddle(lines, i)
			}
		}
	}

	fmt.Println(countRight)
	fmt.Println(countWrong)
}

func CheckIfWrong(lines [][]string, index int, hash map[string][]string) bool {
	for i := 5; i < len(lines[index]); i += 3 {
		for j := 2; j < i; j += 3 {
			secondNumber := strings.Join(lines[index][i-2:i], "")
			firstNumber := strings.Join(lines[index][j-2:j], "")

			arr := hash[secondNumber]
			fmt.Println(arr)
			if slices.Contains(arr, firstNumber) {
				return true
			}
		}
	}

	return false
}

func CalculateMiddle(lines [][]string, index int) int {
	size := len(lines[index])
	middleLast := size / 2
	ret, _ := strconv.Atoi(strings.Join(lines[index][middleLast-1:middleLast+1], ""))
	return ret
}

func FixOrder(lines [][]string, index int, hash map[string][]string) []string {
	init := lines[index]

	for i := 5; i < len(lines[index]); i += 3 {
		for j := 2; j < i; j += 3 {
			secondNumber := strings.Join(lines[index][i-2:i], "")
			firstNumber := strings.Join(lines[index][j-2:j], "")

			arr := hash[secondNumber]
			fmt.Println(arr)
			if slices.Contains(arr, firstNumber) {
				init = DeleteFromArray(init, 1, 2)
				AddBackToArray(init, 1)
			}
		}
	}

	return init
}

func DeleteFromArray(oldArray []string, el1 int, el2 int) []string {
	var newArray []string
	for i := 0; i < len(oldArray); i++ {
		if i != el1 && i != el2 {
			newArray = append(newArray, oldArray[i])
		}
	}

	return newArray
}

func AddBackToArray(oldArray []string, el int) {
}
