package day_3

import (
	"regexp"
	"strconv"
)

var mulExp, _ = regexp.Compile("mul\\((\\d+),(\\d+)\\)")

func evalMul(input string) int {
	match := mulExp.FindStringSubmatch(input)
	parsedA, _ := strconv.ParseInt(match[1], 10, 64)
	parsedB, _ := strconv.ParseInt(match[2], 10, 64)
	return int(parsedA) * int(parsedB)
}
