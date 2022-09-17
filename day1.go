package main

import (
	"fmt"
	"strconv"
)

func day1() {
	last := -1
	increases := 0
	getMesurement("day1input.txt", func(line string) {
		num, _ := strconv.Atoi(line)
		if last > 0 && num > last {
			increases++
		}
		last = num

	})
	fmt.Println(increases)

}
