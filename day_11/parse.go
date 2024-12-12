package day_11

import (
	"aoc24/parser"
	"github.com/samber/lo"
	"regexp"
)

func parseStones(input string) []Stone {
	exp, _ := regexp.Compile("\\d+")
	numbers := exp.FindAllString(input, -1)
	return lo.Map(numbers, func(num string, _ int) int {
		return parser.ParseInt(num)
	})
}
