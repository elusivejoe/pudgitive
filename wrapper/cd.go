package wrapper

import (
	"fmt"
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"
)

func (w *Wrapper) Cd(path string) error {
	navPath, err := pathUtils.NewNavPath(resolveAbsolute(w, pathUtils.NewNormPath(path)))

	if err != nil {
		return err
	}

	pathNorm := navPath.FinalDest()
	where := ""

	if pathNorm.IsAbs() {
		where = strings.TrimPrefix(pathNorm.Path(), "/")
	} else {
		if len(w.where) == 0 {
			where = pathNorm.Path()
		} else {
			where = w.where + "/" + pathNorm.Path()
		}
	}

	exists, err := w.Exists(where)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("path '%s' does not exist", path)
	}

	w.where = where

	return nil
}
