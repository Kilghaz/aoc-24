package day_10

import (
	"github.com/samber/lo"
)

func calculateTrailheadScore(start *Node) int {
	trails := findTrails(start)
	uniqueTops := lo.Uniq(lo.Map(trails, func(trail Trail, _ int) *Node {
		end, _ := lo.Last(trail)
		return end
	}))
	return len(uniqueTops)
}

func Part1(input string) {
	nodes := parseNodes(input)

	trailheads := lo.Filter(nodes, func(item *Node, _ int) bool {
		return item.value == 0
	})

	totalScore := lo.Sum(lo.Map(trailheads, func(head *Node, _ int) int {
		score := calculateTrailheadScore(head)
		return score
	}))

	println(totalScore)
}
