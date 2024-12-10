package day_10

import "aoc24/math"

type Node struct {
	position math.Vector2i
	value    int
	siblings []*Node
}
