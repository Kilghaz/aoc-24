package day_8

import (
	"aoc24/matrix"
	"aoc24/parser"
	"aoc24/vec2"
)

func parseAntennas(input string) []Antenna {
	var antennas []Antenna
	matrix.Traverse(parser.ParseGrid(input), func(frequency rune, x int, y int) {
		if frequency == '.' {
			return
		}
		antennas = append(antennas, Antenna{
			position:  vec2.Vector2i{X: x, Y: y},
			frequency: frequency,
		})
	})
	return antennas
}
