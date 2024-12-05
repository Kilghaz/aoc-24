package day_4

import (
	"aoc24/matrix"
	"github.com/samber/lo"
	"strings"
)

var pattern = [][]rune{
	{'M', ' ', 'M'},
	{' ', 'A', ' '},
	{'S', ' ', 'S'},
}

var patternHashes = []string{
	hashMatrix(matrix.RotateN(pattern, 0)),
	hashMatrix(matrix.RotateN(pattern, 1)),
	hashMatrix(matrix.RotateN(pattern, 2)),
	hashMatrix(matrix.RotateN(pattern, 3)),
}

func hashMatrix(matrix [][]rune) string {
	return string(lo.Filter(lo.Flatten(matrix), func(char rune, i int) bool {
		return i%2 == 0
	}))
}

func matchesPattern(grid [][]rune, x int, y int) bool {
	hash := hashMatrix(lo.Map(grid[x-1:x+2], func(row []rune, _ int) []rune {
		return row[y-1 : y+2]
	}))

	for _, patternHash := range patternHashes {
		if hash == patternHash {
			return true
		}
	}

	return false
}

func Part2(input string) {
	grid := matrix.Invert(lo.Map(strings.Split(input, "\n"), func(row string, _ int) []rune {
		return []rune(row)
	}))

	sum := 0
	matrix.TraverseSymmetricallyOffset(grid, 1, func(_ rune, x int, y int) {
		if matchesPattern(grid, x, y) {
			sum += 1
		}
	})
	println(sum)
}
