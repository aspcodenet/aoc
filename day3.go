package main

import (
	"fmt"
	"strconv"
)

func day3() {
	zeroCount := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	total := 0

	getMesurement("day3input.txt", func(line string) {
		total++
		for i := 0; i < 12; i++ {
			if line[i] == '0' {
				zeroCount[i]++
			}
		}
	})
	gamma := ""
	epsilon := ""
	for i := 0; i < 12; i++ {
		if zeroCount[i] > total-zeroCount[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(g * e)

}
