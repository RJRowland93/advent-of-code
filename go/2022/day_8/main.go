package main

import (
	"advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(input []string) [][]int {
	result := make([][]int, len(input))

	for i, in := range input {
		tree := make([]int, len(in))
		split := strings.Split(in, "")
		for j, t := range split {
			n, _ := strconv.Atoi(t)
			tree[j] = n
		}
		result[i] = tree
	}

	return result
}

// define inside Part1 to not copy trees every iteration?
func IsVisible(trees [][]int, r, c int) bool {
	tree := trees[r][c]

	left := true
	for i := 0; i < c; i++ {
		if trees[r][i] > tree {
			fmt.Println("left: ", r, i, trees[r][i], tree)
			left = false
			break
		}
	}
	right := true
	for i := c + 1; i < len(trees[0]); i++ {
		if trees[r][i] > tree {
			fmt.Println("right: ", r, i, trees[r][i], tree)
			right = false
			break
		}
	}
	up := true
	for i := 0; i < r; i++ {
		if trees[i][c] > tree {
			fmt.Println("up: ", i, c, trees[i][c], tree)
			up = false
			break
		}
	}
	down := true
	for i := r + 1; i < len(trees); i++ {
		if trees[i][c] > tree {
			fmt.Println("down: ", i, c, trees[i][c], tree)
			down = false
			break
		}
	}

	return up || down || left || right
}

func Part1(trees [][]int) int {
	rows := len(trees)
	cols := len(trees[0])

	// count all edges
	visible := rows*2 + cols*2 - 4
	fmt.Println("edges: ", visible)

	fmt.Println(rows, cols)

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if IsVisible(trees, r, c) {
				visible++
			}
		}
	}

	return visible
}

func main() {
	input := common.LoadFileString()
	trees := ParseInput(input)
	fmt.Println(Part1(trees))
}
