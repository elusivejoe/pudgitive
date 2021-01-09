package wrapper

import (
	"fmt"
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
	fmt.Printf("Exists: %s\n", path)
	return false, nil
}

func (w *Wrapper) IsDir(path string) (bool, error) {
	fmt.Printf("IsDir: %s\n", path)
	return false, nil
}
