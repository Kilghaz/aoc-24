package day_10

import (
	"aoc24/io"
	"aoc24/math"
	"aoc24/matrix"
	"errors"
)

var directions = []math.Vector2i{
	{Y: -1},
	{Y: 1},
	{X: -1},
	{X: 1},
}

func parseNodes(input string) []*Node {
	var nodes []*Node
	grid := io.ParseGrid(input)

	matrix.Traverse(grid, func(item rune, x int, y int) {
		position := math.Vector2i{X: x, Y: y}
		node := Node{
			position: position,
			value:    io.ParseInt(string(item)),
			siblings: make([]*Node, 0),
		}
		nodes = append(nodes, &node)
	})

	var findByPosition = func(position math.Vector2i) (*Node, error) {
		for _, node := range nodes {
			if node.position.X == position.X && node.position.Y == position.Y {
				return node, nil
			}
		}
		return nil, errors.New("could not find node")
	}

	for key, node := range nodes {
		siblings := make([]*Node, 0)
		for _, direction := range directions {
			position := math.AddVector2i(node.position, direction)
			sibling, _ := findByPosition(position)
			if sibling == nil {
				continue
			}
			siblings = append(siblings, sibling)
		}
		node.siblings = siblings
		nodes[key] = node
	}

	return nodes
}
