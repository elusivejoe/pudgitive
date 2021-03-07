package meta

import (
	"github.com/recoilme/pudge"
)

func ReadMeta(db *pudge.Db, key string) (Meta, error) {
	metaInfo := &Meta{}

	if err := db.Get(key, metaInfo); err != nil {
		return Meta{}, err
	}

	return *metaInfo, nil
}
