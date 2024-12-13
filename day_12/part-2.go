package day_12

import (
	"aoc24/math"
	"github.com/samber/lo"
)

func calculateSides(region Region) int {
	sides := 0

	lotsByPosition := lo.KeyBy(region, func(lot *GardenLot) Position {
		return lot.Position
	})

	for _, lot := range region {
		for _, direction := range math.DiagonalDirections {
			cornerLot := lotsByPosition[math.AddVector2i(lot.Position, direction)]
			horizontalLot := lotsByPosition[math.AddVector2i(lot.Position, Position{X: direction.X})]
			verticalLot := lotsByPosition[math.AddVector2i(lot.Position, Position{Y: direction.Y})]

			isCorner := hasSamePlant(horizontalLot, verticalLot)
			isInsideCorner := isCorner && !hasSamePlant(horizontalLot, lot)
			isOutsideCorner := isCorner && !isInsideCorner && !hasSamePlant(cornerLot, lot)

			if isInsideCorner {
				sides++
			}

			if isOutsideCorner {
				sides++
			}
		}
	}

	return sides
}

func Part2(input string) {
	gardenLots := parseGardenLots(input)
	regions := findRegions(gardenLots)

	totalPrice := 0
	for _, region := range regions {
		totalPrice += calculateArea(region) * calculateSides(region)
	}

	println(totalPrice)
}
