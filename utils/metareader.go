package utils

import (
	"github.com/elusivejoe/pudgitive/meta"
	"github.com/recoilme/pudge"
)

func ReadMeta(db *pudge.Db, key string) (meta.Meta, error) {
	metaInfo := &meta.Meta{}

	if err := db.Get(key, metaInfo); err != nil {
		return meta.Meta{}, err
	}

	return *metaInfo, nil
}
