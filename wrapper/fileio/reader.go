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
	chunk       *Chunk
	chunkOffset int
}

func NewReader(db *pudge.Db, fileId int) (*Reader, error) {
	return &Reader{db: db, fileId: fileId, chunk: &Chunk{}}, nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	for !r.eof && n < len(p) {
		if r.chunkOffset == len(r.chunk.Payload) {
			*r.chunk = Chunk{}
			chunkId := fmt.Sprintf("%d:%d", r.fileId, r.chunkId)

			if err := r.db.Get(chunkId, r.chunk); err != nil {
				return n, err
			}

			r.chunkId++
			r.chunkOffset = 0
		}

		copied := copy(p[n:], r.chunk.Payload[r.chunkOffset:])
		n += copied

		r.chunkOffset += copied
		r.eof = !r.chunk.HasNext && r.chunkOffset == len(r.chunk.Payload)
	}

	return n, nil
}

func (r *Reader) Eof() bool {
	return r.eof
}
