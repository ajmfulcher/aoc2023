package main

import (
	"fmt"
	"strings"

	"andrewfulcher.dev/aoc2023/input"
)

func segmentToArray(segment string) []string {
	return strings.Fields(strings.TrimSpace(segment))
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func calculatePoints(matches int) int {
	total := 1
	for x := 1; x < matches; x++ {
		total = total * 2
	}
	return total
}

func main() {
	input := input.InputToArray("test.txt")
	rt := 0
	for _, card := range input {
		total := 0
		want := strings.Split(card, ":")[1]
		numbers := segmentToArray(strings.Split(want, "|")[0])
		matches := segmentToArray(strings.Split(want, "|")[1])
		var seen []string
		for _, match := range matches {
			if contains(numbers, match) && !contains(seen, match) {
				total += 1
				seen = append(seen, match)
			}
		}
		if total > 0 {
			rt += calculatePoints(total)
		}
	}

	fmt.Println(rt)
}
