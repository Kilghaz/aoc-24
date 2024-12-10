package day_10

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"slices"
)

type Trail = []*Node

func createId() string {
	id, _ := uuid.NewUUID()
	return id.String()
}

func findTrails(start *Node) []Trail {
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

			if end.value == 9 {
				finishedTrails = append(finishedTrails, trail)
				delete(unfinishedTrails, id)
				continue
			}

			siblings := lo.Filter(end.siblings, func(sibling *Node, index int) bool {
				return sibling.value == end.value+1
			})

			if len(siblings) == 0 {
				delete(unfinishedTrails, id)
				continue
			}

			branchingTrails := lo.Map(siblings, func(node *Node, _ int) Trail {
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
