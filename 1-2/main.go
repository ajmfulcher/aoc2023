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

func digitAsString(d byte) (string, error) {
	char := string(d)
	_, err := strconv.Atoi(char)
	return char, err
}

func isNumberStart(input string) string {
	if len(input) >= 3 && input[0:3] == "one" {
		return "1"
	} else if len(input) >= 3 && input[0:3] == "two" {
		return "2"
	} else if len(input) >= 5 && input[0:5] == "three" {
		return "3"
	} else if len(input) >= 4 && input[0:4] == "four" {
		return "4"
	} else if len(input) >= 4 && input[0:4] == "five" {
		return "5"
	} else if len(input) >= 3 && input[0:3] == "six" {
		return "6"
	} else if len(input) >= 5 && input[0:5] == "seven" {
		return "7"
	} else if len(input) >= 5 && input[0:5] == "eight" {
		return "8"
	} else if len(input) >= 4 && input[0:4] == "nine" {
		return "9"
	} else {
		return ""
	}
}

func isNumberEnd(input string) string {
	end := len(input)
	if len(input) >= 3 && input[end-3:end] == "one" {
		return "1"
	} else if len(input) >= 3 && input[end-3:end] == "two" {
		return "2"
	} else if len(input) >= 5 && input[end-5:end] == "three" {
		return "3"
	} else if len(input) >= 4 && input[end-4:end] == "four" {
		return "4"
	} else if len(input) >= 4 && input[end-4:end] == "five" {
		return "5"
	} else if len(input) >= 3 && input[end-3:end] == "six" {
		return "6"
	} else if len(input) >= 5 && input[end-5:end] == "seven" {
		return "7"
	} else if len(input) >= 5 && input[end-5:end] == "eight" {
		return "8"
	} else if len(input) >= 4 && input[end-4:end] == "nine" {
		return "9"
	} else {
		return ""
	}
}

func main() {
	input := inputToArray("input.txt")

	var result int = 0

	for _, line := range input {
		var asString string = ""
		for i := 0; i < len(line); i++ {
			char, err := digitAsString(line[i])
			if err == nil {
				asString = asString + char
				break
			}
			isNum := isNumberStart(line[i : len(line)-1])
			if isNum != "" {
				asString = asString + isNum
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char, err := digitAsString(line[i])
			if err == nil {
				asString = asString + char
				break
			}
			isNum := isNumberEnd(line[0 : i+1])
			if isNum != "" {
				asString = asString + isNum
				break
			}
		}
		// fmt.Println(asString)
		res, _ := strconv.Atoi(asString)
		result += res
	}
	fmt.Println(result)
}
