package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed puzzle.txt
var puzzle []byte

const (
	OpMul       = "mul"
	OpActivate  = "do("
	OpDeacivate = "don"

	StateActive = iota
	StateDisable
)

func main() {
	r := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don\'t\(\)`)
	matches := r.FindAllStringSubmatch(string(puzzle), -1)
	//fmt.Printf("%v\n", matches)

	res := 0
	state := StateActive

	for _, line := range matches {
		cmd := strings.ToLower(line[0][:3])
		switch cmd {
		case OpMul:
			if state == StateActive {
				a, _ := strconv.Atoi(line[1])
				b, _ := strconv.Atoi(line[2])
				res += a * b
			}
		case OpActivate:
			state = StateActive
		case OpDeacivate:
			state = StateDisable
		}

	}

	fmt.Println(res)
}

func printArray(data []string) {
	for k, v := range data {
		fmt.Printf("%d: %v\n", k, v)
	}
	fmt.Printf("\n")
}
