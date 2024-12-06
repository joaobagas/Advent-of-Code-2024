package day1

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func Load() ([]string, []string) {
	file, err := os.Open("day1/day1.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	var left []string
	var right []string
	for _, record := range records {
		left = append(left, record[0])
		right = append(right, record[1])
	}

	return left, right
}

func CalculateDistance() {
	left, right := Load()

	slices.Sort(left)
	slices.Sort(right)

	d := 0
	for i := 0; i < len(left); i++ {
		l, _ := strconv.Atoi(left[i])
		r, _ := strconv.Atoi(right[i])
		d += int(math.Abs(float64(l - r)))
	}

	fmt.Println(d)
}

func CalculateOcurences() {
	left, right := Load()

	var hm map[string]int
	hm = make(map[string]int)

	d := 0

	for i := 0; i < len(left); i++ {
		l, _ := strconv.Atoi(left[i])
		if hm[left[i]] == 0 {
			c := 0
			for j := 0; j < len(right); j++ {
				r, _ := strconv.Atoi(right[j])
				if l == r {
					c += 1
				}
			}
			hm[left[i]] = c
		}
		d += l * hm[left[i]]
	}

	fmt.Println(d)
}
