package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	var numbers []int
	var boards [][25]int
	var boardProgress [][10]int

	scanner := bufio.NewScanner(file)
	scanner.Split(ScanTwoConsecutiveNewlines)
	scanner.Scan()
	comma := regexp.MustCompile(",")

	for _, value := range comma.Split(scanner.Text(), -1) {
		parsedValue, _ := strconv.Atoi(value)
		numbers = append(numbers, parsedValue)
	}

	for scanner.Scan() {
		var board [25]int

		for i, value := range strings.Fields(scanner.Text()) {
			parsedValue, _ := strconv.Atoi(value)
			board[i] = parsedValue
		}

		boards = append(boards, board)
		var emptyProgress [10]int
		boardProgress = append(boardProgress, emptyProgress)
	}

	var score int

bingo:
	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {
			var board = boards[i]
			var numberIndex = IndexOf(board, number)

			if numberIndex == -1 {
				continue
			}

			// find the coord
			var row = numberIndex / 5
			var column = numberIndex % 5
			boards[i][numberIndex] = -1
			boardProgress[i][row]++
			boardProgress[i][column+5]++

			hasBingo := checkForBingo(boardProgress[i])

			if hasBingo {
				score = calculateScore(boards[i], number)
				break bingo
			}
		}
	}

	fmt.Println(score)
}

func calculateScore(board [25]int, multiplier int) int {
	var total = 0

	for _, value := range board {
		if value == -1 {
			continue
		}

		total += value
	}

	return total * multiplier
}

func checkForBingo(board [10]int) bool {
	var hasBingo = false

	for _, value := range board {
		if value >= 5 {
			hasBingo = true
			break
		}
	}

	return hasBingo
}

func IndexOf(haystack [25]int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}

	return -1
}

var (
	patEols  = regexp.MustCompile(`[\r\n]+`)
	pat2Eols = regexp.MustCompile(`[\r\n]{2}`)
)

func ScanTwoConsecutiveNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if loc := pat2Eols.FindIndex(data); loc != nil && loc[0] >= 0 {
		// Replace newlines within string with a space
		s := patEols.ReplaceAll(data[0:loc[0]+1], []byte(" "))
		// Trim spaces and newlines from string
		s = bytes.Trim(s, "\n ")
		return loc[1], s, nil
	}

	if atEOF {
		// Replace newlines within string with a space
		s := patEols.ReplaceAll(data, []byte(" "))
		// Trim spaces and newlines from string
		s = bytes.Trim(s, "\r\n ")
		return len(data), s, nil
	}

	// Request more data.
	return 0, nil, nil
}
