package day_7

import (
	"aoc24/parser"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func parseEquations(input string) []Equation {
	return lo.Map(strings.Split(input, "\n"), func(line string, _ int) Equation {
		exp, _ := regexp.Compile("(\\d+)")
		match := exp.FindAllString(line, -1)
		result := parser.ParseInt(match[0])
		constants := lo.Map(match[1:], func(it string, _ int) int {
			return parser.ParseInt(it)
		})
		return Equation{
			result:    result,
			constants: constants,
		}
	})
}
