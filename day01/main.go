package main

import (
	"fmt"
	"math"
	"netsak/AdventOfCode2023/core"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func part1(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	reDigit := regexp.MustCompile(`[0-9]`)
	sum := 0
	for _, line := range lines {
		found := map[int]int{}
		matches := reDigit.FindAllStringIndex(line, math.MaxInt64)
		for _, match := range matches {
			n, err := strconv.ParseInt(line[match[0]:match[1]], 10, 64)
			core.FailOnError(err, "ParseInt")
			found[match[0]] = int(n)
		}
		keys := maps.Keys(found)
		first := slices.Min(keys)
		last := slices.Max(keys)
		n := (10 * found[first]) + found[last]
		sum += n
	}
	fmt.Println("sum:", sum)
	return sum
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part2(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sum := 0
	for _, line := range lines {
		found := map[int]int{}
		for key := range numbers {
			if i := strings.Index(line, key); i != -1 {
				found[i] = numbers[key]
			}
			if i := strings.LastIndex(line, key); i != -1 {
				found[i] = numbers[key]
			}
		}
		keys := maps.Keys(found)
		first := slices.Min(keys)
		last := slices.Max(keys)
		n := (10 * found[first]) + found[last]
		sum += n
	}
	fmt.Println("sum:", sum)
	return sum
}

func main() {
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(142, part1("part1-example.txt"))
	core.Assert(55386, part1("part1-input.txt"))
	// part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(281, part2("part2-example.txt"))
	core.Assert(54824, part2("part1-input.txt"))
}
