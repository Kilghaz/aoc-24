package day_13

func Part2(input string) {
	machines := parseClawMachines(input)

	tokens := 0
	for _, machine := range machines {
		if !canSolve(machine) {
			continue
		}

		machine.prize.X += 10000000000000
		machine.prize.Y += 10000000000000

		a, b := solve(machine)
		px, py := press(machine, a, b)

		if px != machine.prize.X || py != machine.prize.Y {
			continue
		}

		tokens += calculateTokens(a, b)
	}

	println(tokens)
}
