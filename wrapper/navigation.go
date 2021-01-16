package wrapper

import (
	"fmt"
)

func (w *Wrapper) assembleEndpoint(path *checkedPath) (*checkedPath, error) {
	endpoint := w.root

	if !path.IsAbs() && len(w.curPosRel) > 0 {
		endpoint += "/" + w.curPosRel
	}

	endpoint += "/" + path.Path()

	return NewCheckedPath(endpoint)
}

func (w *Wrapper) Ls(path string) ([]string, error) {
	fmt.Printf("Ls: %s\n", path)
	return []string{}, nil
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

	endpoint, err := w.assembleEndpoint(pathChecked)

	if err != nil {
		return false, err
	}

	ok, err := w.db.Has(endpoint.Path())

	if err != nil {
		return false, err
	}

	return ok, nil
}

func (w *Wrapper) IsDir(path string) (bool, error) {
	fmt.Printf("IsDir: %s\n", path)
	return false, nil
}
