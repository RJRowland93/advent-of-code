package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	id          int
	items       []int
	op          func(int) int
	div         int
	test        func(int) int
	inspections int
}

func Part1() int {
	monkeys := []*Monkey{
		{
			id:    0,
			items: []int{59, 74, 65, 86},
			op:    func(x int) int { return x * 19 },
			test: func(x int) int {
				if x%7 == 0 {
					return 6
				}
				return 2
			},
			inspections: 0,
		},
		{
			id:    1,
			items: []int{62, 84, 72, 91, 68, 78, 51},
			op:    func(x int) int { return x + 1 },
			test: func(x int) int {
				if x%2 == 0 {
					return 2
				}
				return 0
			},
			inspections: 0,
		},
		{
			id:    2,
			items: []int{78, 84, 96},
			op:    func(x int) int { return x + 8 },
			test: func(x int) int {
				if x%19 == 0 {
					return 6
				}
				return 5
			},
			inspections: 0,
		},
		{
			id:    3,
			items: []int{97, 86},
			op:    func(x int) int { return x * x },
			test: func(x int) int {
				if x%3 == 0 {
					return 1
				}
				return 0
			},
			inspections: 0,
		},
		{
			id:    4,
			items: []int{50},
			op:    func(x int) int { return x + 6 },
			test: func(x int) int {
				if x%13 == 0 {
					return 3
				}
				return 1
			},
			inspections: 0,
		},
		{
			id:    5,
			items: []int{73, 65, 69, 65, 51},
			op:    func(x int) int { return x * 17 },
			test: func(x int) int {
				if x%11 == 0 {
					return 4
				}
				return 7
			},
			inspections: 0,
		},
		{
			id:    6,
			items: []int{69, 82, 97, 93, 82, 84, 58, 63},
			op:    func(x int) int { return x + 5 },
			test: func(x int) int {
				if x%5 == 0 {
					return 5
				}
				return 7
			},
			inspections: 0,
		},
		{
			id:    7,
			items: []int{81, 78, 82, 76, 79, 80},
			op:    func(x int) int { return x + 3 },
			test: func(x int) int {
				if x%17 == 0 {
					return 3
				}
				return 4
			},
			inspections: 0,
		},
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for _, i := range m.items {
				stress := m.op(i)
				m.inspections++
				relief := stress / 3
				next := m.test(relief)
				monkeys[next].items = append(monkeys[next].items, relief)
			}
			m.items = []int{} // reset monkey's items
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	for _, f := range monkeys {
		fmt.Println(f)
	}

	return monkeys[0].inspections * monkeys[1].inspections
}

func Part2() int {
	monkeys := []*Monkey{
		{
			id:    0,
			items: []int{59, 74, 65, 86},
			op:    func(x int) int { return x * 19 },
			div:   7,
			test: func(x int) int {
				if x%7 == 0 {
					return 6
				}
				return 2
			},
			inspections: 0,
		},
		{
			id:    1,
			items: []int{62, 84, 72, 91, 68, 78, 51},
			op:    func(x int) int { return x + 1 },
			div:   2,
			test: func(x int) int {
				if x%2 == 0 {
					return 2
				}
				return 0
			},
			inspections: 0,
		},
		{
			id:    2,
			items: []int{78, 84, 96},
			op:    func(x int) int { return x + 8 },
			div:   19,
			test: func(x int) int {
				if x%19 == 0 {
					return 6
				}
				return 5
			},
			inspections: 0,
		},
		{
			id:    3,
			items: []int{97, 86},
			op:    func(x int) int { return x * x },
			div:   3,
			test: func(x int) int {
				if x%3 == 0 {
					return 1
				}
				return 0
			},
			inspections: 0,
		},
		{
			id:    4,
			items: []int{50},
			op:    func(x int) int { return x + 6 },
			div:   13,
			test: func(x int) int {
				if x%13 == 0 {
					return 3
				}
				return 1
			},
			inspections: 0,
		},
		{
			id:    5,
			items: []int{73, 65, 69, 65, 51},
			op:    func(x int) int { return x * 17 },
			div:   11,
			test: func(x int) int {
				if x%11 == 0 {
					return 4
				}
				return 7
			},
			inspections: 0,
		},
		{
			id:    6,
			items: []int{69, 82, 97, 93, 82, 84, 58, 63},
			op:    func(x int) int { return x + 5 },
			div:   5,
			test: func(x int) int {
				if x%5 == 0 {
					return 5
				}
				return 7
			},
			inspections: 0,
		},
		{
			id:    7,
			items: []int{81, 78, 82, 76, 79, 80},
			op:    func(x int) int { return x + 3 },
			div:   17,
			test: func(x int) int {
				if x%17 == 0 {
					return 3
				}
				return 4
			},
			inspections: 0,
		},
	}

	gcd := 1
	for _, m := range monkeys {
		gcd *= m.div
	}
	fmt.Println("gcd: ", gcd)

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for _, i := range m.items {
				stress := m.op(i)
				m.inspections++
				relief := stress % gcd
				next := m.test(relief)
				monkeys[next].items = append(monkeys[next].items, relief)
			}
			m.items = []int{} // reset monkey's items
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	for _, f := range monkeys {
		fmt.Println(f)
	}

	return monkeys[0].inspections * monkeys[1].inspections
}

func main() {
	// result := Part1()
	result := Part2()
	fmt.Println(result)
}
