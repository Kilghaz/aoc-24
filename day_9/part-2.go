package day_9

import (
	"aoc24/slice"
	"github.com/samber/lo"
	"slices"
)

func Part2(input string) {
	blocks := parseBlocks(input)

	fileBlocks := lo.Filter(blocks, func(block Block, index int) bool {
		return block.blockType == File
	})

	spaceBlocks := lo.Filter(blocks, func(block Block, index int) bool {
		return block.blockType == FreeSpace
	})

	for fileIndex := len(fileBlocks) - 1; fileIndex >= 0; fileIndex-- {
		file := fileBlocks[fileIndex]
		for spaceIndex, space := range spaceBlocks {
			if space.offset > file.offset {
				continue
			}

			if file.size > space.size {
				continue
			}

			freedSpaceBlock := Block{
				blockType: FreeSpace,
				size:      file.size,
				offset:    file.offset,
			}
			spaceBlocks = append(spaceBlocks, freedSpaceBlock)

			file.offset = space.offset
			space.size -= file.size
			space.offset += file.size

			fileBlocks[fileIndex] = file
			spaceBlocks[spaceIndex] = space
			break
		}
	}

	nonEmptySpaceBlocks := lo.Filter(spaceBlocks, func(block Block, index int) bool {
		return block.size > 0
	})

	diskBlocks := make([]Block, 0)
	diskBlocks = append(diskBlocks, fileBlocks...)
	diskBlocks = append(diskBlocks, nonEmptySpaceBlocks...)
	slices.SortFunc(diskBlocks, func(a, b Block) int {
		return a.offset - b.offset
	})

	disk := lo.FlatMap(diskBlocks, func(block Block, _ int) []int {
		return slice.Fill(make([]int, block.size), block.id)
	})

	checksum := calculateChecksum(disk)
	println(checksum)
}
