package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2() {
	horizontal := 0
	depth := 0

	getMesurement("day2input.txt", func(line string) {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return
		}
		moves, _ := strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			horizontal += moves
		} else if parts[0] == "down" {
			depth += moves
		} else if parts[0] == "up" {
			depth -= moves
		}
	})
	total := horizontal * depth
	fmt.Printf("%d", total)
}

func day2part2() {
	horizontal := 0
	depth := 0
	aim := 0

	getMesurement("day2input.txt", func(line string) {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return
		}
		moves, _ := strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			horizontal += moves
			depth += aim * moves
		} else if parts[0] == "down" {
			aim += moves
		} else if parts[0] == "up" {
			aim -= moves
		}
	})
	total := horizontal * depth
	fmt.Printf("%d", total)
}
