package day_11

import "github.com/samber/lo"

type BlinkKey struct {
	stone Stone
	count int
}

var cache = map[BlinkKey]int{}

func calculateStoneCount(stone Stone, blinkCount int) int {
	if lo.HasKey(cache, BlinkKey{stone, blinkCount}) {
		return cache[BlinkKey{stone, blinkCount}]
	}

	stones := blink(stone)

	if blinkCount == 1 {
		return len(stones)
	}

	sum := lo.Map(stones, func(stone Stone, _ int) int {
		calculated := calculateStoneCount(stone, blinkCount-1)
		cache[BlinkKey{stone, blinkCount - 1}] = calculated
		return calculated
	})

	return lo.Sum(sum)
}

func Part2(input string) {
	stones := parseStones(input)
	count := 75

	stones = lo.Map(stones, func(stone Stone, _ int) int {
		return calculateStoneCount(stone, count)
	})

	println(lo.Sum(stones))
}
