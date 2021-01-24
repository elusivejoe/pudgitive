package wrapper

import (
	"fmt"
	"strings"

	"github.com/elusivejoe/pudgitive/pathUtils"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) Ls(path string, limit, offset int, asc bool) ([]Descriptor, error) {
	pathNorm, err := pathUtils.NewNormPath(path)

	if err != nil {
		return nil, err
	}

	endpoint, err := assembleEndpoint(w, pathNorm)

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
	pathNorm, err := pathUtils.NewNormPath(path)

	if err != nil {
		return err
	}

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

func (w *Wrapper) Where() string {
	if len(w.where) > 0 {
		return w.where
	}

	return "/"
}

func (w *Wrapper) Exists(path string) (bool, error) {
	pathNorm, err := pathUtils.NewNormPath(path)

	if err != nil {
		return false, err
	}

	endpoint, err := assembleEndpoint(w, pathNorm)

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
