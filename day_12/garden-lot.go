package day_12

import "aoc24/parser"

type GardenLot = parser.Node[rune]

func hasSamePlant(a, b *GardenLot) bool {
	if a == nil && b == nil {
		return true
	}

	if (a != nil && b == nil) || (a == nil && b != nil) {
		return false
	}

	return a.Value == b.Value
}
