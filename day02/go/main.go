package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed puzzle.txt
var puzzle []byte

type reportType int

const (
	maxDiff               = 3
	toleranceErrThreshold = 1

	reportUndef reportType = iota
	reportDec
	reportInc
)

func main() {
	var safeCount int

	reports := strings.Split(string(puzzle), "\n")

	for _, reportRaw := range reports {
		data := strings.Split(reportRaw, " ")

		if isValidReport(data, 0) {
			safeCount++
		}
	}

	fmt.Println(len(reports), safeCount)
}

func isValidReport(data []string, toleranceAcc int) bool {
	if toleranceAcc > toleranceErrThreshold {
		return false
	}

	current := -1
	rType := reportUndef

	for i := 0; i < len(data); i++ {
		element, _ := strconv.Atoi(data[i])

		// Initial Setup
		if current < 0 {
			current = element
			continue
		}

		if rType == reportUndef {
			if current < element {
				rType = reportInc
			} else {
				rType = reportDec
			}
		}

		// Validation
		diff := element - current
		if diff > maxDiff || diff < -maxDiff || diff == 0 || // check the max allowed difference
			rType == reportInc && diff < 0 || // broken type
			rType == reportDec && diff > 0 {

			// Recursion validation
			// There are 3 cases: 12 10 11
			// 1) we can try to remove the first item and get the valid INCREASING report, 10 11 (12 15 etc)
			// 2) we can try to remove the last of 3 and get the valid DECREASING 12 10 (8 7 etc)
			// 3) remove the middle item and still get the DECREASING report, but probably valid in case of 14 first:  14 11 8

			if i >= 2 && isValidReport(slices.Concat(data[0:i-2], data[i-1:]), toleranceAcc+1) {
				return true
			}

			if isValidReport(slices.Concat(data[0:i-1], data[i:]), toleranceAcc+1) {
				return true
			}

			if isValidReport(slices.Concat(data[:i], data[i+1:]), toleranceAcc+1) {
				return true
			}

			return false
		}

		current = element
	}

	return true
}
