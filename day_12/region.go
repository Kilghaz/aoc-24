package day_12

import (
	"aoc24/slice"
	"github.com/samber/lo"
)

type Region = []*GardenLot

func expandRegion(current *GardenLot, plantType rune, region *Region, visited *[]Position) {
	if lo.Contains(*visited, current.Position) {
		return
	}

	if current.Value != plantType {
		return
	}

	slice.Push(visited, current.Position)
	slice.Push(region, current)

	for _, sibling := range current.Siblings {
		expandRegion(sibling, plantType, region, visited)
	}
}

func findRegions(gardenLots []*GardenLot) []Region {
	lotsByPosition := lo.KeyBy(gardenLots, func(node *GardenLot) Position {
		return node.Position
	})

	var visited []Position
	var regions []Region

	for position, node := range lotsByPosition {
		if lo.Contains(visited, node.Position) {
			continue
		}

		if lo.Contains(visited, position) {
			continue
		}

		var region Region
		expandRegion(node, node.Value, &region, &visited)
		regions = append(regions, region)
	}
	return regions
}
