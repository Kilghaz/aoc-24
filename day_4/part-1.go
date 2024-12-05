package day_4

import (
	"github.com/samber/lo"
	"strings"
)

var word = "XMAS"

var directions = [][]int{
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1},
}

func hasEnoughCharacters(grid [][]rune, x int, y int, direction []int) bool {
	if x+direction[1]*len(word) < -1 || x+direction[1]*len(word) > len(grid[y]) {
		return false
	}

	if y+direction[0]*len(word) < -1 || y+direction[0]*len(word) > len(grid) {
		return false
	}

	return true
}

func getCharacters(grid [][]rune, x int, y int, direction []int) string {
	var result []rune
	for i := 0; i < 4; i++ {
		result = append(result, grid[y+direction[0]*i][x+direction[1]*i])
	}
	return string(result)
}

func findXMAS(grid [][]rune, x int, y int) int {
	if grid[y][x] != []rune(word)[0] {
		return 0
	}

	matchingDirections := lo.Filter(directions, func(direction []int, _ int) bool {
		if !hasEnoughCharacters(grid, x, y, direction) {
			return false
		}

		characters := getCharacters(grid, x, y, direction)
		return characters == word
	})

	return len(matchingDirections)
}

func Part1(input string) {
	var grid [][]rune
	lo.ForEach(strings.Split(input, "\n"), func(row string, _ int) {
		grid = append(grid, []rune(row))
	})
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			sum += findXMAS(grid, x, y)
		}
	}
	println(sum)
}
