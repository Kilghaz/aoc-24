package day_1

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
	"strings"
)

func parseList(input string, n int) []int {
	lines := strings.Split(input, "\n")
	return lo.Map(lines, func(item string, index int) int {
		exp, _ := regexp.Compile("\\d+")
		match := exp.FindAllString(item, -1)[n]
		parsed, _ := strconv.ParseInt(match, 10, 64)
		return int(parsed)
	})
}
