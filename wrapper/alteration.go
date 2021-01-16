package wrapper

import (
	"fmt"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) MkDir(path string) ([]Descriptor, error) {
	pathChecked, err := NewCheckedPath(path)

	if err != nil {
		return nil, err
	}

	currentPos := w.root

	if !pathChecked.IsAbs() && len(w.curPosRel) > 0 {
		currentPos += w.curPosRel
	}

	var descriptors []Descriptor

	for _, part := range pathChecked.Parts() {
		currentPos += "/" + part

		exists, err := w.db.Has(currentPos)

		if err != nil {
			return nil, err
		}

		if exists {
			continue
		}

		meta := meta.NewMeta(part, true)

		if err := w.db.Set(currentPos, meta); err != nil {
			return nil, err
		}

		descriptors = append(descriptors, Descriptor{currentPos, meta})
	}

	return descriptors, nil
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
