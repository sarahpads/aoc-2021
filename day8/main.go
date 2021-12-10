package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int64

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "|")
		// ensure that all of our pattern and output strings are sorted alphabetically
		// for easier comparison
		patterns := SortCollection(strings.Fields(split[0]))
		outputs := SortCollection(strings.Fields(split[1]))
		displays := getDisplays(patterns)

		var result string
		for _, output := range outputs {
			result += displays[output]
		}

		value, _ := strconv.ParseInt(result, 10, 64)
		count += value
	}

	// 514
	fmt.Println(count)
}

func SortCollection(s []string) []string {
	var sorted []string

	// sort the individual strings alphabetically
	for _, value := range s {
		sorted = append(sorted, SortString(value))
	}

	return sorted
}

type Displays map[int]string
type Patterns []string
type Signals map[string]string

func getDisplays(patterns []string) map[string]string {
	// sorth the entire collection by length
	sort.Slice(patterns, func(a, b int) bool {
		return len(patterns[a]) < len(patterns[b])
	})

	var signals = make(map[string]string)
	displays := map[int]string{
		1: patterns[0],
		4: patterns[2],
		7: patterns[1],
		8: patterns[9],
	}

	// we can solve for A by finding the Difference between 1 and 7
	signals["a"] = Diff(displays[1], displays[7])[0]
	// F occurs more than F
	cf := displays[1]
	signals["c"], signals["f"] = GetMostCommon(patterns, string(cf[0]), string(cf[1]))
	// D occurs more than B
	bd := Diff(displays[7], displays[4])
	signals["b"], signals["d"] = GetMostCommon(patterns, bd[0], bd[1])
	// G occurs more than E
	eg := Diff(signals["a"]+signals["b"]+signals["c"]+signals["d"]+signals["f"], displays[8])
	signals["e"], signals["g"] = GetMostCommon(patterns, eg[0], eg[1])

	// remove signals b and f from 8 to get 2
	displays[2] = strings.Replace(strings.Replace(displays[8], signals["b"], "", -1), signals["f"], "", -1)
	// remove signal d from 8 to get 0
	displays[0] = strings.Replace(displays[8], signals["d"], "", -1)
	// remove signal c from 8 to get 6
	displays[6] = strings.Replace(displays[8], signals["c"], "", -1)
	displays[3] = SortString(signals["a"] + signals["c"] + signals["d"] + signals["f"] + signals["g"])
	displays[5] = SortString(signals["a"] + signals["b"] + signals["d"] + signals["f"] + signals["g"])
	displays[9] = strings.Replace(displays[8], signals["e"], "", -1)

	return map[string]string{
		displays[0]: "0",
		displays[1]: "1",
		displays[2]: "2",
		displays[3]: "3",
		displays[4]: "4",
		displays[5]: "5",
		displays[6]: "6",
		displays[7]: "7",
		displays[8]: "8",
		displays[9]: "9",
	}
}

func GetMostCommon(patterns Patterns, a string, b string) (string, string) {
	aCount := Count(patterns, a)
	bCount := Count(patterns, b)

	if aCount < bCount {
		return a, b
	} else {
		return b, a
	}
}

func Count(patterns Patterns, character string) int {
	var count int

	for _, pattern := range patterns {
		if strings.Contains(pattern, character) {
			count++
		}
	}

	return count
}

// calculate the diff between two strings
func Diff(a string, b string) []string {
	var diff []string

	for _, character := range b {
		if !strings.ContainsRune(a, character) {
			diff = append(diff, string(character))
		}
	}

	return diff
}

// custom sort interface for sorting a single string alphabetically
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)

	sort.Sort(sortRunes(r))
	return string(r)
}
