package day_7

import (
	math2 "aoc24/math"
	"aoc24/slice"
	"github.com/samber/lo"
)

type Operator func(a, b int) int

var Multiply Operator = func(a, b int) int {
	return a * b
}

var Add Operator = func(a, b int) int {
	return a + b
}

var Combine Operator = func(a, b int) int {
	if b == 0 {
		return a * 10
	}
	digits := math2.Log10(b) + 1
	return a*math2.Pow(10, digits) + b
}

type Equation struct {
	result    int
	constants []int
	operators []Operator
}

func evaluateEquation(equation Equation) int {
	current := 0
	for i, constant := range equation.constants {
		if current == 0 {
			current = constant
			continue
		}
		current = equation.operators[i-1](current, constant)
	}
	return current
}

func createPermutations(equation Equation, allOperators []Operator) []Equation {
	operatorCount := len(allOperators)
	combinationCount := math2.Pow(operatorCount, len(equation.constants)-1)

	permutations := make([]Equation, combinationCount)
	for i := range combinationCount {
		digits := slice.Pad(math2.ToBase(i, operatorCount), len(equation.constants)-1, 0)
		operators := lo.Map(digits, func(item int, _ int) Operator {
			return allOperators[item]
		})
		permutations[i] = Equation{
			result:    equation.result,
			constants: equation.constants,
			operators: operators,
		}
	}
	return permutations
}
