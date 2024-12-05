package day_5

import (
	"aoc24/io"
	"github.com/samber/lo"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func parseRules(input string) []Rule {
	rulesInput := strings.Split(input, "\n\n")[0]
	return lo.Map(strings.Split(rulesInput, "\n"), func(line string, _ int) Rule {
		parts := strings.Split(line, "|")
		return Rule{
			before: io.ParseInt(parts[0]),
			after:  io.ParseInt(parts[1]),
		}
	})
}

func parseUpdates(input string) [][]int {
	updatesInput := strings.Split(input, "\n\n")[1]
	return lo.Map(strings.Split(updatesInput, "\n"), func(line string, _ int) []int {
		return lo.Map(strings.Split(line, ","), func(page string, _ int) int {
			return io.ParseInt(page)
		})
	})
}
