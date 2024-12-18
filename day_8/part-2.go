package day_8

import (
	"aoc24/functional"
	"aoc24/math"
	"aoc24/matrix"
	"aoc24/parser"
	"github.com/samber/lo"
)

func isInBound(position math.Vector2[int], width, height int) bool {
	return position.X >= 0 && position.X < width && position.Y >= 0 && position.Y < height
}

func calculateAntiNodesWithHarmonics(a, b Antenna, width, height int) []math.Vector2[int] {
	diffA := math.SubVector2(b.position, a.position)
	diffB := math.SubVector2(a.position, b.position)

	nodes := make([]math.Vector2[int], 0)
	i := 1
	for {
		nodeA := math.AddVector2(math.MultiplyScalarVector2(diffA, i), a.position)
		nodeB := math.AddVector2(math.MultiplyScalarVector2(diffB, i), b.position)

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
	nodes := lo.Uniq(lo.FlatMap(pairs, func(pair []Antenna, _ int) []math.Vector2[int] {
		return calculateAntiNodesWithHarmonics(pair[0], pair[1], width, height)
	}))

	println(len(nodes))
}
