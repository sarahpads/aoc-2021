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

	openingTags := "([{<"
	// lookup map where key is opening rune and value is closing rune
	closingTag := map[rune]rune{
		40:  41,
		91:  93,
		123: 125,
		60:  62,
	}

	// maybe rune?
	var invalidCharacters []rune

	for scanner.Scan() {
		var previousTags []rune

	line:
		for _, value := range scanner.Text() {
			// for each character, determine if the next character is an opening or closing
			if strings.ContainsRune(openingTags, value) {
				// if this is an opening tag, we'll process it later
				previousTags = append(previousTags, value)
				continue
			}

			// if this is a closing, does it match the correct opening?
			openingTag := previousTags[len(previousTags)-1]
			if closingTag[openingTag] == value {
				// if it does, remove this openinTag from the queue, as it has been handled
				previousTags = previousTags[:len(previousTags)-1]
			} else {
				// otherwise, exit this line and add the character to our invalid slice
				invalidCharacters = append(invalidCharacters, value)
				break line
			}
		}
	}

	var total int

	for _, rune := range invalidCharacters {
		switch rune {
		// round bracket
		case 41:
			total += 3

			// square bracket
		case 93:
			total += 57

			// brace
		case 125:
			total += 1197

			// arrow
		case 62:
			total += 25137
		}
	}

	// >))>]]])))))]>)}]>>)}}}>>))>>>}]))}})>})}]}>]>>
	// 364389
	fmt.Println(string(invalidCharacters))
	fmt.Println(total)
}
