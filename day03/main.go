package main

import (
	"fmt"
	"netsak/AdventOfCode2023/core"
)

func part1(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	return len(lines)
}

func part2(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	return len(lines)
}

func main() {
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(0, part1("part1-example.txt"))
	core.Assert(0, part1("part1-input.txt"))
	// part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(0, part2("part2-example.txt"))
	core.Assert(0, part2("part2-input.txt"))
}
