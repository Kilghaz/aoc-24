package day_8

import (
	"aoc24/functional"
	"aoc24/matrix"
	"aoc24/parser"
	"aoc24/vec2"
	"github.com/samber/lo"
)

func isInBound(position vec2.Vector2i, width, height int) bool {
	return position.X >= 0 && position.X < width && position.Y >= 0 && position.Y < height
}

func calculateAntiNodesWithHarmonics(a, b Antenna, width, height int) []vec2.Vector2i {
	diffA := vec2.Subtract(b.position, a.position)
	diffB := vec2.Subtract(a.position, b.position)

	nodes := make([]vec2.Vector2i, 0)
	i := 1
	for {
		nodeA := vec2.Add(vec2.MultiplyScalar(diffA, i), a.position)
		nodeB := vec2.Add(vec2.MultiplyScalar(diffB, i), b.position)

		inBoundA := isInBound(nodeA, width, height)
		inBoundB := isInBound(nodeB, width, height)

		if inBoundA {
			nodes = append(nodes, nodeA)
		}

		if inBoundB {
			nodes = append(nodes, nodeB)
		}

		if !inBoundA && !inBoundB {
			break
		}
		i++
	}

	return nodes
}

func Part2(input string) {
	width, height := matrix.Dimension(parser.ParseGrid(input))
	antennasByFrequency := functional.KeyBy(parseAntennas(input), func(antenna Antenna, _ int) rune {
		return antenna.frequency
	})
	pairs := createUniqueAntennaPairs(antennasByFrequency)
	nodes := lo.Uniq(lo.FlatMap(pairs, func(pair []Antenna, _ int) []vec2.Vector2i {
		return calculateAntiNodesWithHarmonics(pair[0], pair[1], width, height)
	}))

	println(len(nodes))
}
