package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var digitRegex, _ = regexp.Compile(`\d`)

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

func isSymbol(char byte) bool {
	asString := string(char)
	return !digitRegex.MatchString(asString) && asString != "."
}

func isDigit(char byte) bool {
	return digitRegex.MatchString(string(char))
}

func max(a, b int) int {
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func min(a, b int) int {
	var min int
	if a > b {
		min = b
	} else {
		min = a
	}
	return min
}

type xy struct {
	x int
	y int
}

func main() {
	input := inputToArray("input.txt")

	var hits []xy

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			// Identify symbols
			if isSymbol(input[y][x]) {
				left := max(x-1, 0)
				top := max(y-1, 0)
				right := min(x+2, len(input[x]))
				bottom := min(y+2, len(input))

				for y2 := top; y2 < bottom; y2++ {
					for x2 := left; x2 < right; x2++ {
						// Identify digits adjacent to symbols
						if isDigit(input[y2][x2]) {
							// fmt.Printf("%d,%d\n", x2, y2)
							hits = append(hits, xy{x: x2, y: y2})
						}
					}
				}
			}
		}
	}

	var total = 0

	visited := make([][]bool, len(input))
	for i := range visited {
		visited[i] = make([]bool, len(input[i]))
	}

	for _, hit := range hits {
		// fmt.Printf("%d,%d\n", hit.x, hit.y)
		if visited[hit.y][hit.x] {
			continue
		}
		pos := hit.x
		edge := false
		for pos > 0 && !edge {
			if !isDigit(input[hit.y][pos-1]) {
				edge = true
			} else {
				pos -= 1
			}
		}
		edge = false
		pn := ""
		for pos < len(input[hit.y]) && !edge {
			if !isDigit(input[hit.y][pos]) {
				edge = true
			} else {
				pn = pn + (string(input[hit.y][pos]))
				visited[hit.y][pos] = true
				pos += 1
			}
		}
		// fmt.Println(pn)
		sum, _ := strconv.Atoi(pn)
		total += sum
	}

	fmt.Println(total)
}
