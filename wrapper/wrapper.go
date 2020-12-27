package wrapper

import (
	"github.com/recoilme/pudge"
)

type Wrapper struct {
	db *pudge.Db
}

func NewWrapper(db *pudge.Db) *Wrapper {
	return &Wrapper{db: db}
}
