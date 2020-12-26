package wrapper

import (
	"github.com/elusivejoe/pudgitive/database"
)

type Wrapper struct {
	db *database.DummyDatabase
}

func NewWrapper(db *database.DummyDatabase) *Wrapper {
	return &Wrapper{db: db}
}
