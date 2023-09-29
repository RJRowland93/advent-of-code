package main

import (
	"advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(input []string) [][2][2]int {
	result := make([][2][2]int, len(input))

	for i, in := range input {
		pairs := strings.Split(in, ",")
		for j, p := range pairs {
			span := strings.Split(p, "-")
			x, _ := strconv.Atoi(span[0])
			y, _ := strconv.Atoi(span[1])
			result[i][j] = [2]int{x, y}
		}
	}

	return result
}

func Part1(input [][2][2]int) int {
	count := 0
	var rejects [][2][2]int

	for _, p := range input {
		// fmt.Println(p)
		if p[0][0] > p[1][0] || (p[0][0] == p[1][0] && p[0][1] < p[1][1]) {
			fmt.Println("swapping: ", p[0], p[1])
			p[0], p[1] = p[1], p[0]
			fmt.Println("after: ", p[0], p[1])
		}

		if p[0][0] <= p[1][0] && p[0][1] >= p[1][1] {
			fmt.Println(p[0], " spans ", p[1])
			count++
		} else {
			rejects = append(rejects, p)
		}
	}

	for _, r := range rejects {
		fmt.Println(r)
	}
	return count
}

func Part2(input [][2][2]int) int {
	count := 0
	var rejects [][2][2]int

	for _, p := range input {
		// fmt.Println(p)
		if p[0][0] > p[1][0] || (p[0][0] == p[1][0] && p[0][1] < p[1][1]) {
			fmt.Println("swapping: ", p[0], p[1])
			p[0], p[1] = p[1], p[0]
			fmt.Println("after: ", p[0], p[1])
		}

		if p[0][1] >= p[1][0] {
			fmt.Println(p[0], " spans ", p[1])
			count++
		} else {
			rejects = append(rejects, p)
		}
	}

	for _, r := range rejects {
		fmt.Println(r)
	}
	return count
}

func main() {
	input := common.LoadFileString()
	parsed := ParseInput(input)
	// result := Part1(parsed)
	result := Part2(parsed)
	fmt.Println(len(input), result)
}
