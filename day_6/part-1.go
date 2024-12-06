package day_6

import (
	"aoc24/io"
	"aoc24/math"
	"github.com/samber/lo"
)

func findStartPosition(grid [][]rune) math.Vector2i {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '^' {
				return math.Vector2i{x, y}
			}
		}
	}
	panic("no start position found")
}

func Part1(input string) {
	grid := io.ParseGrid(input)
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
