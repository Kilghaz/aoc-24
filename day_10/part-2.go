package day_10

import (
	"github.com/samber/lo"
)

func calculateTrailheadRating(start *PathNode) int {
	return len(findTrails(start))
}

func Part2(input string) {
	nodes := parsePathNodes(input)

	trailheads := lo.Filter(nodes, func(item *PathNode, _ int) bool {
		return item.Value == 0
	})

	totalScore := lo.Sum(lo.Map(trailheads, func(head *PathNode, _ int) int {
		score := calculateTrailheadRating(head)
		return score
	}))

	println(totalScore)
}
