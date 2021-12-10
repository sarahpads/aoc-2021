package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var caves [][]int64

	for scanner.Scan() {
		var cave []int64

		for _, value := range scanner.Text() {
			number, _ := strconv.ParseInt(string(value), 10, 16)
			cave = append(cave, number)
		}

		caves = append(caves, cave)
	}

	var lowPoints []int64

	for caveIndex, cave := range caves {
		// if the caveIndex is 0, we don't have upper locations
		var hasTop = caveIndex > 0
		// if this is the last cave, we don't have lower locations
		var hasBottom = caveIndex < len(caves)-1

		for locationIndex, location := range cave {
			// use 10 as our empty state, since all of our locations are 9 and lower
			lowest := float64(10)

			// if the locationIndex is 0, we don't have a left neighbor
			var hasLeft = locationIndex > 0
			// if it's the last location, we don't have a right neighbor
			var hasRight = locationIndex < len(cave)-1

			if hasTop {
				lowest = math.Min(lowest, float64(caves[caveIndex-1][locationIndex]))
			}

			if hasBottom {
				lowest = math.Min(lowest, float64(caves[caveIndex+1][locationIndex]))
			}

			if hasLeft {
				lowest = math.Min(lowest, float64(cave[locationIndex-1]))
			}

			if hasRight {
				lowest = math.Min(lowest, float64(cave[locationIndex+1]))
			}

			// we only want to consider this a low point if it's LOWER, not equal to, its neighbours
			if location < int64(lowest) {
				lowPoints = append(lowPoints, location)
			}
		}
	}

	var total int64

	for _, value := range lowPoints {
		total += value + 1
	}

	fmt.Println(total)
}
