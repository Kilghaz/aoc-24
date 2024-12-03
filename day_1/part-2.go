package day_1

import (
	"github.com/samber/lo"
	"slices"
)

func Part2(input string) {
	firstList := parseList(input, 0)
	secondList := parseList(input, 1)

	slices.Sort(firstList)
	slices.Sort(secondList)

	score := lo.Map(firstList, func(item int, index int) int {
		return item * lo.Count(secondList, item)
	})

	println(lo.Sum(score))
}
