package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Coord struct {
	x, y int
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	points := make(map[Coord]int)

	// [int, int] = int

	for scanner.Scan() {
		// print(scanner.Text())
		parts := strings.FieldsFunc(scanner.Text(), f)
		startX, _ := strconv.Atoi(parts[0])
		startY, _ := strconv.Atoi(parts[1])
		endX, _ := strconv.Atoi(parts[2])
		endY, _ := strconv.Atoi(parts[3])
		start := Coord{startX, startY}
		end := Coord{endX, endY}
		var coords []Coord

		if startX == endX {
			coords = getLineCoords(start.y, end.y, start.x, "vertical")
		} else if startY == endY {
			coords = getLineCoords(start.x, end.x, start.y, "horizontal")
		} else {
			continue
		}

		for _, coord := range coords {
			points[coord]++
		}
	}

	var intersections int = 0

	for _, occurences := range points {
		if occurences > 1 {
			intersections++
		}
	}

	// 5698
	fmt.Println(intersections)
}

func getLineCoords(a int, b int, c int, axis string) []Coord {
	var coords []Coord
	var start int
	var end int

	if a-b > 0 {
		start = b
		end = a
	} else {
		start = a
		end = b
	}

	for i := start; i <= end; i++ {
		if axis == "horizontal" {
			coords = append(coords, Coord{i, c})
		} else {
			coords = append(coords, Coord{c, i})
		}
	}

	return coords
}

func f(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}
