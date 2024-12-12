package main

import (
	"aoc24/day_1"
	"aoc24/day_10"
	"aoc24/day_11"
	"aoc24/day_12"
	"aoc24/day_2"
	"aoc24/day_3"
	"aoc24/day_4"
	"aoc24/day_5"
	"aoc24/day_6"
	"aoc24/day_7"
	"aoc24/day_8"
	"aoc24/day_9"
	"aoc24/io"
	"aoc24/parser"
	"fmt"
	"os"
	"time"
)

var advents = map[int][]func(input string){
	1:  {day_1.Part1, day_1.Part2},
	2:  {day_2.Part1, day_2.Part2},
	3:  {day_3.Part1, day_3.Part2},
	4:  {day_4.Part1, day_4.Part2},
	5:  {day_5.Part1, day_5.Part2},
	6:  {day_6.Part1, day_6.Part2},
	7:  {day_7.Part1, day_7.Part2},
	8:  {day_8.Part1, day_8.Part2},
	9:  {day_9.Part1, day_9.Part2},
	10: {day_10.Part1, day_10.Part2},
	11: {day_11.Part1, day_11.Part2},
	12: {day_12.Part1, day_12.Part2},
}

func main() {
	day := parser.ParseInt(os.Args[1])
	part := parser.ParseInt(os.Args[2]) - 1
	var input string
	if len(os.Args) == 3 || os.Args[3] != "test" {
		input = io.LoadInput(day)
	} else {
		input = io.LoadTestInput(day)
	}

	start := time.Now().UnixMicro()
	advents[day][part](input)
	end := time.Now().UnixMicro()

	println(fmt.Sprintf("\nRuntime:\n%dμs \n%fms", end-start, (float64)(end-start)/1_000))
}
