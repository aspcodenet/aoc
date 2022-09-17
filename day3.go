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

func day3Part2() {
	var allValues []string
	getMesurement("day3input.txt", func(line string) {
		allValues = append(allValues, line)
	})

	oxygen := getOxygen(allValues)
	co2 := getCo2(allValues)

	g, _ := strconv.ParseInt(oxygen, 2, 64)
	e, _ := strconv.ParseInt(co2, 2, 64)

	fmt.Println(g * e)

}

func getOxygen(allValues []string) string {
	return filter(allValues, func(letter byte, antalZeroes, antalOnes int) bool {
		if antalZeroes > antalOnes {
			if letter == '0' {
				return true
			}
		} else {
			if letter == '1' {
				return true
			}
		}
		return false
	})
}

func getCo2(allValues []string) string {
	return filter(allValues, func(letter byte, antalZeroes, antalOnes int) bool {
		if antalZeroes <= antalOnes {
			if letter == '0' {
				return true
			}
		} else {
			if letter == '1' {
				return true
			}
		}
		return false
	})

}

func filter(allValues []string, shouldKeepFunc func(letter byte, antalZeroes int, antalOnes int) bool) string {
	var left []string
	left = allValues

	for bit := 0; bit < len(allValues[0]); bit++ {
		antalZeros, antalOnes := getCount(bit, left)
		var toKeep []string
		for _, line := range left {
			if shouldKeepFunc(line[bit], antalZeros, antalOnes) {
				toKeep = append(toKeep, line)
			}
		}
		left = toKeep
		if len(left) == 1 {
			break
		}
	}
	return left[0]
}

func getCount(bit int, allValues []string) (antalZeros int, antalOnes int) {
	antalZeros = 0
	antalOnes = 0
	for _, line := range allValues {
		if line[bit] == '0' {
			antalZeros++
		} else {
			antalOnes++
		}
	}
	return antalZeros, antalOnes
}
