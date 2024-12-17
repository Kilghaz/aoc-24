package day_14

import "aoc24/math"

type Robot struct {
	position math.Vector2i
	velocity math.Vector2i
}

func step(robot *Robot, steps int) {
	robot.position = math.AddVector2i(robot.position, math.MultiplyScalar(robot.velocity, steps))
}

func teleport(robot *Robot, width, height int) {
	robot.position.X %= width
	robot.position.Y %= height

	if robot.position.X < 0 {
		robot.position.X = width + robot.position.X
	}

	if robot.position.Y < 0 {
		robot.position.Y = height + robot.position.Y
	}
}
