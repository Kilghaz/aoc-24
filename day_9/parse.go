package day_9

import "aoc24/io"

func parseBlocks(input string) []Block {
	blocks := make([]Block, 0)
	offset := 0
	fileId := 0
	for i, r := range []rune(input) {
		blockType := File
		if i%2 != 0 {
			blockType = FreeSpace
		}

		id := 0
		if blockType == File {
			id = fileId
			fileId++
		}

		size := io.ParseInt(string(r))

		blocks = append(blocks, Block{
			blockType: blockType,
			size:      size,
			offset:    offset,
			id:        id,
		})

		offset += size
	}
	return blocks
}
