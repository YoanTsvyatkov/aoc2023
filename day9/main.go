package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func getNextValue(nums []int) int {
	sequances := make([][]int, 0)
	sequances = append(sequances, nums)

	for {
		newSequance := make([]int, len(nums)-1)

		hasNonZero := false
		for i := 0; i < len(nums)-1; i++ {
			newSequance[i] = nums[i+1] - nums[i]
			if newSequance[i] != 0 {
				hasNonZero = true
			}
		}

		sequances = append(sequances, newSequance)

		if !hasNonZero {
			break
		}

		nums = newSequance
	}

	for i := len(sequances) - 1; i >= 0; i-- {
		if i == len(sequances)-1 {
			sequances[i] = append(sequances[i], 0)
		} else {
			valueDown := sequances[i+1][len(sequances[i+1])-1]
			valueLeft := sequances[i][len(sequances[i])-1]
			sequances[i] = append(sequances[i], valueDown+valueLeft)
		}
	}

	return sequances[0][len(sequances[0])-1]
}

func solution1() {
	lines := strings.Split(utils.ReadFile("input"), "\n")

	sum := utils.Reduce(lines, func(acc int, line string) int {
		nums := utils.Map(strings.Fields(line), func(strNum string) int {
			num, _ := strconv.Atoi(strNum)
			return num
		})
		nextValue := getNextValue(nums)
		return acc + nextValue
	}, 0)

	fmt.Println(sum)
}

func getLeftValue(nums []int) int {
	sequances := make([][]int, 0)
	sequances = append(sequances, nums)

	for {
		newSequance := make([]int, len(nums)-1)

		hasNonZero := false
		for i := 0; i < len(nums)-1; i++ {
			newSequance[i] = nums[i+1] - nums[i]
			if newSequance[i] != 0 {
				hasNonZero = true
			}
		}

		sequances = append(sequances, newSequance)

		if !hasNonZero {
			break
		}

		nums = newSequance
	}

	for i := len(sequances) - 1; i >= 0; i-- {
		if i == len(sequances)-1 {
			sequances[i] = append([]int{0}, sequances[i]...)
		} else {
			valueDown := sequances[i+1][0]
			valueRight := sequances[i][0]
			sequances[i] = append([]int{valueRight - valueDown}, sequances[i]...)
		}
	}

	return sequances[0][0]
}

func solution2() {
	lines := strings.Split(utils.ReadFile("input"), "\n")

	sum := utils.Reduce(lines, func(acc int, line string) int {
		nums := utils.Map(strings.Fields(line), func(strNum string) int {
			num, _ := strconv.Atoi(strNum)
			return num
		})
		nextValue := getLeftValue(nums)
		return acc + nextValue
	}, 0)

	fmt.Println(sum)
}

func main() {
	solution1()
	solution2()
}
