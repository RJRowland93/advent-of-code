package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  []string
	result int
}

func TestSplitString(t *testing.T) {
	cases := []struct {
		input  string
		result [2]string
	}{
		{input: "", result: [2]string{}},
		{input: "a", result: [2]string{"", "a"}},
		{input: "ab", result: [2]string{"a", "b"}},
		{input: "abc", result: [2]string{"a", "bc"}},
		{input: "abcd", result: [2]string{"ab", "cd"}},
	}

	for _, c := range cases {
		r := SplitString(c.input)

		if !reflect.DeepEqual(r, c.result) {
			t.Logf("expected %v  but got: %v", c.result, r)
			t.Fail()
		}
	}
}

func TestStringIntersection(t *testing.T) {
	cases := []struct {
		input  [2]string
		result rune
	}{
		{input: [2]string{"ab", "bc"}, result: 'b'},
		{input: [2]string{"", "ab"}, result: 0},
	}

	for _, c := range cases {
		r := StringIntersection(c.input)

		if !reflect.DeepEqual(r, c.result) {
			t.Logf("expected %v  but got: %v", c.result, r)
			t.Fail()
		}
	}
}

func TestPriority(t *testing.T) {
	cases := []struct {
		input  rune
		result int
	}{
		{input: 'a', result: 1},
		{input: 'z', result: 26},
		{input: 'A', result: 27},
	}

	for _, c := range cases {
		r := Priority(c.input)

		if r != c.result {
			t.Logf("expected %v  but got: %v", c.result, r)
			t.Fail()
		}
	}
}

func TestGroupBy(t *testing.T) {
	cases := []struct {
		input  []string
		result [][]string
	}{
		{input: []string{"hello", "world", "goodbye", "universe"},
			result: [][]string{{"hello", "world"}, {"goodbye", "universe"}}},
	}

	for _, c := range cases {
		r := GroupBy(2, c.input)
		if !reflect.DeepEqual(r, c.result) {
			t.Logf("expected %v but got: %v", c.result, r)
			t.Fail()
		}
	}
}

func TestStringIntersectAll(t *testing.T) {
	cases := []struct {
		input  []string
		result rune
	}{
		{
			input:  []string{"abc", "bde", "bfg"},
			result: 'b',
		},
	}

	for _, c := range cases {
		r := StringIntersectionAll(c.input)
		if r != c.result {
			t.Logf("expected %v but got: %v", c.result, r)
			t.Fail()
		}
	}
}
func TestPart1(t *testing.T) {

	cases := []TestCase{
		{
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw"},
			result: 157,
		},
	}

	for _, c := range cases {
		r := Part1(c.input)

		if r != c.result {
			t.Logf("expected %v  but got: %v", c.result, r)
			t.Fail()
		}
	}
}

func TestPart2(t *testing.T) {

	cases := []TestCase{
		{
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			result: 18},
		{
			input: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			result: 52,
		},
	}

	for _, c := range cases {
		r := Part2(c.input)

		if r != c.result {
			t.Logf("expected %v  but got: %v", c.result, r)
			t.Fail()
		}
	}
}
