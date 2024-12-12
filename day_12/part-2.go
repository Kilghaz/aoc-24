package day_12

import (
	"aoc24/math"
	"aoc24/matrix"
	"aoc24/parser"
)

var directions = []Position{
	{Y: -1},
	{Y: 1},
	{X: -1},
	{X: 1},
}

var cornersOffsets = []Position{
	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
}

func isWithinBounds(i Position, maxRow, maxCol int) bool {
	return i.X >= 0 && i.Y >= 0 && i.X < maxRow && i.Y < maxCol
}

func isSamePlant(a, b Position, grid [][]rune) bool {
	width, height := matrix.Dimension(grid)

	if !isWithinBounds(a, width, height) && !isWithinBounds(b, width, height) {
		return true
	}

	if isWithinBounds(a, width, height) && isWithinBounds(b, width, height) {
		return grid[a.X][a.Y] == grid[b.X][b.Y]
	}

	return false
}

func countCorners(lotPosition Position, grid [][]rune, visited map[Position]bool, currPlant rune, area, corners *int) {
	maxRow, maxCol := len(grid), len(grid[0])
	visited[lotPosition] = true

	for _, dir := range directions {
		nextIndex := Position{X: lotPosition.X + dir.X, Y: lotPosition.Y + dir.Y}
		if isWithinBounds(nextIndex, maxRow, maxCol) && grid[nextIndex.X][nextIndex.Y] == currPlant {
			if !visited[nextIndex] {
				*area++
				countCorners(nextIndex, grid, visited, currPlant, area, corners)
			}
		}
	}

	for _, cornerOffset := range cornersOffsets {
		cornerPos := math.AddVector2i(lotPosition, cornerOffset)

		horizontal := math.AddVector2i(lotPosition, Position{X: cornerOffset.X})
		vertical := math.AddVector2i(lotPosition, Position{Y: cornerOffset.Y})

		if !isSamePlant(horizontal, lotPosition, grid) && !isSamePlant(vertical, lotPosition, grid) {
			*corners++
		}

		if isSamePlant(horizontal, lotPosition, grid) && isSamePlant(vertical, lotPosition, grid) && !isSamePlant(lotPosition, cornerPos, grid) {
			*corners++
		}
	}
}

func Part2(input string) {
	grid := parser.ParseGrid(input)
	totalPrice := 0
	visited := make(map[Position]bool)
	matrix.Traverse(grid, func(plant rune, x int, y int) {
		position := Position{X: x, Y: y}
		if !visited[position] {
			var area, corners int
			countCorners(position, grid, visited, plant, &area, &corners)
			totalPrice += (area + 1) * corners
		}
	})
	println(totalPrice)
}
