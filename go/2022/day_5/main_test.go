package main

import (
	"strings"
	"testing"
)

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// CMZ

const raw = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1(t *testing.T) {

	start := [][]rune{
		{'D'},
		{'N', 'C'},
		{'N', 'M', 'P'},
	}
	input := strings.Split(raw, "\n")
	instructions := ParseInput(input)
	result := Part1(start, instructions)
	expected := "CMZ"

	if result != expected {
		t.Logf("expected %s but got: %s", expected, result)
		t.Fail()
	}
}
