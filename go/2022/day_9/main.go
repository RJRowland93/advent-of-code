package main

import (
	"advent-of-code/common"
	"fmt"
	"strconv"
	"strings"
)

type Move struct {
	dir      string
	distance int
}

type Coord struct {
	x int
	y int
}

func Head(c Coord, dir string) Coord {
	switch dir {
	case "L":
		c.x -= 1
	case "R":
		c.x += 1
	case "U":
		c.y += 1
	case "D":
		c.y -= 1
	}
	return c
}

func ParseInput(input []string) []Move {
	moves := make([]Move, len(input))
	for i, in := range input {
		split := strings.Split(in, " ")
		n, _ := strconv.Atoi(split[1])
		moves[i] = Move{split[0], n}
	}
	return moves
}

func Distance(head, tail Coord) Coord {
	return Coord{head.x - tail.x, head.y - tail.y}
}
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Follow(tail, d Coord) Coord {
	x := 0
	y := 0

	if d.x == 2 {
		if Abs(d.y) == 1 {
			y = d.y
		}
		x = 1
	}
	if d.y == 2 {
		if Abs(d.x) == 1 {
			x = d.x
		}
		y = 1
	}
	if d.x == -2 {
		if Abs(d.y) == 1 {
			y = d.y
		}
		x = -1
	}
	if d.y == -2 {
		if Abs(d.x) == 1 {
			x = d.x
		}
		y = -1
	}

	return Coord{tail.x + x, tail.y + y}
}

func Part1(moves []Move) int {
	head := Coord{0, 0}
	tail := Coord{0, 0}
	record := make(map[[2]int]bool)

	for _, m := range moves {
		for i := 0; i < m.distance; i++ {
			head = Head(head, m.dir)
			d := Distance(head, tail)
			tail = Follow(tail, d)
			record[[2]int{tail.x, tail.y}] = true

		}
	}

	return len(record)
}

func Part2(moves []Move) int {
	rope := [10]Coord{
		{0, 0}, // head
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0}, // tail
	}
	record := make(map[[2]int]bool)

	for _, m := range moves {
		for i := 0; i < m.distance; i++ {
			rope[0] = Head(rope[0], m.dir)
			for j := 1; j < len(rope); j++ {
				d := Distance(rope[j-1], rope[j])
				rope[j] = Follow(rope[j], d)
			}
			record[[2]int{rope[9].x, rope[9].y}] = true

		}
	}
	fmt.Println(rope)

	return len(record)
}

func main() {
	input := common.LoadFileString()
	moves := ParseInput(input)
	// tail := Part1(moves) // 6037
	// fmt.Println(tail, 6037)
	tail := Part2(moves) // 2485
	fmt.Println(tail)
}
