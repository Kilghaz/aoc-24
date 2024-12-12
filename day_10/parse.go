package day_10

import "aoc24/parser"

func parsePathNodes(input string) []*PathNode {
	return parser.ParseNodes(input, func(item rune) int {
		return parser.ParseInt(string(item))
	})
}
