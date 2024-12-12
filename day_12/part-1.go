package day_12

import (
	"github.com/samber/lo"
)

func calculatePerimeter(region Region) int {
	perimeter := 0
	for _, lot := range region {
		perimeter += 4 - len(lo.Filter(lot.Siblings, func(item *GardenLot, _ int) bool {
			return item.Value == lot.Value
		}))
	}
	return perimeter
}

func calculateArea(region Region) int {
	return len(region)
}

func Part1(input string) {
	gardenLots := parseGardenLots(input)
	regions := findRegions(gardenLots)

	sum := 0
	for _, region := range regions {
		perimeter := calculatePerimeter(region)
		area := calculateArea(region)
		sum += perimeter * area
	}
	println(sum)
}
