package main

import (
	"advent-of-code/common"
	"fmt"
)

func Part1(input []int) int {
	prev := input[0]
	c := 0

	for _, num := range input {
		if num > prev {
			c++
		}
		prev = num
	}

	return c
}

func Part2(input []int) int {
	prev := input[0] + input[1] + input[2]
	c := 0

	for i := 0; i < len(input)-2; i++ {
		sum := input[i] + input[i+1] + input[i+2]
		if sum > prev {
			c++
		}
		prev = sum
	}

	return c
}

func main() {
	input := common.LoadFileInt()
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
