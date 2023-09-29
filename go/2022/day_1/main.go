package main

import (
	"advent-of-code/common"
	"fmt"
	"sort"
	"strconv"
)

func Part1(input []string) int {
	// s := strings.Split(input, "\n")

	max := 0
	sum := 0

	for _, n := range input {

		if n == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			x, _ := strconv.Atoi(n)
			sum += x
		}
	}
	if sum > max {
		max = sum
	}

	return max
}

func Part2(input []string) int {
	// s := strings.Split(input, "\n")

	var elves []int
	sum := 0

	for _, n := range input {

		if n == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			x, _ := strconv.Atoi(n)
			sum += x
		}
	}
	elves = append(elves, sum)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	max := 0

	for _, x := range elves[:3] {
		max += x
	}

	return max
}

func main() {
	input := common.LoadFileString()

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
