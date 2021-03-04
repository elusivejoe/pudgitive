package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/utils"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) Ls(path string, limit, offset int, asc bool) ([]meta.Meta, error) {
	normPath := utils.NewNormPath(path)
	navPath, err := utils.NewNavPath(utils.ResolveAbsolute(w.pwd, normPath))

	if err != nil {
		return nil, err
	}

	endpoint := utils.NewNormPath(w.root + "/" + navPath.FinalDest().Path()).Path()

	prefixedKeys, err := w.db.KeysByPrefix([]byte(endpoint), 0, 0, asc)

	if err != nil {
		return nil, err
	}

	var metas []meta.Meta
	metaInfo := &meta.Meta{}

	currentOffset := 0

	for _, key := range prefixedKeys {
		subPath := utils.TrimDbKey(w.root, w.pwd, string(key), normPath.IsAbs())

		if isRootElem := len(subPath) == 0; isRootElem {
			continue
		}

		if isNotCurrentLevel := strings.Count(subPath, "/") > 1; isNotCurrentLevel {
			continue
		}

		if offset > 0 && currentOffset != offset {
			currentOffset++
			continue
		}

		key := endpoint + subPath

		if err := w.db.Get(key, metaInfo); err != nil {
			return metas, err
		}

		metas = append(metas, *metaInfo)

		if limit > 0 && len(metas) == limit {
			break
		}
	}

	return metas, nil
}
