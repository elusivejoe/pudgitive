package wrapper

import (
	"github.com/recoilme/pudge"
)

type Wrapper struct {
	db   *pudge.Db
	root string
	pwd  string
}

//TODO: allow user provide path-restriction func
//TODO: allow user provide chunk size for files
func NewWrapper(db *pudge.Db) *Wrapper {
	return &Wrapper{db: db}
}
