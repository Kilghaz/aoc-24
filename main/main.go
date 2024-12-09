package main

import (
	"aoc24/day_1"
	"aoc24/day_2"
	"aoc24/day_3"
	"aoc24/day_4"
	"aoc24/day_5"
	"aoc24/day_6"
	"aoc24/day_7"
	"aoc24/day_8"
	"aoc24/io"
	"os"
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
	4: {
		day_4.Part1,
		day_4.Part2,
	},
	5: {
		day_5.Part1,
		day_5.Part2,
	},
	6: {
		day_6.Part1,
		day_6.Part2,
	},
	7: {
		day_7.Part1,
		day_7.Part2,
	},
	8: {
		day_8.Part1,
		day_8.Part2,
	},
}

func main() {
	day := io.ParseInt(os.Args[1])
	part := io.ParseInt(os.Args[2]) - 1
	var input string
	if len(os.Args) == 3 || os.Args[3] != "test" {
		input = io.LoadInput(day)
	} else {
		input = io.LoadTestInput(day)
	}
	advents[day][part](input)
}
