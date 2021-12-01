package main

import (
	"bufio"
	"fmt"
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
	var lines []int

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}

	file.Close()

	var increased int
	var decreased int

	for index, line := range lines {
		if index == 0 {
			continue
		} else if line > lines[index-1] {
			increased++
		} else {
			decreased++
		}
	}

	fmt.Println(increased)
	fmt.Println(decreased)
}
