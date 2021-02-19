package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"
)

func resolveAbsolute(w *Wrapper, path *pathUtils.NormPath) *pathUtils.NormPath {
	if !path.IsAbs() && len(w.where) > 0 {
		return pathUtils.NewNormPath("/" + w.where + "/" + path.Path())
	}

	return path
}

func trimPosition(w *Wrapper, path string, isAbs bool) string {
	prefix := w.root

	if !isAbs && len(w.where) > 0 {
		prefix += "/" + w.where
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
