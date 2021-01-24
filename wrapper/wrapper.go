package wrapper

import (
	"github.com/recoilme/pudge"
)

type Wrapper struct {
	db    *pudge.Db
	root  string
	where string
}

//TODO: allow user provide path-restriction func
func NewWrapper(db *pudge.Db) *Wrapper {
	return &Wrapper{db: db}
}
