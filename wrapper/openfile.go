package wrapper

import (
	"fmt"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) OpenFile(path string) (meta.Meta, error) {
	fmt.Printf("OpenFile: %s\n", path)
	return meta.NewMeta(path, false), nil
}
