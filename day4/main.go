package main

import (
	"aoc2023/utils"
	"strconv"
	"strings"
)

func Contains(nums []int, value int) bool {
	for _, curr := range nums {
		if curr == value {
			return true
		}
	}
	return false
}

func solution1() {
	lines := strings.Split(utils.ReadFile("input"), "\n")

	sum := utils.Reduce(lines, func(acc int, curr string) int {
		parts := strings.Split(curr, "|")

		userNums := utils.Map(strings.Fields(strings.TrimSpace(parts[1])), func(value string) int {
			v, _ := strconv.Atoi(value)
			return v
		})

		winningNums := utils.Map(strings.Fields(strings.Split(parts[0], ": ")[1]), func(value string) int {
			v, _ := strconv.Atoi(value)
			return v
		})

		count := 0
		for _, num := range userNums {
			if Contains(winningNums, num) {
				if count > 1 {
					count *= 2
				} else {
					count++
				}
			}
		}

		return acc + count
	}, 0)

	println(sum)
}

// Sum of cards
func solution2() {
	lines := strings.Split(utils.ReadFile("input"), "\n")

	cardCount := make(map[int]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cardCount[i+1] = 1
	}

	sum := utils.Reduce(lines, func(acc int, curr string) int {
		parts := strings.Split(curr, "|")
		gameNum, _ := strconv.Atoi(strings.Fields(strings.Split(parts[0], ":")[0])[1])

		userNums := utils.Map(strings.Fields(strings.TrimSpace(parts[1])), func(value string) int {
			v, _ := strconv.Atoi(value)
			return v
		})

		winningNums := utils.Map(strings.Fields(strings.Split(parts[0], ": ")[1]), func(value string) int {
			v, _ := strconv.Atoi(value)
			return v
		})

		count := 0
		for _, num := range userNums {
			if Contains(winningNums, num) {
				count++
			}
		}

		for i := 1; i <= count; i++ {
			cardCount[gameNum+i] += cardCount[gameNum]
		}

		return cardCount[gameNum] + acc
	}, 0)

	println(sum)
}

func main() {
	solution1()
	solution2()
}
