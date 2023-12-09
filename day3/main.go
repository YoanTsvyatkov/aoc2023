package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(ch rune) bool {
	return ch != '.' && !unicode.IsDigit(ch)
}

func solution1() {
	rows := strings.Split(utils.ReadFile("input.txt"), "\n")
	sum := 0
	currentNumber := -1
	isPartOfSum := false

	for i, row := range rows {
		currentNumber = -1
		isPartOfSum = false

		for j, ch := range row {
			if unicode.IsDigit(ch) {
				digit, _ := strconv.Atoi(string(ch))
				//  TOP, TOP-LEFT,  TOP-RIGHT, LEFT, RIGHT,  BOTTOM-RIGHT, BOTTOM-LEFT, BOTTOM
				if i-1 >= 0 && isSymbol(rune(rows[i-1][j])) || i+1 < len(rows) && isSymbol(rune(rows[i+1][j])) ||
					j-1 >= 0 && isSymbol(rune(rows[i][j-1])) || j+1 < len(row) && isSymbol(rune(rows[i][j+1])) ||
					i-1 >= 0 && j-1 >= 0 && isSymbol(rune(rows[i-1][j-1])) || i+1 < len(rows) && j+1 < len(row) && isSymbol(rune(rows[i+1][j+1])) ||
					i+1 < len(rows) && j-1 >= 0 && isSymbol(rune(rows[i+1][j-1])) || i-1 >= 0 && j+1 < len(row) && isSymbol(rune(rows[i-1][j+1])) {
					isPartOfSum = true
				}
				if currentNumber == -1 {
					currentNumber = digit
				} else {
					currentNumber = currentNumber*10 + digit
				}
			} else {
				if currentNumber != -1 {
					if isPartOfSum {
						sum += currentNumber
					}
					currentNumber = -1
					isPartOfSum = false
				}
			}
		}

		if currentNumber != -1 {
			if isPartOfSum {
				sum += currentNumber
			}
			currentNumber = -1
			isPartOfSum = false
		}
	}

	fmt.Println(sum)
}

func solution2() {
	rows := strings.Split(utils.ReadFile("input.txt"), "\n")
	currentNumber := -1
	isNumberAdjacentToSymbol := false
	isDigitInPartNumber := make([][]bool, len(rows))
	digitToNumber := make([][]int, len(rows))

	for i, row := range rows {
		currentNumber = -1
		isNumberAdjacentToSymbol = false
		start := -1
		end := -1
		isDigitInPartNumber[i] = make([]bool, len(rows[i]))
		digitToNumber[i] = make([]int, len(rows[i]))

		for j, ch := range row {
			if unicode.IsDigit(ch) {
				digit, _ := strconv.Atoi(string(ch))
				if i-1 >= 0 && isSymbol(rune(rows[i-1][j])) || i+1 < len(rows) && isSymbol(rune(rows[i+1][j])) ||
					j-1 >= 0 && isSymbol(rune(rows[i][j-1])) || j+1 < len(row) && isSymbol(rune(rows[i][j+1])) ||
					i-1 >= 0 && j-1 >= 0 && isSymbol(rune(rows[i-1][j-1])) || i+1 < len(rows) && j+1 < len(row) && isSymbol(rune(rows[i+1][j+1])) ||
					i+1 < len(rows) && j-1 >= 0 && isSymbol(rune(rows[i+1][j-1])) || i-1 >= 0 && j+1 < len(row) && isSymbol(rune(rows[i-1][j+1])) {
					isNumberAdjacentToSymbol = true
				}
				if currentNumber == -1 {
					currentNumber = digit
					start = j
				} else {
					currentNumber = currentNumber*10 + digit
				}
			} else {
				if currentNumber != -1 {
					if isNumberAdjacentToSymbol {
						end = j - 1
						for iter := start; iter <= end; iter++ {
							isDigitInPartNumber[i][iter] = true
							digitToNumber[i][iter] = currentNumber
						}
					}
					currentNumber = -1
					isNumberAdjacentToSymbol = false
				}
			}
		}

		if currentNumber != -1 {
			if isNumberAdjacentToSymbol {
				end = len(rows[i]) - 1
				for iter := start; iter <= end; iter++ {
					isDigitInPartNumber[i][iter] = true
					digitToNumber[i][iter] = currentNumber
				}
			}
			currentNumber = -1
			isNumberAdjacentToSymbol = false
		}
	}

	finalSum := 0
	for i, row := range rows {
		for j, ch := range rows[i] {
			if ch == '*' {
				countOfAdjacentNums := 0
				prod := 1

				if i+1 < len(rows) && isDigitInPartNumber[i+1][j] {
					countOfAdjacentNums++
					prod *= digitToNumber[i+1][j]
				} else if i+1 < len(rows) && j+1 < len(row) && isDigitInPartNumber[i+1][j+1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i+1][j+1]

					if j-1 >= 0 && isDigitInPartNumber[i+1][j-1] {
						countOfAdjacentNums++
						prod *= digitToNumber[i+1][j-1]
					}
				} else if i+1 < len(rows) && j-1 >= 0 && isDigitInPartNumber[i+1][j-1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i+1][j-1]
				}

				if j+1 < len(row) && isDigitInPartNumber[i][j+1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i][j+1]
				}
				if j-1 >= 0 && isDigitInPartNumber[i][j-1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i][j-1]
				}

				if i-1 >= 0 && isDigitInPartNumber[i-1][j] {
					countOfAdjacentNums++
					prod *= digitToNumber[i-1][j]
				} else if i-1 >= 0 && j+1 < len(row) && isDigitInPartNumber[i-1][j+1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i-1][j+1]

					if j-1 >= 0 && isDigitInPartNumber[i-1][j-1] {
						countOfAdjacentNums++
						prod *= digitToNumber[i-1][j-1]
					}
				} else if i-1 >= 0 && j-1 >= 0 && isDigitInPartNumber[i-1][j-1] {
					countOfAdjacentNums++
					prod *= digitToNumber[i-1][j-1]
				}

				if countOfAdjacentNums == 2 {
					finalSum += prod
				}
			}
		}
	}

	fmt.Println(finalSum)
}

func main() {
	// solution1()
	solution2()
}
