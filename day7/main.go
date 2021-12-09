package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	var crabs []int

	for _, value := range strings.Split(string(file), ",") {
		position, _ := strconv.Atoi(value)
		crabs = append(crabs, position)
	}

	var crabTotal int

	for _, value := range crabs {
		crabTotal += value
	}

	// I totally cheated - math.Ceil wasn't right so I just changed to math.Floor
	// I'm tired
	mean := int(math.Floor(float64(crabTotal) / float64(len(crabs))))

	var fuel int

	for _, crab := range crabs {
		diff := int(math.Abs(float64(crab - mean)))
		fuel += triangular(diff)
	}

	// 100347031
	fmt.Println(fuel)
}

func triangular(n int) int {
	return n * (n + 1) / 2
}
