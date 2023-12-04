package main

import (
	"fmt"
	"math"
	"netsak/AdventOfCode2023/core"
	"strings"

	"golang.org/x/exp/maps"
)

func matchingNumbers(row string) int {
	fields := strings.Fields(row)
	seen := map[string]int{}
	count := 0
	for _, key := range fields {
		seen[key] += 1
		if seen[key] > 1 {
			count++
		}
	}
	return count
}

func part1(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sum := 0
	for _, row := range lines {
		count := matchingNumbers(row)
		sum += int(math.Pow(2, float64(count-1)))
	}
	fmt.Println("How many points are they worth in total?", sum)
	return sum
}

func nextCards(start, count, max int) []int {
	ret := []int{}
	for i := start + 1; i < min(start+count+1, max+1); i++ {
		ret = append(ret, i)
	}
	return ret
}

func processCards(ids []int, cardMatches map[int]int) int {
	total := len(ids)
	for _, i := range ids {
		next := nextCards(i, cardMatches[i], len(cardMatches))
		total += processCards(next, cardMatches)
	}
	return total
}

func part2(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	cardMatches := map[int]int{}
	for i, row := range lines {
		cardMatches[i+1] = matchingNumbers(row)
	}
	sum := processCards(maps.Keys(cardMatches), cardMatches)
	fmt.Println("How many total scratchcards do you end up with?", sum)
	return sum
}

func main() {
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(13, part1("part1-example.txt"))
	core.Assert(21919, part1("part1-input.txt"))
	// part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(30, part2("part1-example.txt"))
	core.Assert(9881048, part2("part1-input.txt"))
}
