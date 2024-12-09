package day_9

import "github.com/samber/lo"

func calculateChecksum(disk []int) int {
	return lo.Sum(lo.Map(disk, func(id int, index int) int {
		return id * index
	}))
}
