package day_8

import (
	"aoc24/io"
	"aoc24/math"
	"aoc24/matrix"
)

func parseAntennas(input string) []Antenna {
	var antennas []Antenna
	matrix.Traverse(io.ParseGrid(input), func(frequency rune, x int, y int) {
		if frequency == '.' {
			return
		}
		antennas = append(antennas, Antenna{
			position:  math.Vector2i{x, y},
			frequency: frequency,
		})
	})
	return antennas
}
