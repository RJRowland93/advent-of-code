package main

import (
	"advent-of-code/common"
	"fmt"
)

func IsStringSliceUnique(str string) bool {
	found := make(map[rune]bool)

	for _, c := range str {
		if _, ok := found[c]; ok {
			return false
		}
		found[c] = true
	}
	return true
}

func Part1(numUnique int, input string) int {
	for i := numUnique; i < len(input)-1; i++ {
		if IsStringSliceUnique(input[i-numUnique : i]) {
			return i
		}
	}
	return 0
}

func main() {
	input := common.LoadFileString()[0]
	fmt.Println(Part1(4, input))
	fmt.Println(Part1(14, input))
}
