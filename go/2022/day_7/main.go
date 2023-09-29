package main

import (
	"advent-of-code/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(input []string) [][]string {
	result := make([][]string, len(input))
	for i, s := range input {
		split := strings.Split(s, " ")
		result[i] = split
	}
	return result
}

type File struct {
	name  string
	size  int
	isDir bool
	files []*File
}

func NewFile(name string, isDir bool, size int) *File {
	files := []*File{}
	return &File{
		name:  name,
		isDir: isDir,
		size:  size,
		files: files,
	}
}

func NewFileTree(instructions [][]string) *File {
	// var cwd []*File
	// var head *File
	head := NewFile("/", true, 0)
	cwd := []*File{head}
	for _, ins := range instructions {
		// fmt.Println(ins)
		switch ins[0] {
		case "$":
			if ins[1] == "cd" {
				// fmt.Printf("%#v\n", cwd)
				if ins[2] == ".." {
					cwd = cwd[:len(cwd)-1]
				} else {
					// fmt.Println("pushing dir")
					for _, f := range head.files {
						// fmt.Printf("iterate file %s\n", f.name)
						if f.name == ins[2] {
							cwd = append(cwd, f)
							break
						}
					}
				}
				head = cwd[len(cwd)-1]
				// fmt.Printf("%#v\n", cwd)
			}
		case "dir":
			head.files = append(head.files, NewFile(ins[1], true, 0))
		default:
			size, _ := strconv.Atoi(ins[0])
			head.files = append(head.files, NewFile(ins[1], false, size))
		}
		// fmt.Println(head)
	}
	// PrintDirs(cwd[0])
	return cwd[0]
}

type DirectorySize struct {
	name string
	size int
}

func DirectorySizes(fs *File) []DirectorySize {
	var sizes []DirectorySize

	var dfs func(*File) int
	dfs = func(f *File) int {
		if !f.isDir {
			return f.size
		}
		for _, c := range f.files {
			// fmt.Println(c.name, c.isDir)
			f.size += dfs(c)
		}

		// fmt.Println(f.name, f.isDir, f.size)
		sizes = append(sizes, DirectorySize{f.name, f.size})

		return f.size
	}

	dfs(fs)

	return sizes
}

func Part1(dirs []DirectorySize) int {
	sum := 0
	for _, d := range dirs {
		if d.size <= 100000 {
			fmt.Println(d.name, d.size)
			sum += d.size
		}
	}
	return sum
}

func Part2(dirs []DirectorySize) int {

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size > dirs[j].size
	})
	disk := 70000000
	required := 30000000
	remaining := disk - dirs[0].size
	free := required - remaining

	for i, d := range dirs {
		fmt.Println(d.name, d.size)
		if d.size < free {
			fmt.Println(d.name, d.size, free)
			return dirs[i-1].size
		}
	}
	return 0
}

func main() {
	input := common.LoadFileString()
	instructions := ParseInput(input)
	fs := NewFileTree(instructions)
	ds := DirectorySizes(fs)
	// r := Part1(ds)
	// fmt.Println(ds)
	r := Part2(ds)
	fmt.Println(r)
}
