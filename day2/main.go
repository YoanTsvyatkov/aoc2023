package main

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Hand struct {
	r int
	g int
	b int
}

type Game struct {
	gameId int
	hands  []Hand
}

func ParseGames() []Game {
	content := utils.ReadFile("input.txt")
	lines := strings.Split(content, "\n")
	allGames := []Game{}

	for _, line := range lines {
		game := strings.Split(line, ":")[0]
		allHandsStr := strings.Split(line, ":")[1]

		id, _ := strconv.Atoi(game[5:])
		hands := strings.Split(allHandsStr, ";")

		gameHands := []Hand{}
		for _, hand := range hands {
			colors := strings.Split(strings.TrimSpace(hand), ", ")
			currentHand := Hand{
				r: 0,
				g: 0,
				b: 0,
			}

			for _, color := range colors {
				numAndColor := strings.Split(color, " ")
				num, _ := strconv.Atoi(numAndColor[0])
				colorStr := numAndColor[1]

				if colorStr == "blue" {
					currentHand.b = num
				} else if colorStr == "red" {
					currentHand.r = num
				} else {
					currentHand.g = num
				}
			}

			gameHands = append(gameHands, currentHand)
		}

		allGames = append(allGames, Game{
			gameId: id,
			hands:  gameHands,
		})
	}

	return allGames
}

func solution1() {
	maxHand := Hand{
		r: 12,
		g: 13,
		b: 14,
	}

	games := ParseGames()
	sum := 0

	for _, game := range games {
		possibleGame := true
		for _, hand := range game.hands {
			if hand.b > maxHand.b || hand.g > maxHand.g || hand.r > maxHand.r {
				possibleGame = false
				break
			}
		}

		if possibleGame {
			sum += game.gameId
		}
	}

	fmt.Println(sum)
}

func solution2() {
	games := ParseGames()
	sum := 0

	for _, game := range games {
		minHand := Hand{
			r: 0,
			g: 0,
			b: 0,
		}

		for _, hand := range game.hands {
			if hand.b > minHand.b {
				minHand.b = hand.b
			}
			if hand.g > minHand.g {
				minHand.g = hand.g
			}
			if hand.r > minHand.r {
				minHand.r = hand.r
			}
		}

		sum += minHand.b * minHand.g * minHand.r
	}

	println(sum)
}

func main() {
	solution1()
	solution2()
}
