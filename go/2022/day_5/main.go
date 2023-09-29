package main

import (
	"advent-of-code/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// [D]                     [N] [F]
// [H] [F]             [L] [J] [H]
// [R] [H]             [F] [V] [G] [H]
// [Z] [Q]         [Z] [W] [L] [J] [B]
// [S] [W] [H]     [B] [H] [D] [C] [M]
// [P] [R] [S] [G] [J] [J] [W] [Z] [V]
// [W] [B] [V] [F] [G] [T] [T] [T] [P]
// [Q] [V] [C] [H] [P] [Q] [Z] [D] [W]
//  1   2   3   4   5   6   7   8   9

func ParseInput(input []string) [][3]int {
	parsed := make([][3]int, len(input))
	for i, s := range input {
		split := strings.Split(s, " ")
		amt, _ := strconv.Atoi(split[1])
		parsed[i][0] = amt
		from, _ := strconv.Atoi(split[3])
		parsed[i][1] = from
		to, _ := strconv.Atoi(split[5])
		parsed[i][2] = to
	}
	return parsed
}

func Part1(stacks [][]rune, instructions [][3]int) string {

	for _, i := range instructions {
		from := i[1] - 1
		to := i[2] - 1
		fmt.Println(i, stacks[from], stacks[to])
		amt := len(stacks[from]) - i[0]
		move := stacks[from][amt:]
		sort.SliceStable(move, func(i, j int) bool {
			return i > j
		})
		stacks[from] = stacks[from][:amt]
		for _, m := range move {
			stacks[to] = append(stacks[to], m)
		}
		fmt.Println(i, stacks[from], stacks[to])
	}

	result := make([]string, len(stacks))
	for j, s := range stacks {
		result[j] = string(s[len(s)-1])
	}
	return strings.Join(result, "")
}

func Part2(stacks [][]rune, instructions [][3]int) string {

	for _, i := range instructions {
		from := i[1] - 1
		to := i[2] - 1
		fmt.Println(i, stacks[from], stacks[to])
		amt := len(stacks[from]) - i[0]
		move := stacks[from][amt:]
		stacks[from] = stacks[from][:amt]
		for _, m := range move {
			stacks[to] = append(stacks[to], m)
		}
		fmt.Println(i, stacks[from], stacks[to])
	}

	result := make([]string, len(stacks))
	for j, s := range stacks {
		result[j] = string(s[len(s)-1])
	}
	return strings.Join(result, "")
}

func main() {
	// 1. split instructions into tuples
	// 2. slice n items from x stack
	// 4. push items onto y stack
	// 5. pop top off all stacks

	start := [][]rune{
		{'Q', 'W', 'P', 'S', 'Z', 'R', 'H', 'D'},
		{'V', 'B', 'R', 'W', 'Q', 'H', 'F'},
		{'C', 'V', 'S', 'H'},
		{'H', 'F', 'G'},
		{'P', 'G', 'J', 'B', 'Z'},
		{'Q', 'T', 'J', 'H', 'W', 'F', 'L'},
		{'Z', 'T', 'W', 'D', 'L', 'V', 'J', 'N'},
		{'D', 'T', 'Z', 'C', 'J', 'G', 'H', 'F'},
		{'W', 'P', 'V', 'M', 'B', 'H'},
	}
	input := common.LoadFileString()
	instructions := ParseInput(input)
	// result := Part1(start, instructions)
	result := Part2(start, instructions)
	// fmt.Printf("%v\n", start)
	fmt.Printf("%s\n", result)
	// Part1()
}
