package day_14

import (
	"aoc24/functional"
	"github.com/samber/lo"
)

func calculateSafetyValue(robots []Robot, width, height int) int {
	withoutCenter := lo.Filter(robots, func(robot Robot, _ int) bool {
		return robot.position.X != width/2 && robot.position.Y != height/2
	})
	byQuadrant := functional.KeyBy(withoutCenter, func(robot Robot, _ int) int {
		return int(float64(robot.position.X)/(float64(width)/2)) + 2*int(float64(robot.position.Y)/(float64(height)/2))
	})
	robotCountByQuadrant := lo.Map(lo.Values(byQuadrant), func(robots []Robot, _ int) int {
		return len(robots)
	})
	return lo.Reduce(robotCountByQuadrant, func(a int, count int, _ int) int {
		return a * count
	}, 1)
}

func Part1(input string) {
	robots := parseRobots(input)

	width := 101
	height := 103
	steps := 100

	for i, robot := range robots {
		step(&robot, steps)
		teleport(&robot, width, height)
		robots[i] = robot
	}

	safety := calculateSafetyValue(robots, width, height)

	println(safety)
}
