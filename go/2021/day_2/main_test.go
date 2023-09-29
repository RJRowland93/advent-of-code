package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input    []Direction
	expected int
}

func TestParseInput(t *testing.T) {
	cases := []struct {
		input    []string
		expected []Direction
	}{
		{input: []string{"up 1", "down 1", "forward 1"}, expected: []Direction{{"up", 1}, {"down", 1}, {"forward", 1}}},
	}

	for _, c := range cases {
		r := ParseInput(c.input)

		if !reflect.DeepEqual(r, c.expected) {
			t.Logf("expected: %v but got: %v", c.expected, r)
			t.Fail()
		}
	}
}

func TestPart1(t *testing.T) {
	cases := []TestCase{
		{input: []Direction{{"forward", 5}, {"down", 2}, {"up", 1}}, expected: 5},
	}

	for _, c := range cases {
		r := Part1(c.input)

		if r != c.expected {
			t.Logf("expected: %d but got: %d", c.expected, r)
			t.Fail()
		}
	}
}

// forward 5 adds 5 to your horizontal position, a total of 5. Because your aim is 0, your depth does not change.
// down 5 adds 5 to your aim, resulting in a value of 5.
// forward 8 adds 8 to your horizontal position, a total of 13. Because your aim is 5, your depth increases by 8*5=40.
// up 3 decreases your aim by 3, resulting in a value of 2.
// down 8 adds 8 to your aim, resulting in a value of 10.
// forward 2 adds 2 to your horizontal position, a total of 15. Because your aim is 10, your depth increases by 2*10=20 to a total of 60.
// After following these new instructions, you would have a horizontal position of 15 and a depth of 60. (Multiplying these produces 900.)

func TestPart2(t *testing.T) {
	cases := []TestCase{
		{input: []Direction{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}, expected: 900},
	}

	for _, c := range cases {
		r := Part2(c.input)

		if r != c.expected {
			t.Logf("expected: %d but got: %d", c.expected, r)
			t.Fail()
		}
	}
}
