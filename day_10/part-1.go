package day_10

import (
	"github.com/samber/lo"
)

func calculateTrailheadScore(start *PathNode) int {
	trails := findTrails(start)
	uniqueTops := lo.Uniq(lo.Map(trails, func(trail Trail, _ int) *PathNode {
		end, _ := lo.Last(trail)
		return end
	}))
	return len(uniqueTops)
}

func Part1(input string) {
	nodes := parsePathNodes(input)

	trailheads := lo.Filter(nodes, func(item *PathNode, _ int) bool {
		return item.Value == 0
	})

	totalScore := lo.Sum(lo.Map(trailheads, func(head *PathNode, _ int) int {
		score := calculateTrailheadScore(head)
		return score
	}))

	println(totalScore)
}
