package day_6

import (
	"aoc24/vec2"
	"github.com/samber/lo"
)

type StopReason string

const (
	Loop        StopReason = "loop"
	OutOfBounds StopReason = "oob"
)

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

var directionVectors = map[Direction]vec2.Vector2i{
	Up:    {0, -1},
	Right: {1, 0},
	Down:  {0, 1},
	Left:  {-1, 0},
}

type Guard struct {
	position  vec2.Vector2i
	direction Direction
}

func rotate(guard *Guard) {
	switch guard.direction {
	case Up:
		guard.direction = Right
	case Right:
		guard.direction = Down
	case Down:
		guard.direction = Left
	case Left:
		guard.direction = Up
	}
}

func move(guard *Guard, visited [][][]Direction, grid [][]rune) StopReason {
	for {
		position := vec2.Add(guard.position, directionVectors[guard.direction])

		if position.X >= len(grid) || position.X < 0 {
			return OutOfBounds
		}

		if position.Y >= len(grid[0]) || position.Y < 0 {
			return OutOfBounds
		}

		if grid[position.X][position.Y] == '#' {
			rotate(guard)
			continue
		}

		if lo.Contains(visited[position.X][position.Y], guard.direction) {
			return Loop
		}

		guard.position = position
		visited[position.X][position.Y] = append(visited[position.X][position.Y], guard.direction)
		continue
	}
}
