package day_9

type BlockType string

const (
	File      BlockType = "file"
	FreeSpace BlockType = "free"
)

type Block struct {
	blockType BlockType
	size      int
	offset    int
	id        int
}
