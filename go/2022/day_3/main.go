package main

import (
	"advent-of-code/common"
	"fmt"
)

const priorities = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GroupBy(n int, s []string) [][]string {
	var result [][]string
	for i := 0; i < len(s)/n; i++ {
		result = append(result, s[i*n:i*n+n])
	}
	return result
}

func SplitString(s string) [2]string {
	mid := len(s) / 2

	return [2]string{s[:mid], s[mid:]}
}

func StringIntersection(s [2]string) rune {
	s1 := s[0]
	s2 := s[1]
	set1 := make(map[rune]struct{})
	// var intersection strings.Builder

	for _, c := range s1 {
		set1[c] = struct{}{}
	}

	for _, c := range s2 {
		_, exists := set1[c]
		if exists {
			return c
			// intersection.WriteString(string(c))
		}
	}

	return 0
}

func StringIntersectionAll(s []string) rune {
	s1 := s[0]
	s2 := s[1]
	s3 := s[2]
	set1 := make(map[rune]struct{})
	set2 := make(map[rune]struct{})
	// var intersection strings.Builder

	for _, c := range s1 {
		set1[c] = struct{}{}
	}

	for _, c := range s2 {
		set2[c] = struct{}{}
	}

	for _, c := range s3 {
		_, exists1 := set1[c]
		_, exists2 := set2[c]

		if exists1 && exists2 {
			return c
			// intersection.WriteString(string(c))
		}
	}

	return 0
}

func Priority(r rune) int {
	// return int(r) - int('a') + 1
	for i, c := range priorities {
		if r == c {
			return i
		}
	}
	return 0
}

func Part1(input []string) int {
	sum := 0
	// 1. split input in half
	// 2. get intersection of halves
	// 3. map union to priority
	// 4. get sum of priorities

	for _, s := range input {
		sum += Priority(StringIntersection(SplitString(s)))
	}

	return sum
}

func Part2(input []string) int {
	// 1. group input by 3
	// 2. get intersection triplets
	// 3. map union to priority
	// 4. get sum of priorities
	sum := 0

	groups := GroupBy(3, input)

	for _, g := range groups {
		sum += Priority(StringIntersectionAll(g))
	}

	return sum
}

func main() {
	input := common.LoadFileString()
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}
