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

	median := crabs[len(crabs)/2]

	var fuel float64

	for _, crab := range crabs {
		fuel += math.Abs(float64(crab - median))
	}

	// 356922
	fmt.Println(fuel)
}
