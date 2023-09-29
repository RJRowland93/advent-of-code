package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFileString() []string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("%s", err)
	}

	input := strings.Split(string(b), "\n")

	return input
}

func LoadFileInt() []int {
	input := LoadFileString()

	ints := make([]int, len(input))

	for i, s := range input {
		cur, _ := strconv.Atoi(s)
		ints[i] = cur
	}

	return ints
}
