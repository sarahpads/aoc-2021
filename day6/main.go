package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	ogFishes := strings.Split(string(file), ",")

	var fishes [9]int

	for _, fish := range ogFishes {
		number, _ := strconv.Atoi(fish)
		fishes[number]++
	}

	fmt.Println(fishes)

	for i := 1; i <= 80; i++ {
		var newFishes [9]int

		// 0 index each create a new fish
		newFishes[8] = fishes[0]
		// and then their timers reset to 6
		newFishes[6] = fishes[0]

		// we already handled eight, now we need to decrement all other fishes
		for j := 0; j <= 7; j++ {
			newFishes[j] += fishes[j+1]
		}

		fishes = newFishes
	}

	total := 0

	for _, fish := range fishes {
		total += fish
	}

	fmt.Println(total)
}
