package wrapper

import (
	"github.com/elusivejoe/pudgitive/pathUtils"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) MkDir(path string) ([]Descriptor, error) {
	navPath, err := pathUtils.NewNavPath(resolveAbsolute(w, pathUtils.NewNormPath(path)))

	if err != nil {
		return nil, err
	}

	var descriptors []Descriptor

	for _, path := range navPath.DestList() {
		subDescriptors, err := w.mkDir(path)

		for _, desc := range subDescriptors {
			descriptors = append(descriptors, desc)
		}

		if err != nil {
			return descriptors, err
		}
	}

	return descriptors, nil
}

func (w *Wrapper) mkDir(path *pathUtils.NormPath) ([]Descriptor, error) {
	currentPos := w.root

	var descriptors []Descriptor

	for _, part := range path.Parts() {
		currentPos += "/" + part

		exists, err := w.db.Has(currentPos)

		if err != nil {
			return descriptors, err
		}

		if exists {
			continue
		}

		meta := meta.NewMeta(part, true)

		if err := w.db.Set(currentPos, meta); err != nil {
			return descriptors, err
		}

		descriptors = append(descriptors, Descriptor{currentPos, meta})
	}

	return descriptors, nil
}
