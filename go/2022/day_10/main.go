package main

import (
	"advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	wait   int
	amount int
}

func ParseInput(input []string) []Instruction {
	instructions := make([]Instruction, len(input))

	for i, in := range input {
		split := strings.Split(in, " ")
		wait := 1
		if split[0] == "addx" {
			wait = 2
		}
		amount := 0
		if len(split) == 2 {
			n, _ := strconv.Atoi(split[1])
			amount = n
		}
		instructions[i] = Instruction{wait, amount}
	}
	return instructions
}

func Part1(ins []Instruction) int {
	cycle := 1
	signal := 1
	total := 0

	for _, in := range ins {
		for j := 0; j < in.wait; j++ {
			if cycle == 20 || (cycle-20)%40 == 0 {
				strength := cycle * signal
				total += strength
			}
			cycle++
		}
		signal += in.amount
	}

	return total
}

func Part2(ins []Instruction) [][]string {
	cycle := 0
	signal := 1
	sprite := []int{0, 1, 2}
	overlap := make([]string, 240)

	for x, _ := range overlap {
		overlap[x] = " "
	}

	for _, in := range ins {
		for j := 0; j < in.wait; j++ {
			for k := 0; k < 3; k++ {
				if sprite[k] == cycle%40 {
					overlap[cycle] = "#"
				}
			}
			// fmt.Println(cycle, signal, sprite, overlap[cycle])
			cycle++
		}
		signal += in.amount
		sprite = []int{signal - 1, signal, signal + 1}
	}
	fmt.Println(cycle)

	crt := make([][]string, 6)
	crt[0] = overlap[1:40]
	crt[1] = overlap[41:80]
	crt[2] = overlap[81:120]
	crt[3] = overlap[121:160]
	crt[4] = overlap[161:200]
	crt[5] = overlap[201:240]

	return crt
}

func main() {
	input := common.LoadFileString()
	parsed := ParseInput(input)
	// fmt.Println(Part1(parsed))

	r := Part2(parsed)
	for _, z := range r {
		fmt.Println(z)
	}
}
