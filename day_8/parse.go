package day_8

import (
	"aoc24/math"
	"aoc24/matrix"
	"aoc24/parser"
)

func parseAntennas(input string) []Antenna {
	var antennas []Antenna
	matrix.Traverse(parser.ParseGrid(input), func(frequency rune, x int, y int) {
		if frequency == '.' {
			return
		}
		antennas = append(antennas, Antenna{
			position:  math.Vector2[int]{x, y},
			frequency: frequency,
		})
	})
	return antennas
}
