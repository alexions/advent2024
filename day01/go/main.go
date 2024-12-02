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

	// Prepare the base stuff
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

	// Calculate the first part of puzzle
	diff := distance(left, right)
	fmt.Println(diff)

	// Calculate the second one
	s := similarity(left, right)
	fmt.Println(s)
}

func distance(left, right []int) int {

	sort.Ints(left)
	sort.Ints(right)

	// Assume sizes are equal
	diff := 0
	for i := 0; i < len(left); i++ {
		diff += absDiffInt(left[i], right[i])
	}

	return diff
}

func similarity(left, right []int) int {
	// Prepare the map of presence
	times := make(map[int]int)
	for _, n := range right {
		times[n]++
	}

	score := 0
	for _, n := range left {
		score += n * times[n]
	}

	return score
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
