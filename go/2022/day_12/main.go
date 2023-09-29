package main

import (
	"advent-of-code/common"
	"errors"
	"fmt"
)

type Node struct {
	position [2]int
	height   byte
	parent   *Node
}

type Path map[Node]Node

type Queue struct {
	items []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(item *Node) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() *Node {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item
}

func (q *Queue) Len() int {
	return len(q.items)
}

func FindStart(graph []string, begin rune) (*Node, error) {
	for y, s := range graph {
		for x, c := range s {
			if c == begin {
				return &Node{[2]int{y, x}, 'a', nil}, nil
			}
		}
	}

	return &Node{}, errors.New("start not found")
}

func Neighbors(graph []string, node *Node) []*Node {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var neighbors []*Node

	for _, d := range directions {
		y := d[0] + node.position[0]
		x := d[1] + node.position[1]
		if y >= 0 && y < len(graph) && x >= 0 && x < len(graph[0]) {
			neighbors = append(neighbors, &Node{[2]int{y, x}, graph[y][x], node})
		}
	}

	return neighbors
}

func BFS(graph []string, begin rune, end rune) ([]*Node, error) {
	var paths []*Node
	visited := make(map[*Node]bool)
	start, err := FindStart(graph, begin)
	if err != nil {
		return paths, err
	}

	queue := NewQueue()
	queue.Push(start)
	for queue.Len() > 0 {
		current := queue.Pop()
		visited[current] = true
		if current.height == byte(end) {
			paths = append(paths, current)
		} else {
			neighbors := Neighbors(graph, current)
			for _, n := range neighbors {
				_, seen := visited[n]
				incline := n.height - current.height
				fmt.Println(current, n, seen, incline)
				if !seen && incline < 2 {
					queue.Push(n)
				}
			}
		}
	}

	return paths, nil
}

// func GetPath(Node) []Node {}

// func GetMinPath([]Node) int {}

// func Part1(input []string) int {
// }

func main() {
	input := common.LoadFileString()

	paths, err := BFS(input, 'S', 'E')
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range paths {
		fmt.Println(p)
	}

}
