package day_7

func Part2(input string) {
	equations := parseEquations(input)
	allOperators := []Operator{Add, Multiply, Combine}

	sum := 0
	for _, equation := range equations {
		permutations := createPermutations(equation, allOperators)
		for _, permutation := range permutations {
			if evaluateEquation(permutation) == permutation.result {
				sum += permutation.result
				break
			}
		}
	}
	println(sum)
}
