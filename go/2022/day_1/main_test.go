package main

import "testing"

type TestCase struct {
	input    []string
	expected int
}

func TestPart1(t *testing.T) {
	cases := []TestCase{
		{
			input: []string{"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000"},
			expected: 24000,
		},
	}

	for _, c := range cases {
		r := Part1(c.input)

		if r != c.expected {
			t.Logf("expected %d but got %d", c.expected, r)
			t.Fail()
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []TestCase{
		{
			input: []string{"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000"},
			expected: 45000,
		},
	}

	for _, c := range cases {
		r := Part2(c.input)

		if r != c.expected {
			t.Logf("expected %d but got %d", c.expected, r)
			t.Fail()
		}
	}
}
