package day_1

import (
	"github.com/samber/lo"
	"math"
	"slices"
)

func Part1(input string) {
	firstList := parseList(input, 0)
	secondList := parseList(input, 1)

	slices.Sort(firstList)
	slices.Sort(secondList)

	diffs := lo.Map(firstList, func(item int, index int) int {
		return int(math.Abs(float64(item - secondList[index])))
	})

	println(lo.Sum(diffs))
}
