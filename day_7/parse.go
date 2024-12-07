package day_7

import (
	"aoc24/io"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func parseEquations(input string) []Equation {
	return lo.Map(strings.Split(input, "\n"), func(line string, _ int) Equation {
		exp, _ := regexp.Compile("(\\d+)")
		match := exp.FindAllString(line, -1)
		result := io.ParseInt(match[0])
		constants := lo.Map(match[1:], func(it string, _ int) int {
			return io.ParseInt(it)
		})
		return Equation{
			result:    result,
			constants: constants,
		}
	})
}
