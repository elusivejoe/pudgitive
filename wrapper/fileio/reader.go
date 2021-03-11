package fileio

import (
	"fmt"

	"github.com/recoilme/pudge"
)

type Reader struct {
	db          *pudge.Db
	eof         bool
	fileId      int
	chunkId     int
	chunkOffset int
}

func NewReader(db *pudge.Db, fileId int) (*Reader, error) {
	return &Reader{db: db, fileId: fileId}, nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	outCap := cap(p)

	if outCap == 0 || r.eof {
		return 0, nil
	}

	chunk := &Chunk{}

	for requestNextChunk := true; requestNextChunk; {
		curChunkId := fmt.Sprintf("%d:%d", r.fileId, r.chunkId)

		if err := r.db.Get(curChunkId, chunk); err != nil {
			return n, err
		}

		copied := copy(p[n:], chunk.Payload[r.chunkOffset:])
		n += copied
		requestNextChunk = chunk.HasNext && n < outCap

		if requestNextChunk {
			r.chunkId++
			r.chunkOffset = 0
			chunk.HasNext = false
		} else {
			r.chunkOffset += copied
		}

		r.eof = !chunk.HasNext && r.chunkOffset == len(chunk.Payload)
	}

	return n, nil
}

func (r *Reader) Eof() bool {
	return r.eof
}
