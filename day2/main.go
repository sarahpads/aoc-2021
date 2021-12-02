package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FORWARD string = "forward"
	DOWN    string = "down"
	UP      string = "up"
)

func main() {
	// horizontal and vertical coords
	coords := [...]int{0, 0}
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		command, distance := parseCommand(scanner.Text())

		switch command {
		case FORWARD:
			coords[0] = coords[0] + distance

		case DOWN:
			coords[1] = coords[1] + distance

		case UP:
			coords[1] = coords[1] - distance

		default:
			panic(fmt.Sprintf("Unknown command %v", command))
		}
	}

	print(coords[0] * coords[1])
}

func parseCommand(input string) (string, int) {
	words := strings.Fields(input)
	distance, _ := strconv.Atoi(words[1])

	return words[0], distance
}
