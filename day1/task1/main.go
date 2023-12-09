package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSumOfFirstAndLastDigitsInStr(str string) int {
	first := 0
	lastNum := 0
	foundFirst := false
	foundSecond := false

	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			if !foundFirst {
				first = int(ch-'0') * 10
				foundFirst = true
			} else {
				lastNum = int(ch - '0')
				foundSecond = true
			}
		}
	}

	if !foundSecond {
		return first + first/10
	}

	return first + lastNum
}

func readFile(fileName string) ([]string, error) {
	readFile, err := os.Open(fileName)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}

func solution1() {
	fileLines, err := readFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	fmt.Print(len(fileLines))
	for _, line := range fileLines {
		curr := getSumOfFirstAndLastDigitsInStr(line)
		fmt.Println(curr)
		sum += curr
	}

	println(sum)
}

func main() {
	solution1()
}
