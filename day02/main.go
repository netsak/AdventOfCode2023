package main

import (
	"fmt"
	"netsak/AdventOfCode2023/core"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sumOfGameIDs := 0
	re := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
nextGame:
	for gameNo, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			count, err := strconv.Atoi(match[1])
			core.FailOnError(err, "strconv")
			if count > cubes[match[2]] {
				continue nextGame
			}
		}
		sumOfGameIDs += gameNo + 1
	}
	fmt.Println("What is the sum of the IDs of those games?", sumOfGameIDs)
	return sumOfGameIDs
}

func part2(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sumOfPowers := 0
	re := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
	for _, line := range lines {
		game := map[string]int{}
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			count, err := strconv.Atoi(match[1])
			core.FailOnError(err, "strconv")
			game[match[2]] = max(game[match[2]], count)
		}
		sumOfPowers += (game["red"] * game["blue"] * game["green"])
	}
	fmt.Println("What is the sum of the power of these sets?", sumOfPowers)
	return sumOfPowers
}

func part1scanner(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sumOfGameIDs := 0
	var s scanner.Scanner
nextGame:
	for gameNo, line := range lines {
		s.Init(strings.NewReader(line))
		var count int
		game := map[string]int{}
		for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
			text := s.TokenText()
			switch token {
			case scanner.Ident:
				game[text] = max(game[text], count)
			case scanner.Int:
				count, err = strconv.Atoi(text)
				core.FailOnError(err, "scan integer")
			}
		}
		for color, maxNumber := range cubes {
			if game[color] > maxNumber {
				continue nextGame
			}
		}
		sumOfGameIDs += gameNo + 1
	}
	fmt.Println("What is the sum of the IDs of those games?", sumOfGameIDs)
	return sumOfGameIDs
}

func part2scanner(filename string) int {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	sumOfPowers := 0
	var s scanner.Scanner
	for _, line := range lines {
		s.Init(strings.NewReader(line))
		var count int
		game := map[string]int{}
		for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
			text := s.TokenText()
			switch token {
			case scanner.Ident:
				game[text] = max(game[text], count)
			case scanner.Int:
				count, err = strconv.Atoi(text)
				core.FailOnError(err, "scan integer")
			}
		}
		sumOfPowers += (game["red"] * game["blue"] * game["green"])
	}
	fmt.Println("What is the sum of the power of these sets?", sumOfPowers)
	return sumOfPowers
}

func main() {
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(8, part1("part1-example.txt"))
	core.Assert(8, part1scanner("part1-example.txt"))
	core.Assert(2447, part1("part1-input.txt"))
	core.Assert(2447, part1scanner("part1-input.txt"))
	// part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(2286, part2("part1-example.txt"))
	core.Assert(2286, part2scanner("part1-example.txt"))
	core.Assert(56322, part2("part1-input.txt"))
	core.Assert(56322, part2scanner("part1-input.txt"))
}
