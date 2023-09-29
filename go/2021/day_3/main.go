package main

import (
	"advent-of-code/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseInput(input []string) [][]int {
	result := make([][]int, len(input))

	for i, s := range input {
		ds := strings.Split(s, "")
		result[i] = make([]int, len(ds))

		for j, d := range ds {
			n, _ := strconv.Atoi(d)
			result[i][j] = n
		}
	}

	return result
}

func CountColumns(input [][]int) []int {
	counts := make([]int, len(input[0]))

	for _, s := range input {
		for i, n := range s {
			counts[i] += n
		}
	}

	return counts
}

func ToOnesAndZeroes(counts []int, half int) []int {
	bits := make([]int, len(counts))
	// half := int(len(input) / 2)

	for j, b := range counts {
		if b > half {
			bits[j] = 1
		} else {
			bits[j] = 0
		}
	}
	return bits
}

func Xor(bits []int) []int {
	xor := make([]int, len(bits))
	for i, b := range bits {
		xor[i] = b ^ 1
	}
	return xor
}

func BitsToNumber(bits []int) float64 {
	number := 0.0
	for i, k := 0, len(bits)-1; k >= 0; i, k = i+1, k-1 {
		something := math.Pow(2, float64(k))
		d := bits[i]
		number += float64(d) * something
	}
	return number
}

func Part1(input [][]int) float64 {

	counts := CountColumns(input)
	bits := ToOnesAndZeroes(counts, len(input)/2)

	gamma := BitsToNumber(bits)
	epsilon := BitsToNumber(Xor(bits))

	fmt.Println(bits, epsilon, gamma)

	return gamma * epsilon
}

func main() {
	input := common.LoadFileString()
	parsed := ParseInput(input)
	fmt.Println(Part1(parsed))
}
