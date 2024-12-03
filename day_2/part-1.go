package day_2

import (
	"aoc24/math"
	"github.com/samber/lo"
)

func isSafe(report []int) bool {
	reportSign := math.Sign(report[0] - report[1])
	for i := 1; i < len(report); i++ {
		diff := math.Difference(report[i-1], report[i])
		if diff < 1 || diff > 3 || math.Sign(report[i-1]-report[i]) != reportSign {
			return false
		}
	}
	return true
}

func Part1(input string) {
	reports := parseReports(input)
	println(len(lo.Filter(reports, func(report []int, index int) bool {
		return isSafe(report)
	})))
}
