package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	cases := []struct {
		input    []string
		expected [][]int
	}{
		{
			input: []string{
				"00100",
				"11110",
			},
			expected: [][]int{
				{0, 0, 1, 0, 0},
				{1, 1, 1, 1, 0},
			},
		},
	}

	for _, c := range cases {
		result := ParseInput(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Logf("got %v, expected %v", result, c.expected)
			t.Fail()
		}
	}
}

func TestCountColumns(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []int
	}{
		{
			input: [][]int{
				{1, 0, 1, 0, 0},
				{1, 1, 1, 1, 0},
			},
			expected: []int{
				2, 1, 2, 1, 0,
			},
		},
	}

	for _, c := range cases {
		result := CountColumns(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Logf("got %v, expected %v", result, c.expected)
			t.Fail()
		}
	}
}

func TestToOnesAndZeroes(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				10, 11, 12, 1, 0,
			},
			expected: []int{
				1, 1, 1, 0, 0,
			},
		},
	}

	for _, c := range cases {
		result := ToOnesAndZeroes(c.input, len(c.input)/2)
		if !reflect.DeepEqual(result, c.expected) {
			t.Logf("got %v, expected %v", result, c.expected)
			t.Fail()
		}
	}
}

func TestBitsToNumber(t *testing.T) {
	cases := []struct {
		input    []int
		expected float64
	}{
		{
			input: []int{
				1, 1, 1, 0, 0,
			},
			expected: 28.0,
		},
		{
			input: []int{
				1, 0, 1, 1, 0,
			},
			expected: 22.0,
		},
		{
			input: []int{
				0, 1, 0, 0, 1,
			},
			expected: 9.0,
		},
	}

	for _, c := range cases {
		result := BitsToNumber(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Logf("got %v, expected %v", result, c.expected)
			t.Fail()
		}
	}
}

func TestXor(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				1, 1, 1, 0, 0,
				0, 0, 0, 0, 0,
				1, 1, 1, 1, 1,
			},
			expected: []int{
				0, 0, 0, 1, 1,
				1, 1, 1, 1, 1,
				0, 0, 0, 0, 0,
			},
		},
	}

	for _, c := range cases {
		result := Xor(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Logf("got %v, expected %v", result, c.expected)
			t.Fail()
		}
	}
}

func TestPart1(t *testing.T) {
	cases := []struct {
		input    []string
		expected float64
	}{
		{
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expected: 198.0, // 10110 -> 22 * 01001 -> 9
		},
	}

	for _, c := range cases {
		result := Part1(ParseInput(c.input))
		if result != c.expected {
			t.Logf("got %.2f, expected %.2f", result, c.expected)
			t.Fail()
		}
	}
}
