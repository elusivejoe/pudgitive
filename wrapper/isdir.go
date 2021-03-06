package wrapper

import "github.com/elusivejoe/pudgitive/utils"

func (w *Wrapper) IsDir(path string) (bool, error) {
	normPath := utils.NewNormPath(path)
	navPath, err := utils.NewNavPath(utils.ResolveAbsolute(w.pwd, normPath))

	if err != nil {
		return false, err
	}

	endpoint := utils.NewNormPath(w.root + "/" + navPath.FinalDest().Path()).Path()

	metaInfo, err := utils.ReadMeta(w.db, endpoint)

	if err != nil {
		return false, err
	}

	return metaInfo.Attrs.IsDir, nil
}
