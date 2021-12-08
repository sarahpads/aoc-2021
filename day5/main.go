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

	for scanner.Scan() {
		parts := strings.FieldsFunc(scanner.Text(), f)
		startX, _ := strconv.Atoi(parts[0])
		startY, _ := strconv.Atoi(parts[1])
		endX, _ := strconv.Atoi(parts[2])
		endY, _ := strconv.Atoi(parts[3])
		start := Coord{startX, startY}
		end := Coord{endX, endY}
		var coords []Coord

		if startX == endX {
			coords = getVerticalCoords(start, end)
		} else if startY == endY {
			coords = getHorizontalCoords(start, end)
		} else {
			coords = getDiagonalCoords(start, end)
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

	// 15463
	fmt.Println(intersections)
}

func getHorizontalCoords(a Coord, b Coord) []Coord {
	coords := []Coord{b}
	inc := getInc(a.x, b.x)

	coord := a
	for coord != b {
		coords = append(coords, coord)
		coord.x += inc
	}

	return coords
}

func getVerticalCoords(a Coord, b Coord) []Coord {
	coords := []Coord{b}
	inc := getInc(a.y, b.y)

	coord := a
	for coord != b {
		coords = append(coords, coord)
		coord.y += inc
	}

	return coords
}

func getDiagonalCoords(a Coord, b Coord) []Coord {
	coords := []Coord{b}
	xInc := getInc(a.x, b.x)
	yInc := getInc(a.y, b.y)

	coord := a
	for coord != b {
		coords = append(coords, coord)
		coord.x += xInc
		coord.y += yInc
	}

	return coords
}

func getInc(a int, b int) int {
	if a-b > 0 {
		return -1
	} else {
		return 1
	}
}

func f(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}
