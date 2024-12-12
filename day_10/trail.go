package day_10

import (
	"aoc24/parser"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"slices"
)

type PathNode = parser.Node[int]
type Trail = []*PathNode

func createId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func findTrails(start *PathNode) []Trail {
	unfinishedTrails := map[string]Trail{
		createId(): {start},
	}

	finishedTrails := make([]Trail, 0)

	for {
		if len(unfinishedTrails) == 0 {
			break
		}

		for id, trail := range unfinishedTrails {
			end, _ := lo.Last(trail)

			if end.Value == 9 {
				finishedTrails = append(finishedTrails, trail)
				delete(unfinishedTrails, id)
				continue
			}

			siblings := lo.Filter(end.Siblings, func(sibling *PathNode, index int) bool {
				return sibling.Value == end.Value+1
			})

			if len(siblings) == 0 {
				delete(unfinishedTrails, id)
				continue
			}

			branchingTrails := lo.Map(siblings, func(node *PathNode, _ int) Trail {
				return append(slices.Clone(trail), node)
			})

			unfinishedTrails[id] = branchingTrails[0]

			if len(branchingTrails) > 1 {
				for _, branch := range branchingTrails[1:] {
					unfinishedTrails[createId()] = branch
				}
			}
		}
	}

	return finishedTrails
}
