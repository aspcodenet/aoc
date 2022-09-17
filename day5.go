package main

import (
	"fmt"
	"strconv"
	"strings"
)

type VentItem struct {
	XStart int
	YStart int
	XEnd   int
	YEnd   int
}

func day5() {
	var lines []*VentItem
	xmax := 0
	ymax := 0
	getMesurement("day5input.txt", func(line string) {
		parts := strings.Split(line, "->")
		xposStart, yposStart := splitPos(parts[0])
		xposEnd, yposEnd := splitPos(parts[1])
		if !(xposStart == xposEnd || yposStart == yposEnd) {
			return
		}
		item := &VentItem{
			XStart: xposStart,
			YStart: yposStart,
			XEnd:   xposEnd,
			YEnd:   yposEnd,
		}
		if xposEnd > xmax {
			xmax = xposEnd
		}
		if yposEnd > ymax {
			ymax = yposEnd
		}
		if xposStart > xmax {
			xmax = xposStart
		}
		if yposStart > ymax {
			ymax = yposStart
		}

		lines = append(lines, item)
	})
	points := make([][]int, ymax+1)
	for i := range points {
		points[i] = make([]int, xmax+1)
	}

	for _, line := range lines {
		plotFrom(points, line.XStart, line.YStart, line.XEnd, line.YEnd)
	}

	antal := 0
	for row := range points {
		for col := 0; col <= xmax; col++ {
			if points[row][col] > 1 {
				antal++
			}
		}
	}
	fmt.Printf("%d", antal)

}

func plotFrom(points [][]int, xStart, yStart, xEnd, yEnd int) {
	if xStart == xEnd {
		for i := yStart; i <= yEnd; i++ {
			points[i][xStart]++
		}
		for i := yStart; i >= yEnd; i-- {
			points[i][xStart]++
		}
	}
	if yStart == yEnd {
		for i := xStart; i <= xEnd; i++ {
			points[yStart][i]++
		}
		for i := xStart; i >= xEnd; i-- {
			points[yStart][i]++
		}
	}
}

func splitPos(csv string) (xpos, ypos int) {
	parts := strings.Split(csv, ",")
	xpos, _ = strconv.Atoi(strings.Trim(parts[0], " "))
	ypos, _ = strconv.Atoi(strings.Trim(parts[1], " "))
	return xpos, ypos
}
