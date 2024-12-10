package day_10

import "github.com/samber/lo"

func calculateTrailheadRating(start *Node) int {
	return len(findTrails(start))
}

func Part2(input string) {
	nodes := parseNodes(input)

	trailheads := lo.Filter(nodes, func(item *Node, _ int) bool {
		return item.value == 0
	})

	totalScore := lo.Sum(lo.Map(trailheads, func(head *Node, _ int) int {
		score := calculateTrailheadRating(head)
		return score
	}))

	println(totalScore)
}
