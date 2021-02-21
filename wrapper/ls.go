package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) Ls(path string, limit, offset int, asc bool) ([]meta.Meta, error) {
	normPath := pathUtils.NewNormPath(path)
	navPath, err := pathUtils.NewNavPath(resolveAbsolute(w, normPath))

	if err != nil {
		return nil, err
	}

	endpoint := pathUtils.NewNormPath(w.root + "/" + navPath.FinalDest().Path()).Path()

	prefixedKeys, err := w.db.KeysByPrefix([]byte(endpoint), 0, 0, asc)

	if err != nil {
		return nil, err
	}

	var metas []meta.Meta
	metaInfo := &meta.Meta{}

	currentOffset := 0

	for _, prefix := range prefixedKeys {
		subPath := trimPosition(w, string(prefix), normPath.IsAbs())

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
