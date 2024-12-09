package day_9

import (
	"aoc24/slice"
	"github.com/samber/lo"
)

func Part1(input string) {
	blocks := parseBlocks(input)

	fileBlocks := lo.Filter(blocks, func(block Block, index int) bool {
		return block.blockType == File
	})

	fileData := lo.FlatMap(fileBlocks, func(block Block, index int) []int {
		return slice.Fill(make([]int, block.size), index)
	})

	disk := make([]int, 0)
	for _, block := range blocks {
		if len(fileData) == 0 {
			break
		}

		dataSize := block.size
		if dataSize > len(fileData) {
			dataSize = len(fileData)
		}

		if block.blockType == File {
			disk = append(disk, slice.Splice(&fileData, 0, dataSize)...)
		}
		if block.blockType == FreeSpace {
			disk = append(disk, slice.PopN(&fileData, dataSize)...)
		}
	}

	checksum := calculateChecksum(disk)
	println(checksum)
}
