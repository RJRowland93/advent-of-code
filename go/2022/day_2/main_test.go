package main

import "testing"

type TestCase struct {
	input    []string
	expected int
}

func TestPart1(t *testing.T) {
	cases := []TestCase{
		{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
			expected: 15,
		},
	}

	for _, c := range cases {
		result := Part1(c.input)

		if result != c.expected {
			t.Logf("expected %d but got: %d", c.expected, result)
			t.Fail()
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []TestCase{
		{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
			expected: 12,
		},
	}

	for _, c := range cases {
		result := Part2(c.input)

		if result != c.expected {
			t.Logf("expected %d but got: %d", c.expected, result)
			t.Fail()
		}
	}
}
