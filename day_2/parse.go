package day_2

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func mapToInt(item string, index int) int {
	parsed, _ := strconv.ParseInt(item, 10, 64)
	return int(parsed)
}

func parseReports(input string) [][]int {
	lines := strings.Split(input, "\n")

	return lo.Map(lines, func(item string, index int) []int {
		return lo.Map(strings.Split(item, " "), mapToInt)
	})
}
