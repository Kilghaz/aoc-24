package day_8

import (
	"aoc24/functional"
	"aoc24/math"
	"aoc24/matrix"
	"aoc24/parser"
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

func calculateAntiNodes(a, b Antenna) []math.Vector2[int] {
	diffA := math.SubVector2(a.position, b.position)
	diffB := math.SubVector2(b.position, a.position)

	return []math.Vector2[int]{
		math.AddVector2(diffA, a.position),
		math.AddVector2(diffB, b.position),
	}
}

func Part1(input string) {
	width, height := matrix.Dimension(parser.ParseGrid(input))
	antennasByFrequency := functional.KeyBy(parseAntennas(input), func(antenna Antenna, _ int) rune {
		return antenna.frequency
	})
	pairs := createUniqueAntennaPairs(antennasByFrequency)
	antinodes := lo.Filter(lo.Uniq(lo.FlatMap(pairs, func(pair []Antenna, _ int) []math.Vector2[int] {
		return calculateAntiNodes(pair[0], pair[1])
	})), func(node math.Vector2[int], _ int) bool {
		return node.X >= 0 && node.X < width && node.Y >= 0 && node.Y < height
	})

	println(len(antinodes))
}
