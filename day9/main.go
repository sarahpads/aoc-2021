package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var basins []int

	for caveIndex := 0; caveIndex < len(caves); caveIndex++ {
		for locationIndex := 0; locationIndex < len(caves[caveIndex]); locationIndex++ {
			basin := GetBasinSize(&caves, caveIndex, locationIndex)

			if basin > 0 {
				basins = append(basins, basin)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	var total int = basins[0]

	for _, value := range basins[1:3] {
		total *= value
	}

	fmt.Println(total)
}

// caves needs to be a pointer, so that we can mutate the collection
func GetBasinSize(caves *[][]int64, caveIndex int, locationIndex int) int {
	var basinSize int
	cave := (*caves)[caveIndex]
	location := cave[locationIndex]

	// if this number isn't 9, recursively find all of its non-9 neighbours
	if location >= 9 {
		return 0
	}

	// as we identify pieces that belong to a basin, change their value to 10 so that we don't count them twice
	basinSize++
	(*caves)[caveIndex][locationIndex] = 10

	// if the caveIndex is 0, we don't have upper locations
	var hasTop = caveIndex > 0
	// if this is the last cave, we don't have lower locations
	var hasBottom = caveIndex < len(*caves)-1
	// if the locationIndex is 0, we don't have a left neighbor
	var hasLeft = locationIndex > 0
	// if it's the last location, we don't have a right neighbor
	var hasRight = locationIndex < len(cave)-1

	if hasTop && (*caves)[caveIndex-1][locationIndex] < 9 {
		basinSize += GetBasinSize(caves, caveIndex-1, locationIndex)
	}

	if hasBottom && (*caves)[caveIndex+1][locationIndex] < 9 {
		basinSize += GetBasinSize(caves, caveIndex+1, locationIndex)
	}

	if hasLeft && cave[locationIndex-1] < 9 {
		basinSize += GetBasinSize(caves, caveIndex, locationIndex-1)
	}

	if hasRight && cave[locationIndex+1] < 9 {
		basinSize += GetBasinSize(caves, caveIndex, locationIndex+1)
	}

	return basinSize
}
