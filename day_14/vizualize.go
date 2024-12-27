package day_14

import (
	"aoc24/functional"
	"aoc24/vec2"
)

func printRobotGrid(robots []Robot, width, height int) {
	byPosition := functional.KeyBy(robots, func(robot Robot, _ int) vec2.Vector2i {
		return robot.position
	})

	for y := range height {
		for x := range width {
			count := len(byPosition[vec2.Vector2i{X: x, Y: y}])
			if count == 0 {
				print(".")
			} else {
				print(count)
			}
		}
		print("\n")
	}
}
