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
	digitRegex, _ := regexp.Compile(`(\d+)`)
	colourRegex, _ := regexp.Compile(`([a-z]+)`)
	sum := 0
	for _, line := range input {
		// fmt.Println(gameIdx)
		setSegment := strings.Split(line, ":")[1]
		sets := strings.Split(setSegment, ";")
		results := make(map[string]int)

		for _, set := range sets {
			// fmt.Println(set)

			draws := strings.Split(set, ",")

			for _, draw := range draws {
				digit, _ := strconv.Atoi(digitRegex.FindStringSubmatch(draw)[1])
				colour := colourRegex.FindStringSubmatch(draw)[1]
				if results[colour] < digit {
					results[colour] = digit
				}
			}

		}
		// fmt.Printf("Red %d Green %d Blue %d\n", results["red"], results["green"], results["blue"])
		sum += results["red"] * results["green"] * results["blue"]
	}
	fmt.Println(sum)
}
