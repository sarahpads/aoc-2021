package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// use map instead of array/slice, since go does not have a 'contains' method
	// this provides an easier way to check for existance within a set
	lengths := map[int]int{2: 2, 3: 3, 4: 4, 7: 7}

	var count int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "|")
		output := split[1]

		for _, value := range strings.Fields(output) {
			if _, ok := lengths[len(value)]; ok {
				count++
			}
		}
	}

	// 514
	fmt.Println(count)
}
