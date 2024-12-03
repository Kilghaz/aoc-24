package main

import (
	"aoc24/day_1"
	"aoc24/day_2"
	"aoc24/day_3"
	"aoc24/io"
	"os"
	"strconv"
)

var advents = map[int][]func(input string){
	1: {
		day_1.Part1,
		day_1.Part2,
	},
	2: {
		day_2.Part1,
		day_2.Part2,
	},
	3: {
		day_3.Part1,
		day_3.Part2,
	},
}

func main() {
	day := toInt(os.Args[1])
	part := toInt(os.Args[2]) - 1
	input := io.LoadInput(day)
	advents[day][part](input)
}

func toInt(value string) int {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(parsed)
}
