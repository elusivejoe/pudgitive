package fileio

type Chunk struct {
	payload   []byte
	nextChunk string
}
