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

	var num1s [12]int
	var total int = 0

	// count the number of 1s, if it's over 50% of the collection length,
	// then we know it's more common
	for scanner.Scan() {
		total++

		for i, rune := range scanner.Text() {
			if string(rune) == "1" {
				num1s[i]++
			}
		}
	}

	var gamma int = 0
	var epsilon int = 0

	for i, value := range num1s {
		if value >= (total / 2) {
			// if 1 is the most common number, we need to set that bit to 1
			gamma += 1 << (11 - i)
		} else {
			// otherwise, that bit should be 1 for epsilon
			epsilon += 1 << (11 - i)
		}
	}

	// 1110 1101 0101
	// 3797
	// 0001 0010 1010
	// 298
	fmt.Println(epsilon * gamma)
}
