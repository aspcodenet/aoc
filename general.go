package main

import (
	"bufio"
	"log"
	"os"
)

func getMesurement(filePath string, processLineFunc func(line string)) {
	file, err := os.Open(filePath)
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
