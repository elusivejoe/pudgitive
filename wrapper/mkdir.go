package wrapper

import (
	"github.com/elusivejoe/pudgitive/utils"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) MkDir(path string) ([]meta.Meta, error) {
	navPath, err := utils.NewNavPath(utils.ResolveAbsolute(w.pwd, utils.NewNormPath(path)))

	if err != nil {
		return nil, err
	}

	var metas []meta.Meta

	for _, path := range navPath.DestList() {
		subMetas, err := w.mkDir(path)

		for _, meta := range subMetas {
			metas = append(metas, meta)
		}

		if err != nil {
			return metas, err
		}
	}

	return metas, nil
}

func (w *Wrapper) mkDir(path *utils.NormPath) ([]meta.Meta, error) {
	currentPos := w.root

	var metas []meta.Meta

	for _, part := range path.Parts() {
		currentPos += "/" + part

		exists, err := w.db.Has(currentPos)

		if err != nil {
			return metas, err
		}

		if exists {
			continue
		}

		meta := meta.NewMeta(part, true)

		if err := w.db.Set(currentPos, meta); err != nil {
			return metas, err
		}

		metas = append(metas, meta)
	}

	return metas, nil
}
