package main

import (
	"strings"
	"testing"
)

const raw = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	input := strings.Split(raw, "\n")

	trees := ParseInput(input)
	expected := 21
	result := Part1(trees)

	if result != expected {
		t.Logf("expected %d but got: %d", expected, result)
		t.Fail()
	}
}
