package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var scores []int
	openingTags := "([{<"
	// lookup map where key is opening rune and value is closing rune
	closingTag := map[rune]rune{
		40:  41,
		91:  93,
		123: 125,
		60:  62,
	}

	for scanner.Scan() {
		var score int
		var previousTags []rune
		line := scanner.Text()

	line:
		for _, value := range line {
			// for each character, determine if the next character is an opening or closing
			if strings.ContainsRune(openingTags, value) {
				// if this is an opening tag, we'll process it later
				previousTags = append(previousTags, value)
				continue
			}

			// if this is a closing, does it match the correct opening?
			openingTag := previousTags[len(previousTags)-1]
			if closingTag[openingTag] == value {
				// if it does, remove this openingTag from the queue, as it has been handled
				previousTags = previousTags[:len(previousTags)-1]
			} else {
				// otherwise, the line is corrupt so exit
				previousTags = nil
				break line
			}
		}

		// did we process the whole line without closing all of our tags?
		if len(previousTags) > 0 {
			score = CalculateScore(previousTags)
		}

		// don't count empty scores for corrupt lines
		if score > 0 {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	middle := len(scores) / 2

	// 2870201088
	fmt.Println(scores[middle])
}

func CalculateScore(characters []rune) int {
	var score int

	// we need to count up the scores in closing tag order, so reverse our opening tags
	for i := len(characters) - 1; i >= 0; i-- {
		rune := characters[i]
		score *= 5

		// for each unclosed opening tag, count up the score
		switch rune {
		// round bracket
		case 40:
			score += 1

			// square bracket
		case 91:
			score += 2

			// brace
		case 123:
			score += 3

			// arrow
		case 60:
			score += 4
		}
	}

	return score
}
