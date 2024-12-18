package day_14

import (
	"aoc24/functional"
	"aoc24/math"
)

func printRobotGrid(robots []Robot, width, height int) {
	byPosition := functional.KeyBy(robots, func(robot Robot, _ int) math.Vector2[int] {
		return robot.position
	})

	for y := range height {
		for x := range width {
			count := len(byPosition[math.Vector2[int]{X: x, Y: y}])
			if count == 0 {
				print(".")
			} else {
				print(count)
			}
		}
		print("\n")
	}
}
