package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Number int
	Set    bool
}

type Grid struct {
	Row1 []*BingoNumber
	Row2 []*BingoNumber
	Row3 []*BingoNumber
	Row4 []*BingoNumber
	Row5 []*BingoNumber
}

func (g *Grid) DrawOnConsole() {
	fmt.Println()
	draw(g.Row1)
	draw(g.Row2)
	draw(g.Row3)
	draw(g.Row4)
	draw(g.Row5)
}

func draw(bingoNumber []*BingoNumber) {
	s := ""
	for i := 0; i < 5; i++ {
		if bingoNumber[i].Set {
			s += "**"
		} else {

			if bingoNumber[i].Number < 10 {
				s += " " + strconv.Itoa(bingoNumber[i].Number)
			} else {
				s += strconv.Itoa(bingoNumber[i].Number)

			}
		}
		s += " "
	}
	fmt.Println(s)

}

func splitTwoAndTwo(line string) []string {
	var result []string
	result = append(result, strings.Trim(line[0:2], " "))
	result = append(result, strings.Trim(line[3:5], " "))
	result = append(result, strings.Trim(line[6:8], " "))
	result = append(result, strings.Trim(line[9:11], " "))
	result = append(result, strings.Trim(line[12:14], " "))
	return result
}

func (g *Grid) Setrow(rowNo int, line string) {
	for _, numstring := range splitTwoAndTwo(line) {
		n, _ := strconv.Atoi(numstring)
		bingoNumber := &BingoNumber{Number: n}
		switch rowNo {
		case 1:
			g.Row1 = append(g.Row1, bingoNumber)
		case 2:
			g.Row2 = append(g.Row2, bingoNumber)
		case 3:
			g.Row3 = append(g.Row3, bingoNumber)
		case 4:
			g.Row4 = append(g.Row4, bingoNumber)
		case 5:
			g.Row5 = append(g.Row5, bingoNumber)
		}
	}
}

func (g *Grid) SumAllUnmarkedNumbers() int {
	sum := 0
	for _, bingoNumber := range g.Row1 {
		if !bingoNumber.Set {
			sum += bingoNumber.Number
		}
	}
	for _, bingoNumber := range g.Row2 {
		if !bingoNumber.Set {
			sum += bingoNumber.Number
		}
	}
	for _, bingoNumber := range g.Row3 {
		if !bingoNumber.Set {
			sum += bingoNumber.Number
		}
	}
	for _, bingoNumber := range g.Row4 {
		if !bingoNumber.Set {
			sum += bingoNumber.Number
		}
	}
	for _, bingoNumber := range g.Row5 {
		if !bingoNumber.Set {
			sum += bingoNumber.Number
		}
	}

	return sum
}

func (g *Grid) HasBingo() bool {
	if checkRow(g.Row1) || checkRow(g.Row2) || checkRow(g.Row3) || checkRow(g.Row4) || checkRow(g.Row5) {
		return true
	}
	for col := 0; col < 5; col++ {
		if checkCol(col, g.Row1, g.Row2, g.Row3, g.Row4, g.Row5) {
			return true
		}
	}
	return false
}

func checkCol(col int, bingoNumber1 []*BingoNumber, bingoNumber2 []*BingoNumber, bingoNumber3 []*BingoNumber, bingoNumber4 []*BingoNumber, bingoNumber5 []*BingoNumber) bool {
	return bingoNumber1[col].Set &&
		bingoNumber2[col].Set &&
		bingoNumber3[col].Set &&
		bingoNumber4[col].Set &&
		bingoNumber5[col].Set
}

func checkRow(Row []*BingoNumber) bool {
	for i := 0; i < 5; i++ {
		if !Row[i].Set {
			return false
		}
	}
	return true
}

func (g *Grid) Setnumber(number int) {
	for _, bingoNumber := range g.Row1 {
		if bingoNumber.Number == number {
			bingoNumber.Set = true
			return
		}
	}
	for _, bingoNumber := range g.Row2 {
		if bingoNumber.Number == number {
			bingoNumber.Set = true
			return
		}
	}
	for _, bingoNumber := range g.Row3 {
		if bingoNumber.Number == number {
			bingoNumber.Set = true
			return
		}
	}
	for _, bingoNumber := range g.Row4 {
		if bingoNumber.Number == number {
			bingoNumber.Set = true
			return
		}
	}
	for _, bingoNumber := range g.Row5 {
		if bingoNumber.Number == number {
			bingoNumber.Set = true
			return
		}
	}
}

func day4() {
	var drawNumbers []int
	var grids []*Grid

	fileRowNo := 0
	boardno := 0
	currentGridRow := 0

	getMesurement("day4input.txt", func(line string) {
		fileRowNo++
		if fileRowNo == 1 {
			for _, numString := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(numString)
				drawNumbers = append(drawNumbers, num)
			}
		} else if len(strings.Trim(line, " ")) == 0 {
			boardno++
		} else {
			if boardno > len(grids) {
				grids = append(grids, &Grid{})
				currentGridRow = 1
			} else {
				currentGridRow++
			}
			grids[boardno-1].Setrow(currentGridRow, line)
		}
	})

	var winningGrid *Grid
	var winningNum int
	for _, drawNum := range drawNumbers {
		fmt.Printf("Drawing %d\n", drawNum)
		for _, grid := range grids {
			grid.Setnumber(drawNum)
		}
		for _, grid := range grids {
			grid.DrawOnConsole()
		}
		for _, grid := range grids {
			if grid.HasBingo() {
				winningGrid = grid
				winningNum = drawNum
				break
			}
		}
		if winningGrid != nil {
			break
		}

	}
	num := winningGrid.SumAllUnmarkedNumbers()
	num *= winningNum
	fmt.Printf("%d", num)
}

func day4Part2() {
	var drawNumbers []int
	var grids []*Grid

	fileRowNo := 0
	boardno := 0
	currentGridRow := 0

	getMesurement("day4input.txt", func(line string) {
		fileRowNo++
		if fileRowNo == 1 {
			for _, numString := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(numString)
				drawNumbers = append(drawNumbers, num)
			}
		} else if len(strings.Trim(line, " ")) == 0 {
			boardno++
		} else {
			if boardno > len(grids) {
				grids = append(grids, &Grid{})
				currentGridRow = 1
			} else {
				currentGridRow++
			}
			grids[boardno-1].Setrow(currentGridRow, line)
		}
	})

	var winningGrid *Grid
	var winningNum int
	var stillPlaying = grids
	for _, drawNum := range drawNumbers {
		fmt.Printf("Drawing %d\n", drawNum)
		for _, grid := range stillPlaying {
			grid.Setnumber(drawNum)
		}
		for _, grid := range stillPlaying {
			grid.DrawOnConsole()
		}

		antalSomHarBingo := 0
		for _, grid := range stillPlaying {
			if grid.HasBingo() {
				antalSomHarBingo++
			}
		}
		for _, grid := range stillPlaying {
			if grid.HasBingo() && antalSomHarBingo == len(stillPlaying) {
				winningGrid = grid
				winningNum = drawNum
				break
			}
		}
		if winningGrid != nil {
			break
		}
		stillPlaying = nil
		for _, grid := range grids {
			if !grid.HasBingo() {
				stillPlaying = append(stillPlaying, grid)
			}
		}

	}
	num := winningGrid.SumAllUnmarkedNumbers()
	winningGrid.DrawOnConsole()
	num *= winningNum
	fmt.Printf("%d", num)
}
