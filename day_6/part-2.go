package day_6

import (
	"aoc24/io"
	"aoc24/matrix"
	"github.com/samber/lo"
)

func createGridPermutations(grid [][]rune) [][][]rune {
	permutations := make([][][]rune, 0)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '#' {
				continue
			}

			if grid[x][y] == '^' {
				continue
			}

			clone := matrix.Clone(grid)
			clone[x][y] = '#'
			permutations = append(permutations, clone)
		}
	}

	return permutations
}

func Part2(input string) {
	grid := io.ParseGrid(input)
	startPosition := findStartPosition(grid)
	visited := lo.Map(grid, func(row []rune, _ int) [][]Direction {
		return lo.Map(row, func(item rune, _ int) []Direction {
			return []Direction{}
		})
	})

	sum := 0
	permutations := createGridPermutations(grid)
	guard := Guard{
		position:  startPosition,
		direction: Up,
	}

	for _, permutation := range permutations {
		guard.position = startPosition
		guard.direction = Up

		if move(&guard, matrix.Clone(visited), permutation) == Loop {
			sum++
		}
	}

	print(sum)
}
