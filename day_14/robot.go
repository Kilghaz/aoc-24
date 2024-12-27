package day_14

import (
	"aoc24/vec2"
)

type Robot struct {
	position vec2.Vector2i
	velocity vec2.Vector2i
}

func step(robot *Robot, steps int) {
	robot.position = vec2.Add(robot.position, vec2.MultiplyScalar(robot.velocity, steps))
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
