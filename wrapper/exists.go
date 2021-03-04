package wrapper

import "github.com/elusivejoe/pudgitive/utils"

func (w *Wrapper) Exists(path string) (bool, error) {
	navPath, err := utils.NewNavPath(resolveAbsolute(w.pwd, utils.NewNormPath(path)))

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
