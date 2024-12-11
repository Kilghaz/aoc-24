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

	sum := 0
	for _, stone := range stones {
		calculated := calculateStoneCount(stone, blinkCount-1)
		cache[BlinkKey{stone, blinkCount - 1}] = calculated
		sum += calculated
	}

	return sum
}

func Part2(input string) {
	stones := parseStones(input)
	count := 75

	sum := 0
	for _, stone := range stones {
		sum += calculateStoneCount(stone, count)
	}
	println(sum)
}
