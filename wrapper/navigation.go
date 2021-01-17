package wrapper

import (
	"fmt"
	"strings"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) Ls(path string, limit, offset int, asc bool) ([]Descriptor, error) {
	pathChecked, err := NewCheckedPath(path)

	if err != nil {
		return nil, err
	}

	endpoint, err := assembleEndpoint(w, pathChecked)

	if err != nil {
		return nil, err
	}

	prefixedKeys, err := w.db.KeysByPrefix([]byte(endpoint), 0, 0, asc)

	if err != nil {
		return nil, err
	}

	var descriptors []Descriptor
	metaInfo := &meta.Meta{}

	currentOffset := 0

	for _, prefix := range prefixedKeys {
		subPath := trimPosition(w, string(prefix))

		if isRootElem := len(subPath) == 0; isRootElem {
			continue
		}

		if isNotCurrentLevel := strings.Count(subPath, "/") > 1; isNotCurrentLevel {
			continue
		}

		if offset > 0 && currentOffset != offset {
			currentOffset++
			continue
		}

		key := endpoint + subPath

		if err := w.db.Get(key, metaInfo); err != nil {
			return descriptors, err
		}

		pathNorm := trimPosition(w, key)

		descriptors = append(descriptors, Descriptor{Path: pathNorm, Meta: *metaInfo})

		if limit > 0 && len(descriptors) == limit {
			break
		}
	}

	return descriptors, nil
}

func (w *Wrapper) Cd(path string) error {
	fmt.Printf("Cd: %s\n", path)
	return nil
}

func (w *Wrapper) Exists(path string) (bool, error) {
	pathChecked, err := NewCheckedPath(path)

	if err != nil {
		return false, err
	}

	endpoint, err := assembleEndpoint(w, pathChecked)

	if err != nil {
		return false, err
	}

	ok, err := w.db.Has(endpoint)

	if err != nil {
		return false, err
	}

	return ok, nil
}

func (w *Wrapper) IsDir(path string) (bool, error) {
	fmt.Printf("IsDir: %s\n", path)
	return false, nil
}
