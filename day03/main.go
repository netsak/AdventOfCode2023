package main

import (
	"fmt"
	"netsak/AdventOfCode2023/core"
	"regexp"
	"strconv"
)

func part1(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	rows := core.ExtendRows(lines, ".", 1)
	reNumber := regexp.MustCompile(`\d+`)
	reSymbol := regexp.MustCompile(`[^.0-9]+`)
	sum := 0
	for y, row := range rows {
		for _, match := range reNumber.FindAllStringSubmatchIndex(row, -1) {
			x1, x2 := match[0], match[1]
			number, _ := strconv.Atoi(row[x1:x2])
			neighbors := []string{
				rows[y-1][x1-1 : x2+1],
				rows[y+0][x1-1 : x2+1],
				rows[y+1][x1-1 : x2+1],
			}
			for _, row := range neighbors {
				if reSymbol.MatchString(row) {
					sum += number
				}
			}
		}
	}
	fmt.Println("What is the sum of all of the part numbers in the engine schematic?", sum)
	return sum
}

func part2(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	rows := core.ExtendRows(lines, ".", 1)
	reGear := regexp.MustCompile(`\*`)
	reNumber := regexp.MustCompile(`\d+`)
	sum := 0
	for y, row := range rows {
		for _, match := range reGear.FindAllStringSubmatchIndex(row, -1) {
			x := match[0]
			neighbors := []string{
				rows[y-1],
				rows[y+0],
				rows[y+1],
			}
			numbers := []int{}
			for _, row := range neighbors {
				for _, match := range reNumber.FindAllStringSubmatchIndex(row, -1) {
					xStart, xEnd := match[0]-1, match[1]
					if xStart <= x && xEnd >= x {
						strNumber := row[match[0]:match[1]]
						number, _ := strconv.Atoi(strNumber)
						numbers = append(numbers, number)
					}
				}
			}
			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			}
		}
	}
	fmt.Println("What is the sum of all of the gear ratios in your engine schematic?", sum)
	return sum
}

func main() {
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(4361, part1("part1-example.txt"))
	core.Assert(540131, part1("part1-input.txt"))
	// // part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(467835, part2("part1-example.txt"))
	core.Assert(86879020, part2("part1-input.txt"))
}
