package wrapper

import "github.com/elusivejoe/pudgitive/utils"

func (w *Wrapper) Exists(path string) (bool, error) {
	normAbs := utils.ResolveAbsolute(w.pwd, utils.NewNormPath(path))
	navPath, err := utils.NewNavPath(utils.ResolveAbsolute(w.pwd, normAbs))

	if err != nil {
		return false, err
	}

	endpoint := utils.NewNormPath(w.root + "/" + navPath.FinalDest().Path()).Path()

	ok, err := w.db.Has(endpoint)

	if err != nil {
		return false, err
	}

	return ok, nil
}
