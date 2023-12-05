package main

import (
	"fmt"
	"math"
	"netsak/AdventOfCode2023/core"
	"strconv"
	"strings"
)

const (
	SeedToSoil            = "seed-to-soil"
	SoilToFertilizer      = "soil-to-fertilizer"
	FertilizerToWater     = "fertilizer-to-water"
	WaterToLight          = "water-to-light"
	LightToTemperature    = "light-to-temperature"
	TemperatureToHumidity = "temperature-to-humidity"
	HumidityToLocation    = "humidity-to-location"
)

type Almanac struct {
	Seeds []int64
	Maps  map[string][]DSL
}

type DSL struct {
	Dst    int64
	Src    int64
	Length int64
}

func (a *Almanac) AddToMap(name string, fields []string) {
	dstStart, err := strconv.ParseInt(fields[0], 10, 64)
	core.FailOnError(err, "parse dstStart")
	srcStart, err := strconv.ParseInt(fields[1], 10, 64)
	core.FailOnError(err, "parse srcStart")
	length, err := strconv.ParseInt(fields[2], 10, 64)
	core.FailOnError(err, "parse length")
	a.Maps[name] = append(a.Maps[name], DSL{
		Dst:    dstStart,
		Src:    srcStart,
		Length: length,
	})
}

func (a *Almanac) Map(name string, idx int64) int64 {
	for _, s := range a.Maps[name] {
		if idx >= s.Src && idx < s.Src+s.Length {
			return s.Dst + (idx - s.Src)
		}
	}
	return idx
}

func (a *Almanac) Location(seed int64) int64 {
	soil := a.Map(SeedToSoil, seed)
	fertilizer := a.Map(SoilToFertilizer, soil)
	water := a.Map(FertilizerToWater, fertilizer)
	light := a.Map(WaterToLight, water)
	temperature := a.Map(LightToTemperature, light)
	humidity := a.Map(TemperatureToHumidity, temperature)
	location := a.Map(HumidityToLocation, humidity)
	return location
}

func newAlmanac(rows []string) Almanac {
	a := Almanac{
		Seeds: []int64{},
		Maps:  map[string][]DSL{},
	}
	var currentMap string
	for _, row := range rows {
		fields := strings.Fields(row)
		if fields[0] == "seeds:" {
			for _, seed := range fields[1:] {
				number, err := strconv.ParseInt(seed, 10, 64)
				core.FailOnError(err, "parse seed")
				a.Seeds = append(a.Seeds, number)
			}
			continue
		}
		if fields[1] == "map:" {
			currentMap = fields[0]
			continue
		}
		a.AddToMap(currentMap, fields)
	}
	return a
}

func part0(filename string) Almanac {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	return newAlmanac(lines)
}

func part1(filename string) int64 {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	almanac := newAlmanac(lines)
	lowestLocation := int64(math.MaxInt64)
	for _, seed := range almanac.Seeds {
		location := almanac.Location(seed)
		lowestLocation = min(lowestLocation, location)
	}
	fmt.Println("What is the lowest location number that corresponds to any of the initial seed numbers?", lowestLocation)
	return lowestLocation
}

func part2(filename string) int64 {
	lines, err := core.ReadLines(filename)
	core.FailOnError(err, "ReadLines")
	almanac := newAlmanac(lines)
	start := int64(0)
	lowestLocation := int64(math.MaxInt64)
	for i, n := range almanac.Seeds {
		if i%2 == 0 {
			start = n
			continue
		}
		for seed := start; seed < start+n; seed++ {
			location := almanac.Location(seed)
			lowestLocation = min(lowestLocation, location)
		}
	}
	fmt.Println("What is the lowest location number that corresponds to any of the initial seed numbers?", lowestLocation)
	return lowestLocation
}

func main() {
	// part 0
	fmt.Println("---------- Part 0 ----------")
	a := part0("part1-example.txt")
	core.Assert(52, a.Map(SeedToSoil, 50))
	core.Assert(57, a.Map(SeedToSoil, 55))
	core.Assert(81, a.Map(SeedToSoil, 79))
	core.Assert(14, a.Map(SeedToSoil, 14))
	core.Assert(13, a.Map(SeedToSoil, 13))
	core.Assert(1, a.Map(SeedToSoil, 1))
	// part 1
	fmt.Println("---------- Part 1 ----------")
	core.Assert(35, part1("part1-example.txt"))
	core.Assert(331445006, part1("part1-input.txt"))
	// part2
	fmt.Println("---------- Part 2 ----------")
	core.Assert(46, part2("part1-example.txt"))
	core.Assert(6472060, part2("part1-input.txt"))
}
