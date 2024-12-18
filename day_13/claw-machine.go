package day_13

import "aoc24/math"

type Delta = math.Vector2[int]
type Location = math.Vector2[int]

type ClawMachine struct {
	buttonA Delta
	buttonB Delta
	prize   Location
}

// this will solve the closest match with the fewest possible presses rounded down
// it is sometimes possible to solve a machine by doing half a/b presses
func solve(machine ClawMachine) (int, int) {
	a := (machine.buttonB.Y*machine.prize.X - machine.buttonB.X*machine.prize.Y) / (machine.buttonA.X*machine.buttonB.Y - machine.buttonA.Y*machine.buttonB.X)
	b := (machine.prize.X - a*machine.buttonA.X) / machine.buttonB.X
	return a, b
}

func canSolve(machine ClawMachine) bool {
	return (machine.buttonA.X*machine.buttonB.Y - machine.buttonA.Y*machine.buttonB.X) != 0
}

func press(machine ClawMachine, a, b int) (int, int) {
	return a*machine.buttonA.X + b*machine.buttonB.X, a*machine.buttonA.Y + b*machine.buttonB.Y
}
