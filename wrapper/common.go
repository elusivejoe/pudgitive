package wrapper

import (
	"strings"

	"github.com/elusivejoe/pudgitive/utils"
)

func resolveAbsolute(pwd string, path *utils.NormPath) *utils.NormPath {
	if !path.IsAbs() && len(pwd) > 0 {
		return utils.NewNormPath("/" + pwd + "/" + path.Path())
	}

	return path
}

func trimPosition(root, pwd string, path string, isAbs bool) string {
	prefix := root

	if !isAbs && len(pwd) > 0 {
		prefix += "/" + pwd
	}

	path = strings.TrimPrefix(path, prefix)

	return path
}
