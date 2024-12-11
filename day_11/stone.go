package day_11

import "aoc24/math"

type Stone = int

type BlinkOperation struct {
	canApply func(stone Stone) bool
	apply    func(stone Stone) []Stone
}

func hasEvenDigits(stone Stone) bool {
	return math.CountDigits(stone)%2 == 0
}

func isZero(stone Stone) bool {
	return stone == 0
}

func splitDigits(stone Stone) []Stone {
	count := math.CountDigits(stone) / 2
	power10count := math.Pow(10, count)
	secondStone := stone % power10count
	for ; stone > power10count-1; stone /= 10 {
	}
	return []Stone{
		stone,
		secondStone,
	}
}

func flipToOne(stone Stone) []Stone {
	return []Stone{1}
}

func multiplyWith2024(stone Stone) []Stone {
	return []Stone{stone * 2024}
}

var stoneOperations = []BlinkOperation{
	{isZero, flipToOne},
	{hasEvenDigits, splitDigits},
}

var fallbackOperation = multiplyWith2024

func blink(stone Stone) []Stone {
	for _, operation := range stoneOperations {
		if operation.canApply(stone) {
			return operation.apply(stone)
		}
	}
	return fallbackOperation(stone)
}
