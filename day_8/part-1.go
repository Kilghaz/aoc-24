package day_8

import (
	"aoc24/functional"
	"aoc24/matrix"
	"aoc24/parser"
	"aoc24/vec2"
	"github.com/samber/lo"
)

func createUniqueAntennaPairs(antennasByFrequency map[rune][]Antenna) [][]Antenna {
	antennaPairs := make([][]Antenna, 0)
	for _, antennas := range antennasByFrequency {
		for i, antenna := range antennas {
			for j := i; j < len(antennas); j++ {
				if i == j {
					continue
				}
				antennaPairs = append(antennaPairs, []Antenna{antenna, antennas[j]})
			}
		}
	}
	return antennaPairs
}

func calculateAntiNodes(a, b Antenna) []vec2.Vector2i {
	diffA := vec2.Subtract(a.position, b.position)
	diffB := vec2.Subtract(b.position, a.position)

	return []vec2.Vector2i{
		vec2.Add(diffA, a.position),
		vec2.Add(diffB, b.position),
	}
}

func Part1(input string) {
	width, height := matrix.Dimension(parser.ParseGrid(input))
	antennasByFrequency := functional.KeyBy(parseAntennas(input), func(antenna Antenna, _ int) rune {
		return antenna.frequency
	})
	pairs := createUniqueAntennaPairs(antennasByFrequency)
	antinodes := lo.Filter(lo.Uniq(lo.FlatMap(pairs, func(pair []Antenna, _ int) []vec2.Vector2i {
		return calculateAntiNodes(pair[0], pair[1])
	})), func(node vec2.Vector2i, _ int) bool {
		return node.X >= 0 && node.X < width && node.Y >= 0 && node.Y < height
	})

	println(len(antinodes))
}
