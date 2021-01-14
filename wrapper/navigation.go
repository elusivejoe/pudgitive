package wrapper

import (
	"fmt"
	"strings"
)

func (w *Wrapper) Ls(path string) ([]string, error) {
	fmt.Printf("Ls: %s\n", path)
	return []string{}, nil
}

func (w *Wrapper) Cd(path string) error {
	fmt.Printf("Cd: %s\n", path)
	return nil
}

func (w *Wrapper) Exists(path string) (bool, error) {
	endpoint := w.root

	if relative := !strings.HasPrefix(path, "/"); relative {
		if len(w.curPosRel) > 0 {
			endpoint += "/" + w.curPosRel
		}
	}

	if lastSlashIdx := strings.LastIndex(path, "/"); lastSlashIdx != -1 {
		endpoint += "/" + path[:lastSlashIdx]
	} else {
		endpoint += "/" + path
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
