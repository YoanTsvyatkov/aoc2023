package main

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getUniqueChars(str string) map[rune]bool {
	result := make(map[rune]bool)
	for _, ch := range str {
		result[ch] = true
	}
	return result
}

func isFiveOfKind(str string) bool {
	return len(getUniqueChars(str)) == 1
}

func isFourOfAKind(str string) bool {
	s := strings.Split(str, "")
	slices.SortFunc(s, func(a, b string) int {
		return strings.Compare(a, b) * -1
	})

	if s[0] == s[1] && s[1] == s[2] && s[2] == s[3] && s[3] != s[4] {
		return true
	}

	slices.SortFunc(s, func(a, b string) int {
		return strings.Compare(a, b)
	})

	return s[0] == s[1] && s[1] == s[2] && s[2] == s[3] && s[3] != s[4]
}

func isFullHouse(str string) bool {
	s := strings.Split(str, "")

	slices.SortFunc(s, func(a, b string) int {
		return strings.Compare(a, b) * -1
	})

	if s[0] == s[1] && s[1] == s[2] && s[3] == s[4] && s[3] != s[2] {
		return true
	}

	slices.SortFunc(s, func(a, b string) int {
		return strings.Compare(a, b)
	})

	return s[0] == s[1] && s[1] == s[2] && s[3] == s[4] && s[3] != s[2]
}

func isOnePair(str string) bool {
	return len(getUniqueChars(str)) == 4
}

func isThreeOfAKind(str string) bool {
	s := strings.Split(str, "")
	countOfCh := make(map[string]int)

	for _, ch := range s {
		cnt, found := countOfCh[ch]

		if !found {
			countOfCh[ch] = 1
		} else {
			countOfCh[ch] = cnt + 1
		}
	}

	threeCh := ""

	for _, ch := range s {
		cnt := countOfCh[ch]

		if cnt == 3 {
			threeCh = ch
			continue
		}

		if cnt >= 4 {
			return false
		}

		if threeCh != "" && cnt == 2 {
			return false
		}
	}

	return threeCh != ""
}

func isHighCard(str string) bool {
	return len(getUniqueChars(str)) == 5
}

func getCountOfJ(str string) int {
	cnt := 0
	for _, ch := range str {
		if ch == 'J' {
			cnt++
		}
	}
	return cnt
}

func getValueOfHand(str string) int {
	if isFiveOfKind(str) {
		return 7
	}

	if isFourOfAKind(str) {
		return 6
	}

	if isFullHouse(str) {
		return 5
	}

	if isThreeOfAKind(str) {
		return 4
	}

	if isOnePair(str) {
		return 2
	}

	if isHighCard(str) {
		return 1
	}

	return 3
}

func getValueOfHandWithJoker(str string) int {
	valueOfHand := getValueOfHand(str)

	cntOfJ := getCountOfJ(str)
	if (valueOfHand == 6 || valueOfHand == 5) && cntOfJ != 0 {
		valueOfHand = 7
	}

	if valueOfHand == 4 && cntOfJ != 0 {
		valueOfHand = 6
	}

	if valueOfHand == 3 && cntOfJ != 0 {
		if cntOfJ == 2 {
			valueOfHand = 6
		} else {
			valueOfHand = 5
		}
	}

	if valueOfHand == 2 && cntOfJ != 0 {
		valueOfHand = 4
	}

	if valueOfHand == 1 && cntOfJ != 0 {
		valueOfHand = 2
	}

	return valueOfHand
}

func getCardValue(str string) int {
	if str == "A" {
		return 14
	}

	if str == "K" {
		return 13
	}

	if str == "Q" {
		return 12
	}

	if str == "J" {
		return 1
	}

	if str == "T" {
		return 10
	}

	value, _ := strconv.Atoi(str)
	return value
}

func solution1() {
	cardsAndBids := utils.Map(strings.Split(utils.ReadFile("input"), "\n"), func(line string) [2]string {
		return [2]string(strings.Split(line, " "))
	})

	slices.SortFunc(cardsAndBids, func(a, b [2]string) int {
		handAValue := getValueOfHand(a[0])
		handBValue := getValueOfHand(b[0])

		if handAValue != handBValue {
			return handAValue - handBValue
		}

		chars1 := strings.Split(a[0], "")
		chars2 := strings.Split(b[0], "")
		for i := 0; i < len(chars1); i++ {
			cardValue1 := getCardValue(chars1[i])
			cardValue2 := getCardValue(chars2[i])
			if cardValue1 != cardValue2 {
				return cardValue1 - cardValue2
			}
		}

		return 0
	})

	fmt.Println(cardsAndBids)

	multiplier := 1
	result := utils.Reduce(cardsAndBids, func(acc int, handAndBid [2]string) int {
		bid, _ := strconv.Atoi(handAndBid[1])
		res := bid * multiplier
		multiplier++
		return acc + res
	}, 0)

	fmt.Println(result)

}

func solution2() {
	cardsAndBids := utils.Map(strings.Split(utils.ReadFile("input"), "\n"), func(line string) [2]string {
		return [2]string(strings.Split(line, " "))
	})

	slices.SortFunc(cardsAndBids, func(a, b [2]string) int {
		handAValueWithJoker := getValueOfHandWithJoker(a[0])
		handBValueWithJoker := getValueOfHandWithJoker(b[0])

		if handAValueWithJoker != handBValueWithJoker {
			return handAValueWithJoker - handBValueWithJoker
		}

		chars1 := strings.Split(a[0], "")
		chars2 := strings.Split(b[0], "")
		for i := 0; i < len(chars1); i++ {
			cardValue1 := getCardValue(chars1[i])
			cardValue2 := getCardValue(chars2[i])
			if cardValue1 != cardValue2 {
				return cardValue1 - cardValue2
			}
		}

		return 0
	})

	multiplier := 1
	result := utils.Reduce(cardsAndBids, func(acc int, handAndBid [2]string) int {
		bid, _ := strconv.Atoi(handAndBid[1])
		res := bid * multiplier
		multiplier++
		return acc + res
	}, 0)

	fmt.Println(result)
}

func main() {
	// solution1()
	solution2()
	// fmt.Println(isFullHouse("2a2a2"))
}
