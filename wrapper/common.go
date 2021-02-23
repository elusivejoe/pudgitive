package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"
)

func resolveAbsolute(w *Wrapper, path *pathUtils.NormPath) *pathUtils.NormPath {
	if !path.IsAbs() && len(w.pwd) > 0 {
		return pathUtils.NewNormPath("/" + w.pwd + "/" + path.Path())
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
