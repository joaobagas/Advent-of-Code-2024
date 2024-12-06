package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateMultiplications() {
	file, _ := os.Open("day3/day3.txt")
	scanner := bufio.NewScanner(file)

	defer file.Close()

	firstVal := 0
	secondVal := 0
	tempVal := 0
	isDo := true
	for scanner.Scan() {
		line := scanner.Text()
		firstVal += DecypherOnlyMul(strings.Split(line, ""))

		tempVal, isDo = DecypherMulAndDo(strings.Split(line, ""), isDo)
		secondVal += tempVal
	}

	fmt.Println(firstVal)
	fmt.Println(secondVal)
}

func DecypherOnlyMul(txt []string) int {
	ret := 0
	for i := 3; i < len(txt); i++ {
		if txt[i-3] == "m" && txt[i-2] == "u" && txt[i-1] == "l" && txt[i] == "(" {
			commaIndex := 0
			rightIndex := 0
			firstNumber := 0
			secondNumber := 0

			for j := i + 2; j < i+5; j++ {
				if txt[j] == "," {
					commaIndex = j
					break
				}
			}

			for j := commaIndex + 2; j < commaIndex+5; j++ {
				if txt[j] == ")" {
					rightIndex = j
					break
				}
			}

			if commaIndex != 0 && rightIndex != 0 {
				firstNumber, _ = strconv.Atoi(strings.Join(txt[i+1:commaIndex], ""))
				secondNumber, _ = strconv.Atoi(strings.Join(txt[commaIndex+1:rightIndex], ""))
				ret += firstNumber * secondNumber
			}
		}
	}
	return ret
}

func DecypherMulAndDo(txt []string, isDo bool) (int, bool) {
	ret := 0
	for i := 3; i < len(txt); i++ {
		// Check mul()

		if txt[i-3] == "m" && txt[i-2] == "u" && txt[i-1] == "l" && txt[i] == "(" {
			commaIndex := 0
			rightIndex := 0
			firstNumber := 0
			secondNumber := 0

			for j := i + 2; j < i+5; j++ {
				if txt[j] == "," {
					commaIndex = j
					break
				}
			}

			for j := commaIndex + 2; j < commaIndex+5; j++ {
				if txt[j] == ")" {
					rightIndex = j
					break
				}
			}

			if commaIndex != 0 && rightIndex != 0 && isDo {
				firstNumber, _ = strconv.Atoi(strings.Join(txt[i+1:commaIndex], ""))
				secondNumber, _ = strconv.Atoi(strings.Join(txt[commaIndex+1:rightIndex], ""))
				ret += firstNumber * secondNumber
			}
		}

		// Check do and don't
		if txt[i-3] == "d" && txt[i-2] == "o" && txt[i-1] == "n" && txt[i] == "'" {
			if txt[i+1] == "t" && txt[i+2] == "(" && txt[i+3] == ")" {
				isDo = false
			}
		} else if txt[i-3] == "d" && txt[i-2] == "o" && txt[i-1] == "(" && txt[i] == ")" {
			isDo = true
		}

	}
	return ret, isDo
}
