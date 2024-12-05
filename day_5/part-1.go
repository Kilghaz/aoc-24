package day_5

import "github.com/samber/lo"

func isValidUpdate(update []int, rules map[int]PageRule) bool {
	for i, page := range update {
		sliceBefore := update[0:i]
		sliceAfter := update[i+1:]
		for _, before := range rules[page].before {
			if lo.Contains(sliceBefore, before) {
				return false
			}
		}

		for _, after := range rules[page].after {
			if lo.Contains(sliceAfter, after) {
				return false
			}
		}
	}

	return true
}

func Part1(input string) {
	rules := createRulesMap(parseRules(input))
	updates := parseUpdates(input)

	sum := 0
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			sum += update[(len(update)-1)/2]
		}
	}
	println(sum)
}
