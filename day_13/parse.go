package day_13

import (
	"aoc24/parser"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func parseClawMachines(input string) []ClawMachine {
	return lo.Map(strings.Split(input, "\n\n"), func(machineInput string, _ int) ClawMachine {
		exp, _ := regexp.Compile("\\d+")
		numbers := exp.FindAllString(machineInput, -1)
		return ClawMachine{
			buttonA: Delta{X: parser.ParseInt(numbers[0]), Y: parser.ParseInt(numbers[1])},
			buttonB: Delta{X: parser.ParseInt(numbers[2]), Y: parser.ParseInt(numbers[3])},
			prize:   Location{X: parser.ParseInt(numbers[4]), Y: parser.ParseInt(numbers[5])},
		}
	})
}
