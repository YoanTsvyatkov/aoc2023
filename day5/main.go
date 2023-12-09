package main

import (
	"aoc2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func solution1() {
	lines := strings.Split(utils.ReadFile("input"), "\n")
	seeds := utils.Map(strings.Fields(strings.Split(lines[0], ": ")[1]), func(seed string) int {
		value, _ := strconv.Atoi(seed)
		return value
	})

	seedMaps := lines[2:]

	seedRanges := make([][][3]int, 0)
	currRanges := make([][3]int, 0)

	for _, currLine := range seedMaps {
		if strings.HasSuffix(currLine, ":") {
			//DO nothing
		} else if len(currLine) == 0 {
			seedRanges = append(seedRanges, currRanges)
			currRanges = make([][3]int, 0)
		} else {
			newRange := utils.Map(strings.Fields(currLine), func(seed string) int {
				value, _ := strconv.Atoi(seed)
				return value
			})
			currRanges = append(currRanges, [3]int(newRange))
		}
	}

	seedRanges = append(seedRanges, currRanges)
	min := math.MaxInt64

	for _, seed := range seeds {
		currentValue := seed
		// 81 - 25 => 56

		for _, currentSeedArrOfRanges := range seedRanges {
			for _, currRange := range currentSeedArrOfRanges {
				startSource, endSource := currRange[1], currRange[1]+currRange[2]-1
				if currentValue >= startSource && currentValue <= endSource {
					currentValue = currRange[0] + currentValue - startSource
					break
				}
			}

		}

		if min > currentValue {
			min = currentValue
		}
	}

	fmt.Println(min)
}

func findMin(startValue int, endValue int, currentRangIndex int, seedRanges [][][3]int) [2]int {
	if currentRangIndex >= len(seedRanges) {
		return [2]int{startValue, endValue}
	}

	ranges := seedRanges[currentRangIndex]

	for _, currRange := range ranges {
		startSource, endSource := currRange[1], currRange[1]+currRange[2]-1
		if startValue >= startSource && endValue <= endSource {
			return findMin(
				currRange[0]+startValue-startSource, (currRange[0]+currRange[2]-1)+endValue-endSource, currentRangIndex+1, seedRanges)
		} else if startValue < startSource && endValue < startSource || startValue > endSource {
			continue
		} else if startValue > startSource {
			x, y := startValue, endSource
			z, n := endSource+1, endValue
			first := findMin(x, y, currentRangIndex+1, seedRanges)
			second := findMin(z, n, currentRangIndex+1, seedRanges)
			if first[0] > second[0] {
				return second
			} else {
				return first
			}
		} else {
			x, y := startValue, startSource
			z, n := startSource+1, endValue
			first := findMin(x, y, currentRangIndex+1, seedRanges)
			second := findMin(z, n, currentRangIndex+1, seedRanges)
			if first[0] > second[0] {
				return second
			} else {
				return first
			}
		}
	}

	return findMin(
		startValue, endValue, currentRangIndex+1, seedRanges)
}

func solution2() {
	lines := strings.Split(utils.ReadFile("input"), "\n")
	seeds := utils.Map(strings.Fields(strings.Split(lines[0], ": ")[1]), func(seed string) int {
		value, _ := strconv.Atoi(seed)
		return value
	})

	seedMaps := lines[2:]

	seedRanges := make([][][3]int, 0)
	currRanges := make([][3]int, 0)

	for _, currLine := range seedMaps {
		if strings.HasSuffix(currLine, ":") {
			//DO nothing
		} else if len(currLine) == 0 {
			seedRanges = append(seedRanges, currRanges)
			currRanges = make([][3]int, 0)
		} else {
			newRange := utils.Map(strings.Fields(currLine), func(seed string) int {
				value, _ := strconv.Atoi(seed)
				return value
			})
			currRanges = append(currRanges, [3]int(newRange))
		}
	}

	seedRanges = append(seedRanges, currRanges)

	min := math.MaxInt64
	start := 0
	end := 1

	for end < len(seeds) {
		seedStartValue := seeds[start]
		seedEndValue := seeds[start] + seeds[end]
		result := findMin(seedStartValue, seedEndValue, 0, seedRanges)

		if min > result[0] {
			min = result[0]
		}

		start = end + 1
		end = start + 1
	}

	fmt.Println(min)

}

func main() {
	solution2()
}
