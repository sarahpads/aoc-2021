package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var flashes int
	var octopi [][]int

	for scanner.Scan() {
		var row []int

		for _, value := range scanner.Text() {
			// For those who is wondering how int(r-'0') even work, short explanation: "subtracting the value of rune '0' from any rune '0' through '9' will give you an integer 0 through 9". Resulting type after subtraction r-'0' will be int32 (base type of runes), that is why if you need int type - you will also need int() conversion
			// https://stackoverflow.com/a/21322694
			number := int(value - '0')
			row = append(row, number)
		}

		octopi = append(octopi, row)
	}

	for i := 0; i < 100; i++ {
		flashes += PerformStep(&octopi)
	}

	fmt.Println(flashes)
}

func PerformStep(octopi *[][]int) int {
	// my recursive fun caused a stack overflow, so will need to iterate over the collection
	// multiple times until there are no 10s
	var flashCounter int
	// the flashes that occured that we need to process
	var flashStack [][2]int

	for rowIndex := 0; rowIndex < len(*octopi); rowIndex++ {
		for octopusIndex := 0; octopusIndex < len((*octopi)[rowIndex]); octopusIndex++ {
			if (*octopi)[rowIndex][octopusIndex] == 9 {
				// if our octopus is going to flash, add it to the stack and reset to 0
				flashStack = append(flashStack, [2]int{rowIndex, octopusIndex})
				(*octopi)[rowIndex][octopusIndex] = 0
			} else {
				// otherwise, just increment
				(*octopi)[rowIndex][octopusIndex]++
			}
		}
	}

	// handle any flashes
	for len(flashStack) > 0 {
		flashCounter++
		// find out if this flash causes any other flashes
		newFlashes := ProcessFlash(octopi, flashStack[0][0], flashStack[0][1])
		// remove this octopus from our stack
		flashStack = flashStack[1:]
		// append any new flashes to our stack
		flashStack = append(flashStack, newFlashes...)
	}

	return flashCounter
}

func ProcessFlash(octopi *[][]int, rowIndex int, octopusIndex int) [][2]int {
	var newFlashes [][2]int
	row := (*octopi)[rowIndex]
	var neighbours [][2]int

	var hasTop = rowIndex > 0
	var hasBottom = rowIndex < len(*octopi)-1
	var hasLeft = octopusIndex > 0
	var hasRight = octopusIndex < len(row)-1

	if hasTop {
		neighbours = append(neighbours, [2]int{rowIndex - 1, octopusIndex})
	}

	if hasBottom {
		neighbours = append(neighbours, [2]int{rowIndex + 1, octopusIndex})
	}

	if hasLeft {
		neighbours = append(neighbours, [2]int{rowIndex, octopusIndex - 1})
	}

	if hasRight {
		neighbours = append(neighbours, [2]int{rowIndex, octopusIndex + 1})
	}

	if hasTop && hasLeft {
		neighbours = append(neighbours, [2]int{rowIndex - 1, octopusIndex - 1})
	}

	if hasTop && hasRight {
		neighbours = append(neighbours, [2]int{rowIndex - 1, octopusIndex + 1})
	}

	if hasBottom && hasLeft {
		neighbours = append(neighbours, [2]int{rowIndex + 1, octopusIndex - 1})
	}

	if hasBottom && hasRight {
		neighbours = append(neighbours, [2]int{rowIndex + 1, octopusIndex + 1})
	}

	for _, neighbour := range neighbours {
		var energy int = (*octopi)[neighbour[0]][neighbour[1]]

		// if this neighbour already has an energy level of nine, it will flash
		if energy == 9 {
			newFlashes = append(newFlashes, neighbour)
			(*octopi)[neighbour[0]][neighbour[1]] = 0
		} else if energy > 0 {
			// we only want to increment if this hasn't already flashed
			(*octopi)[neighbour[0]][neighbour[1]]++
		}
	}

	return newFlashes
}
