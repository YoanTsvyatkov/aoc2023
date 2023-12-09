package main

import (
	"aoc2023/utils"
	"fmt"
	"strings"
)

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(value string) *Node {
	node := Node{value: value, left: nil, right: nil}
	return &node
}

var lines = strings.Split(utils.ReadFile("input"), "\n")
var moves = strings.Split(lines[0], "")
var nodesList = lines[2:]

func getMovesTillEnd(root *Node, index int) int {
	if root.value == "ZZZ" {
		return 0
	}

	move := moves[index]
	if move == "L" {
		return 1 + getMovesTillEnd(root.left, (index+1)%len(moves))
	} else {
		return 1 + getMovesTillEnd(root.right, (index+1)%len(moves))
	}
}

func getMovesTillEndSlice(nodes []*Node, index int) int {
	allEndingWithZ := true
	for _, node := range nodes {
		if node.value[2] != 'Z' {
			allEndingWithZ = false
			break
		}
	}

	if allEndingWithZ {
		return 0
	}

	move := moves[index]

	for i := 0; i < len(nodes); i++ {
		if move == "L" {
			nodes[i] = nodes[i].left
		} else {
			nodes[i] = nodes[i].right
		}
	}

	return 1 + getMovesTillEndSlice(nodes, (index+1)%len(moves))
}

func solution1() {
	nodesMap := make(map[string]*Node)

	for _, nodeStr := range nodesList {
		nodeKey := strings.Split(nodeStr, " = ")[0]
		nodeChilds := strings.Split(nodeStr, " = ")[1]

		newNode, ok := nodesMap[nodeKey]
		if !ok {
			newNode = NewNode(nodeKey)
			nodesMap[nodeKey] = newNode
		}

		left := strings.Split(nodeChilds, ", ")[0][1:]
		right := strings.Split(nodeChilds, ", ")[1][0:3]

		leftNode, ok := nodesMap[left]
		if !ok {
			leftNode = NewNode(left)
			nodesMap[left] = leftNode
		}

		rightNode, ok := nodesMap[right]
		if !ok {
			rightNode = NewNode(right)
			nodesMap[right] = rightNode
		}

		newNode.left = leftNode
		newNode.right = rightNode
	}

	root := nodesMap["AAA"]
	cnt := getMovesTillEnd(root, 0)
	fmt.Println(cnt)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func solution2() {
	nodesMap := make(map[string]*Node)
	endingWithA := make([]*Node, 0)

	for _, nodeStr := range nodesList {
		nodeKey := strings.Split(nodeStr, " = ")[0]
		nodeChilds := strings.Split(nodeStr, " = ")[1]

		newNode, ok := nodesMap[nodeKey]
		if !ok {
			newNode = NewNode(nodeKey)
			nodesMap[nodeKey] = newNode
		}

		left := strings.Split(nodeChilds, ", ")[0][1:]
		right := strings.Split(nodeChilds, ", ")[1][0:3]

		leftNode, ok := nodesMap[left]
		if !ok {
			leftNode = NewNode(left)
			nodesMap[left] = leftNode
		}

		rightNode, ok := nodesMap[right]
		if !ok {
			rightNode = NewNode(right)
			nodesMap[right] = rightNode
		}

		newNode.left = leftNode
		newNode.right = rightNode

		if strings.Split(newNode.value, "")[2] == "A" {
			endingWithA = append(endingWithA, newNode)
		}
	}

	stepsTillEndingZ := make(map[string]int)
	initialNodeKeys := make([]string, 0)

	for _, node := range endingWithA {
		cnt := 0
		moveIndex := 0
		initialNodeValue := node.value

		for {
			if node.value[2] == 'Z' {
				break
			}

			move := moves[moveIndex]

			if move == "L" {
				node = node.left
			} else {
				node = node.right
			}

			cnt++
			moveIndex = (moveIndex + 1) % len(moves)
		}

		stepsTillEndingZ[initialNodeValue] = cnt
		initialNodeKeys = append(initialNodeKeys, initialNodeValue)
	}

	result := utils.Reduce(initialNodeKeys[1:], func(acc int, curr string) int {
		return LCM(acc, stepsTillEndingZ[curr])
	}, stepsTillEndingZ[initialNodeKeys[0]])

	fmt.Println(result)
}

func main() {
	// solution1()
	solution2()
}
