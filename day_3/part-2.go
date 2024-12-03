package day_3

import (
	"regexp"
)

func parseExp(input string, index *int, exp *regexp.Regexp, onMatch func(match string)) {
	matchIndex := exp.FindStringIndex(input[*index:])
	if len(matchIndex) == 0 || matchIndex[0] != 0 {
		return
	}
	match := exp.FindString(input[*index:])
	*index += len(match) - 1
	onMatch(match)
}

func Part2(input string) {
	skip := false
	sum := 0
	var doExp, _ = regexp.Compile("do\\(\\)")
	var dontExp, _ = regexp.Compile("don't\\(\\)")

	var parseDont = func(index *int) {
		parseExp(input, index, dontExp, func(_ string) {
			skip = true
		})
	}

	var parseDo = func(index *int) {
		parseExp(input, index, doExp, func(_ string) {
			skip = false
		})
	}

	var parseMul = func(index *int) {
		if skip {
			return
		}

		parseExp(input, index, mulExp, func(match string) {
			sum += evalMul(match)
		})
	}

	for i := 0; i < len(input); i++ {
		parseDont(&i)
		parseDo(&i)
		parseMul(&i)
	}

	println(sum)
}
