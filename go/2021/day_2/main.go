package main

import (
	"advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

type Direction struct {
	direction string
	value     int
}

func ParseInput(input []string) []Direction {
	directions := make([]Direction, len(input))

	for i, s := range input {
		result := strings.Split(s, " ")
		n, _ := strconv.Atoi(result[1])
		directions[i] = Direction{result[0], n}
	}

	return directions
}

func Part1(input []Direction) int {
	depth := 0
	distance := 0

	for _, d := range input {
		switch d.direction {
		case "up":
			depth -= d.value
		case "down":
			depth += d.value
		case "forward":
			distance += d.value
		default:
			fmt.Printf("%s not supported", d.direction)

		}
	}

	return depth * distance
}

func Part2(input []Direction) int {
	depth := 0
	distance := 0
	aim := 0

	for _, d := range input {
		switch d.direction {
		case "up":
			aim -= d.value
		case "down":
			aim += d.value
		case "forward":
			distance += d.value
			depth += (aim * d.value)
		default:
			fmt.Printf("%s not supported", d.direction)
		}
	}

	return depth * distance
}

func main() {

	input := common.LoadFileString()
	parsed := ParseInput(input)

	fmt.Println(Part1(parsed))
	fmt.Println(Part2(parsed))
}
