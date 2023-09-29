package main

import "testing"

type TestCase struct {
	input    []int
	expected int
}

func TestPart1(t *testing.T) {

	cases := []TestCase{
		{input: []int{1, 2, 3}, expected: 2},
		{input: []int{1, 2, 2}, expected: 1},
		{input: []int{1, 2, 3, 4}, expected: 3},
	}

	for _, c := range cases {
		r := Part1(c.input)

		if r != c.expected {
			t.Logf("expected: %d but got: %d", c.expected, r)
			t.Fail()
		}
	}
}

func TestPart2(t *testing.T) {

	cases := []TestCase{
		// {input: []int{}, expected: 2},
		{input: []int{1, 2, 3}, expected: 2},
		{input: []int{1, 2, 2}, expected: 1},
		{input: []int{1, 2, 3, 4}, expected: 3},
	}

	for _, c := range cases {
		r := Part2(c.input)

		if r != c.expected {
			t.Logf("expected: %d but got: %d", c.expected, r)
			t.Fail()
		}
	}
}
