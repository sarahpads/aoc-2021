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

	// part1(scanner)
	part2(scanner)
}

func part1(scanner *bufio.Scanner) {
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

func part2(scanner *bufio.Scanner) {
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	oxygen := lines
	carbon := lines

	for depth := 0; depth < 12; depth++ {
		oxygen = filter(&oxygen, depth, 1)
		carbon = filter(&carbon, depth, -1)
	}

	// 1111 1111 1001
	oxygenValue, _ := strconv.ParseInt(oxygen[0], 2, 64)
	// 0111 1000 0011
	carbonValue, _ := strconv.ParseInt(carbon[0], 2, 64)

	// 7863147
	fmt.Println(oxygenValue * carbonValue)
}

// accept pointer of lines so that the array isn't copied; no need to waste memory
// modifier indicates if we're looking for the most common (1) or least common (-1)
func filter(lines *[]string, depth int, modifier int) []string {
	if len(*lines) == 1 {
		return *lines
	}

	var ones []string
	var zeros []string

	for _, value := range *lines {
		if string(value[depth]) == "1" {
			ones = append(ones, value)
		} else {
			zeros = append(zeros, value)
		}
	}

	diff := len(ones) - len(zeros)

	if diff == 0 {
		if modifier > 0 {
			return ones
		} else {
			return zeros
		}
	}

	if diff*modifier > 0 {
		return ones
	} else {
		return zeros
	}
}
