package wrapper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) MkDir(path string) (meta.Meta, error) {
	if len(path) == 0 {
		return meta.Meta{}, errors.New("wrapper: empty path provided")
	}

	parts := strings.Split(path, "/")

	ok, err := w.Exists(parts[0])

	if err != nil {
		return meta.Meta{}, err
	}

	if ok {
		return meta.Meta{}, fmt.Errorf("wrapper: dir '%s' already exists", path)
	}

	currentPos := w.root

	if relative := !strings.HasPrefix(path, "/"); relative {
		if len(w.curPosRel) > 0 {
			currentPos += w.curPosRel
		}
	}

	for _, part := range parts {
		currentPos += "/" + part
		w.db.Set(currentPos, meta.NewMeta(part, true))
	}

	return meta.NewMeta(parts[0], true), nil
}

func (w *Wrapper) MkFile(path string) (meta.Meta, error) {
	fmt.Printf("MkFile: %s\n", path)
	return meta.NewMeta(path, false), nil
}

func (w *Wrapper) RmDir(path string, recursive bool) error {
	fmt.Printf("RmDir: %s recursive: %t\n", path, recursive)
	return nil
}

func (w *Wrapper) RmFile(path string) error {
	fmt.Printf("RmFile: %s\n", path)
	return nil
}

func (w *Wrapper) Mv(src string, dst string) error {
	fmt.Printf("Mv: from %s to: %s\n", src, dst)
	return nil
}
