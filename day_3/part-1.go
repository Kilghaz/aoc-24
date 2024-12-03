package day_3

import (
	"github.com/samber/lo"
)

func Part1(input string) {
	multiplications := lo.Map(mulExp.FindAllString(input, -1), func(match string, index int) int {
		return evalMul(match)
	})
	println(lo.Sum(multiplications))
}
