package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var strToValue = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func readFileLines(fileName string) ([]string, error) {
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

func getKeysFromMap() []string {
	keys := make([]string, 0, len(strToValue))
	for k, _ := range strToValue {
		keys = append(keys, k)
	}
	return keys
}

func getValuesFromMap() []int {
	values := make([]int, 0, len(strToValue))
	for _, v := range strToValue {
		values = append(values, v)
	}
	return values
}

var digitsAsText = getKeysFromMap()
var digits = getValuesFromMap()

func getSumOfFirstAndLastDigitsInStr(str string) int {
	firstIndex, firstNum := -1, 0
	secondIndex, secondNum := -1, 0

	for _, digitAsText := range digitsAsText {
		first := strings.Index(str, digitAsText)
		last := strings.LastIndex(str, digitAsText)

		if first == -1 && last == -1 {
			continue
		}

		if last > secondIndex || secondIndex == -1 {
			secondIndex = last
			secondNum = strToValue[digitAsText]
		}

		if first < firstIndex || firstIndex == -1 {
			firstIndex = first
			firstNum = strToValue[digitAsText]
		}
	}

	for _, digit := range digits {
		first := strings.Index(str, strconv.Itoa(digit))
		last := strings.LastIndex(str, strconv.Itoa(digit))

		if first == -1 && last == -1 {
			continue
		}

		if last > secondIndex || secondIndex == -1 {
			secondIndex = last
			secondNum = digit
		}

		if first < firstIndex || firstIndex == -1 {
			firstIndex = first
			firstNum = digit
		}
	}

	return firstNum*10 + secondNum
}

func solution1() {
	fileLines, err := readFileLines("../input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for _, line := range fileLines {
		sum += getSumOfFirstAndLastDigitsInStr(line)
	}

	println(sum)
}

func main() {
	solution1()
}
