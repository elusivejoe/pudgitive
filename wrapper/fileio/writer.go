package fileio

import (
	"fmt"
	"math"

	"github.com/recoilme/pudge"
)

type Writer struct {
	db        *pudge.Db
	chunkSize int
	fileId    int
	chunkId   int
	chunk     *Chunk
}

func NewWriter(db *pudge.Db, fileId, chunkSize int) (*Writer, error) {
	return &Writer{db: db, fileId: fileId, chunkSize: chunkSize, chunk: &Chunk{}}, nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	for len(p) > 0 {
		chunkCap := w.chunkSize - len(w.chunk.Payload)
		copied := int(math.Min(float64(len(p)), float64(chunkCap)))

		w.chunk.Payload = append(w.chunk.Payload, p[:copied]...)
		p = p[copied:]
		w.chunk.HasNext = len(p) > 0

		chunkId := fmt.Sprintf("%d:%d", w.fileId, w.chunkId)

		if err := w.db.Set(chunkId, w.chunk); err != nil {
			return n, err
		}

		if w.chunk.HasNext {
			*w.chunk = Chunk{}
			w.chunkId++
		}

		n = copied
	}

	return n, nil
}
