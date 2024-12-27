package day_14

import (
	"aoc24/parser"
	"aoc24/vec2"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func parseRobots(input string) []Robot {
	return lo.Map(strings.Split(input, "\n"), func(line string, _ int) Robot {
		exp, _ := regexp.Compile("-?\\d+")
		numbers := exp.FindAllString(line, -1)
		return Robot{
			position: vec2.Vector2i{X: parser.ParseInt(numbers[0]), Y: parser.ParseInt(numbers[1])},
			velocity: vec2.Vector2i{X: parser.ParseInt(numbers[2]), Y: parser.ParseInt(numbers[3])},
		}
	})
}
