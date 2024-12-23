package main

import (
	"aoc24/day_0"
	"aoc24/day_1"
	"aoc24/day_10"
	"aoc24/day_11"
	"aoc24/day_12"
	"aoc24/day_13"
	"aoc24/day_14"
	"aoc24/day_15"
	"aoc24/day_16"
	"aoc24/day_17"
	"aoc24/day_18"
	"aoc24/day_19"
	"aoc24/day_2"
	"aoc24/day_20"
	"aoc24/day_21"
	"aoc24/day_22"
	"aoc24/day_23"
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
	13: {day_13.Part1, day_13.Part2},
	14: {day_14.Part1, day_14.Part2},
	0:  {day_0.Part1, day_0.Part2},
	15: {day_15.Part1, day_15.Part2},
	16: {day_16.Part1, day_16.Part2},
	17: {day_17.Part1, day_17.Part2},
	18: {day_18.Part1, day_18.Part2},
	19: {day_19.Part1, day_19.Part2},
	20: {day_20.Part1, day_20.Part2},
	21: {day_21.Part1, day_21.Part2},
	22: {day_22.Part1, day_22.Part2},
	23: {day_23.Part1, day_23.Part2},
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
