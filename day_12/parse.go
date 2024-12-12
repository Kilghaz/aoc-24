package day_12

import (
	"aoc24/functional"
	"aoc24/parser"
)

type GardenLot = parser.Node[rune]

func parseGardenLots(input string) []*GardenLot {
	return parser.ParseNodes(input, functional.Identity[rune]())
}
