package day_14

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

var patternFile = "day_14/pattern.png"

func calculateSimilarity(robots []Robot, pattern Pattern) float64 {
	avgError := 0.0
	for _, robot := range robots {
		x := robot.position.X
		y := robot.position.Y
		avgError += pattern[y][x]
	}
	return 1.0 - (avgError / float64(len(robots)))
}

func Part2(input string) {
	if _, err := os.Stat(patternFile); errors.Is(err, os.ErrNotExist) {
		trainPattern(patternFile)
	}
	pattern := loadPattern(patternFile)
	robots := parseRobots(input)

	width := 101
	height := 103
	steps := 0

	increments := 1
	threshold := 0.75
	bestMatchSteps := 0
	bestMatchSimilarity := 0.0
	var bestRobots []Robot

	for bestMatchSimilarity < threshold && steps < 200000 {
		robotClones := slices.Clone(robots)
		for i, robot := range robotClones {
			step(&robot, steps)
			teleport(&robot, width, height)
			robotClones[i] = robot
		}

		similarity := calculateSimilarity(robotClones, pattern)

		if similarity > bestMatchSimilarity {
			bestMatchSimilarity = similarity
			bestMatchSteps = steps
			bestRobots = robotClones
		}

		steps += increments
	}

	println(fmt.Sprintf("similarity %f, steps: %d", bestMatchSimilarity, bestMatchSteps))
	printRobotGrid(bestRobots, width, height)

	println(bestMatchSteps)
}
