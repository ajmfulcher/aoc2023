package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func inputToArray(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func main() {
	input := inputToArray("input.txt")
	gameRegex, _ := regexp.Compile(`Game (\d+)`)
	digitRegex, _ := regexp.Compile(`(\d+)`)
	colourRegex, _ := regexp.Compile(`([a-z]+)`)
	sumIds := 0
	for _, line := range input {
		gameIdx := gameRegex.FindStringSubmatch(line)[1]
		// fmt.Println(gameIdx)
		setSegment := strings.Split(line, ":")[1]
		sets := strings.Split(setSegment, ";")
		possible := true
		for _, set := range sets {
			// fmt.Println(set)
			results := make(map[string]int)
			draws := strings.Split(set, ",")

			for _, draw := range draws {
				digit, _ := strconv.Atoi(digitRegex.FindStringSubmatch(draw)[1])
				colour := colourRegex.FindStringSubmatch(draw)[1]
				results[colour] = results[colour] + digit
			}
			if results["red"] > 12 || results["green"] > 13 || results["blue"] > 14 {
				possible = false
				break
			}
		}
		if possible {
			numId, _ := strconv.Atoi(gameIdx)
			sumIds += numId
		}
	}
	fmt.Println(sumIds)
}
