package wrapper

import (
	"fmt"
	"strings"

	"github.com/elusivejoe/pudgitive/utils"
)

func (w *Wrapper) Cd(path string) error {
	navPath, err := utils.NewNavPath(resolveAbsolute(w, utils.NewNormPath(path)))

	if err != nil {
		return err
	}

	pathNorm := navPath.FinalDest()
	pwd := ""

	if pathNorm.Path() != "/" {
		pwd = pathNorm.Path()
	}

	exists, err := w.Exists(pwd)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("path '%s' does not exist", path)
	}

	w.pwd = strings.TrimPrefix(pwd, "/")

	return nil
}
