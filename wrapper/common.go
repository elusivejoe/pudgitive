package wrapper

import (
	"strings"
)

func assembleEndpoint(w *Wrapper, path *checkedPath) (string, error) {
	endpoint := w.root

	if !path.IsAbs() && len(w.curPosRel) > 0 {
		endpoint += "/" + w.curPosRel
	}

	endpoint += "/" + path.Path()

	validated, err := NewCheckedPath(endpoint)

	if err != nil {
		return "", err
	}

	return validated.Path(), nil
}

func trimPosition(w *Wrapper, path string) string {
	prefix := w.root

	if len(w.curPosRel) > 0 {
		prefix += "/" + w.curPosRel
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
