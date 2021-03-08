package fileio

import "github.com/recoilme/pudge"

type Reader struct {
	db     *pudge.Db
	fileId int
}

func (b *Reader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func NewReader(db *pudge.Db, fileId int) (*Reader, error) {
	return &Reader{db: db, fileId: fileId}, nil
}
