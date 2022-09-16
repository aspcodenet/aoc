package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var file = "day1input.txt"

func getMesurement(processLineFunc func(line string)) {
	file, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rd := bufio.NewScanner(file)
	for rd.Scan() {
		read_line := rd.Text()

		// other code what work with parsed line...
		processLineFunc(read_line)
	}

}

func main() {

	last := -1
	increases := 0
	getMesurement(func(line string) {
		num, _ := strconv.Atoi(line)
		if last > 0 && num > last {
			increases++
		}
		last = num

	})
	fmt.Println(increases)

}
