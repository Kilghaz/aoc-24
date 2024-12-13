package parser

import (
	"aoc24/math"
	"aoc24/matrix"
	"errors"
)

type Node[T any] struct {
	Position math.Vector2i
	Value    T
	Siblings []*Node[T]
}

func ParseNodes[T any](input string, iterator func(item rune) T) []*Node[T] {
	var nodes []*Node[T]
	grid := ParseGrid(input)

	matrix.Traverse(grid, func(item rune, x int, y int) {
		position := math.Vector2i{X: x, Y: y}
		node := Node[T]{
			Position: position,
			Value:    iterator(item),
			Siblings: make([]*Node[T], 0),
		}
		nodes = append(nodes, &node)
	})

	var findByPosition = func(position math.Vector2i) (*Node[T], error) {
		for _, node := range nodes {
			if node.Position.X == position.X && node.Position.Y == position.Y {
				return node, nil
			}
		}
		return nil, errors.New("could not find node")
	}

	for key, node := range nodes {
		siblings := make([]*Node[T], 0)
		for _, direction := range math.OrthogonalDirections {
			position := math.AddVector2i(node.Position, direction)
			sibling, _ := findByPosition(position)
			if sibling == nil {
				continue
			}
			siblings = append(siblings, sibling)
		}
		node.Siblings = siblings
		nodes[key] = node
	}

	return nodes
}
