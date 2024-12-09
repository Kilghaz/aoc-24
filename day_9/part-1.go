package day_9

import (
	"aoc24/slice"
	"github.com/samber/lo"
)

var blockTypeOperation = map[BlockType]func(data *[]int, size int) []int{
	File:      slice.ShiftN[int],
	FreeSpace: slice.PopN[int],
}

func Part1(input string) {
	blocks := parseBlocks(input)

	fileData := lo.FlatMap(blocks, func(block Block, index int) []int {
		if block.blockType == File {
			return slice.Fill(make([]int, block.size), block.id)
		}
		return []int{}
	})

	disk := make([]int, 0)
	for _, block := range blocks {
		dataSize := min(block.size, len(fileData))
		disk = append(disk, blockTypeOperation[block.blockType](&fileData, dataSize)...)
	}

	checksum := calculateChecksum(disk)
	println(checksum)
}
