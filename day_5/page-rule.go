package day_5

import "github.com/samber/lo"

type PageRule struct {
	before []int
	after  []int
}

func createRulesMap(rules []Rule) map[int]PageRule {
	mapping := make(map[int]PageRule)
	lo.ForEach(rules, func(item Rule, _ int) {
		if !lo.HasKey(mapping, item.before) {
			mapping[item.before] = PageRule{
				before: make([]int, 0),
				after:  make([]int, 0),
			}
		}

		if !lo.HasKey(mapping, item.after) {
			mapping[item.after] = PageRule{
				before: make([]int, 0),
				after:  make([]int, 0),
			}
		}

		pageRuleBefore := mapping[item.before]
		pageRuleBefore.before = append(pageRuleBefore.before, item.after)
		mapping[item.before] = pageRuleBefore

		pageRuleAfter := mapping[item.after]
		pageRuleAfter.after = append(pageRuleAfter.after, item.before)
		mapping[item.after] = pageRuleAfter
	})
	return mapping
}
