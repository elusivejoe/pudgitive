package fileio

import (
	"fmt"
	"math"

	"github.com/recoilme/pudge"
)

type Writer struct {
	db           *pudge.Db
	chunkSize    int
	fileId       int
	chunkId      int
	currentChunk *Chunk
}

func NewWriter(db *pudge.Db, fileId, chunkSize int) (*Writer, error) {
	return &Writer{db: db, fileId: fileId, chunkSize: chunkSize}, nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.currentChunk == nil {
		w.currentChunk = &Chunk{Payload: nil, HasNext: false}
	}

	for len(p) > 0 {
		curChunkCap := w.chunkSize - len(w.currentChunk.Payload)
		copied := int(math.Min(float64(len(p)), float64(curChunkCap)))

		w.currentChunk.Payload = append(w.currentChunk.Payload, p[:copied]...)
		p = p[copied:]
		w.currentChunk.HasNext = len(p) > 0

		curChunkId := fmt.Sprintf("%d:%d", w.fileId, w.chunkId)

		if err := w.db.Set(curChunkId, w.currentChunk); err != nil {
			return n, err
		}

		if w.currentChunk.HasNext {
			w.currentChunk = &Chunk{Payload: nil, HasNext: false}
			w.chunkId++
		}

		n = copied
	}

	return n, nil
}
