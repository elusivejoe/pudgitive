package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"
)

func assembleEndpoint(w *Wrapper, path *pathUtils.NormPath) (string, error) {
	endpoint := w.root

	if !path.IsAbs() && len(w.where) > 0 {
		endpoint += "/" + w.where
	}

	endpoint += "/" + path.Path()

	validated, err := pathUtils.NewNormPath(endpoint)

	if err != nil {
		return "", err
	}

	return validated.Path(), nil
}

func trimPosition(w *Wrapper, path string) string {
	prefix := w.root

	if len(w.where) > 0 {
		prefix += "/" + w.where
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
