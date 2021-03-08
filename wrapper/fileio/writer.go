package fileio

import "github.com/recoilme/pudge"

type Writer struct {
	db        *pudge.Db
	fileId    int
	chunkSize int
	chunkId   int
	offset    int
}

func NewWriter(db *pudge.Db, fileId, chunkSize int) (*Writer, error) {
	return &Writer{db: db, fileId: fileId, chunkSize: chunkSize}, nil
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return 0, err
}
