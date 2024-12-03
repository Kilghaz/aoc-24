package day_2

import (
	"github.com/samber/lo"
	"slices"
)

func withoutLevel(report []int, index int) []int {
	return append(slices.Clone(report)[:index], report[index+1:]...)
}

func isSafeDampened(report []int) bool {
	for i := 0; i < len(report); i++ {
		if isSafe(withoutLevel(report, i)) {
			return true
		}
	}
	return false
}

func Part2(input string) {
	reports := parseReports(input)
	println(len(lo.Filter(reports, func(report []int, index int) bool {
		return isSafeDampened(report)
	})))
}
