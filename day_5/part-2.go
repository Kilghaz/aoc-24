package day_5

import (
	"github.com/samber/lo"
	"slices"
)

func sortUpdate(update []int, rules map[int]PageRule) []int {
	sorted := slices.Clone(update)
	slices.SortFunc(sorted, func(a, b int) int {
		if lo.Contains(rules[a].before, b) {
			return -1
		}

		if lo.Contains(rules[a].after, b) {
			return 1
		}

		return 0
	})
	return sorted
}

func Part2(input string) {
	rules := createRulesMap(parseRules(input))
	updates := parseUpdates(input)

	sum := 0
	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			fixedUpdate := sortUpdate(update, rules)
			sum += fixedUpdate[(len(update)-1)/2]
		}
	}
	println(sum)
}
