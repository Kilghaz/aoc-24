package day_11

import (
	"github.com/samber/lo"
)

func Part1(input string) {
	stones := parseStones(input)

	for range 25 {
		stones = lo.FlatMap(stones, func(stone Stone, _ int) []Stone {
			return blink(stone)
		})
	}

	println(len(stones))
}
