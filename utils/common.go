package utils

import (
	"strings"
)

func ResolveAbsolute(pwd string, path *NormPath) *NormPath {
	if !path.IsAbs() && len(pwd) > 0 {
		return NewNormPath("/" + pwd + "/" + path.Path())
	}

	return path
}

func TrimPosition(root, pwd string, path string, isAbs bool) string {
	prefix := root

	if !isAbs && len(pwd) > 0 {
		prefix += "/" + pwd
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
