package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func solution1() {
	lines := strings.Split(utils.ReadFile("input"), "\n")
	times := utils.Map(strings.Fields(strings.Split(lines[0], ":")[1]), func(time string) int {
		value, _ := strconv.Atoi(time)
		return value
	})
	records := utils.Map(strings.Fields(strings.Split(lines[1], ":")[1]), func(record string) int {
		value, _ := strconv.Atoi(record)
		return value
	})

	prod := 1
	for i := 0; i < len(times); i++ {
		record := records[i]
		time := times[i]
		recordCnt := 0
		speed := 1

		for holdTime := 1; holdTime < times[i]; holdTime++ {
			finishTime := (time - holdTime) * speed
			if finishTime > record {
				recordCnt++
			}
			speed++
		}

		prod *= recordCnt
	}

	fmt.Println(prod)
}

func solution2() {
	lines := strings.Split(utils.ReadFile("input"), "\n")
	time, _ := strconv.Atoi(
		utils.Reduce(strings.Fields(strings.Split(lines[0], ":")[1]), func(acc string, time string) string {
			return acc + time
		}, ""))
	record, _ := strconv.Atoi(
		utils.Reduce(strings.Fields(strings.Split(lines[1], ":")[1]), func(acc string, time string) string {
			return acc + time
		}, ""))

	recordCnt := 0
	speed := 1

	fmt.Println("Time: ", time)
	fmt.Println("Record: ", record)

	for holdTime := 1; holdTime < time; holdTime++ {
		finishTime := (time - holdTime) * speed
		if finishTime > record {
			recordCnt++
		}
		speed++
	}

	fmt.Println(recordCnt)
}
func main() {
	// solution1()
	solution2()
}
