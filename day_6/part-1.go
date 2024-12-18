package day_6

import (
	"aoc24/math"
	"aoc24/parser"
	"github.com/samber/lo"
)

func findStartPosition(grid [][]rune) math.Vector2[int] {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '^' {
				return math.Vector2[int]{x, y}
			}
		}
	}
	panic("no start position found")
}

func Part1(input string) {
	grid := parser.ParseGrid(input)
	startPosition := findStartPosition(grid)
	visited := lo.Map(grid, func(row []rune, _ int) [][]Direction {
		return lo.Map(row, func(item rune, _ int) []Direction {
			return []Direction{}
		})
	})
	guard := Guard{
		position:  startPosition,
		direction: Up,
	}
	move(&guard, visited, grid)

	sum := 0
	for _, row := range visited {
		for _, col := range row {
			if len(col) > 0 {
				sum += 1
			}
		}
	}
	print(sum)
}
