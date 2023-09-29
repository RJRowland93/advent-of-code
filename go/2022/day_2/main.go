package main

import (
	"advent-of-code/common"
	"fmt"
	"strings"
)

type Move int

const (
	A Move = 1 // rock
	B Move = 2 // paper
	C Move = 3 // scizzors

	X Move = 1
	Y Move = 2
	Z Move = 3

	lose = 0
	draw = 3
	win  = 6
)

func ParseInput(input []string) [][]string {
	result := make([][]string, len(input))
	for i, s := range input {
		split := strings.Split(s, " ")
		result[i] = split

	}
	return result
}

func ScoreRound(moves []string) int {
	score := 0

	switch moves[0] {
	case "A":
		switch moves[1] {
		case "X":
			score = 1 + 3
		case "Y":
			score = 2 + 6
		case "Z":
			score = 3 + 0
		}
	case "B":
		switch moves[1] {
		case "X":
			score = 1 + 0
		case "Y":
			score = 2 + 3
		case "Z":
			score = 3 + 6
		}
	case "C":
		switch moves[1] {
		case "X":
			score = 1 + 6
		case "Y":
			score = 2 + 0
		case "Z":
			score = 3 + 3
		}
	}
	return score
}

func Part1(input []string) int {
	moves := ParseInput(input)
	sum := 0
	for _, m := range moves {
		sum += ScoreRound(m)
	}
	return sum
}

// X = lose
// Y = draw
// Z = win
func MatchRound(moves []string) int {
	score := 0

	switch moves[0] {
	case "A":
		switch moves[1] {
		case "X":
			// rock; lose = scissors
			score = 3 + 0
		case "Y":
			// rock; draw = rock
			score = 1 + 3
		case "Z":
			// rock; win = paper
			score = 2 + 6
		}
	case "B":
		switch moves[1] {
		case "X":
			// paper; lose = rock
			score = 1 + 0
		case "Y":
			// paper; draw = paper
			score = 2 + 3
		case "Z":
			// paper; win = scissors
			score = 3 + 6
		}
	case "C":
		switch moves[1] {
		case "X":
			// scissors; lose = paper
			score = 2 + 0
		case "Y":
			// scissors; draw = scissors
			score = 3 + 3
		case "Z":
			// scissors; win = rock
			score = 1 + 6
		}
	}
	return score
}

func Part2(input []string) int {
	moves := ParseInput(input)
	sum := 0
	for _, m := range moves {
		sum += MatchRound(m)
	}
	return sum
}

func main() {
	input := common.LoadFileString()
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
