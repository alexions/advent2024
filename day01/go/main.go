package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed puzzle.txt
var puzzle []byte

func main() {
	var left, right []int

	re := regexp.MustCompile(`[0-9]+`)
	lines := strings.Split(string(puzzle), "\n")

	for _, line := range lines {
		match := re.FindAllString(line, -1)
		// Assume we have a correct input
		lnum, _ := strconv.Atoi(match[0])
		rnum, _ := strconv.Atoi(match[1])

		left = append(left, lnum)
		right = append(right, rnum)
	}
	sort.Ints(left)
	sort.Ints(right)

	// Assume sizes are equal
	diff := 0
	for i := 0; i < len(lines); i++ {
		diff += absDiffInt(left[i], right[i])
	}

	fmt.Println(diff)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
