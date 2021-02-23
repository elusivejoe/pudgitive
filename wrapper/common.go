package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/utils"
)

func resolveAbsolute(w *Wrapper, path *utils.NormPath) *utils.NormPath {
	if !path.IsAbs() && len(w.pwd) > 0 {
		return utils.NewNormPath("/" + w.pwd + "/" + path.Path())
	}

	return path
}

func trimPosition(w *Wrapper, path string, isAbs bool) string {
	prefix := w.root

	if !isAbs && len(w.pwd) > 0 {
		prefix += "/" + w.pwd
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
