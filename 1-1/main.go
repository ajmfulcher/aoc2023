package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var result int = 0

	for _, line := range input {
		var asString string = ""
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			_, err := strconv.Atoi(char)
			if err == nil {
				asString = asString + char
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := string(line[i])
			_, err := strconv.Atoi(char)
			if err == nil {
				asString = asString + char
				break
			}
		}
		// fmt.Println(asString)
		res, _ := strconv.Atoi(asString)
		result += res
	}
	fmt.Println(result)
}
