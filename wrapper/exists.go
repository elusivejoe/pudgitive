package wrapper

import "github.com/elusivejoe/pudgitive/pathUtils"

func (w *Wrapper) Exists(path string) (bool, error) {
	navPath, err := pathUtils.NewNavPath(resolveAbsolute(w, pathUtils.NewNormPath(path)))

	if err != nil {
		return false, err
	}

	endpoint := pathUtils.NewNormPath(w.root + "/" + navPath.FinalDest().Path()).Path()

	ok, err := w.db.Has(endpoint)

	if err != nil {
		return false, err
	}

	return ok, nil
}
