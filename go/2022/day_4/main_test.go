package main

import (
	"strings"
	"testing"
)

const raw = `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

func TestPart1(t *testing.T) {
	input := strings.Split(raw, "\n")
	expected := 2
	result := Part1(ParseInput(input))

	if result != expected {
		t.Logf("expected %d but got: %d", expected, result)
		t.Fail()
	}
}
