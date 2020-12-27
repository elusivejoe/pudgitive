package wrapper

import (
	"fmt"

	"github.com/elusivejoe/pudgitive/tree"
)

func (w *Wrapper) ls(path string) ([]tree.Node, error) {
	fmt.Printf("ls: %s\n", path)
	return []tree.Node{}, nil
}

func (w *Wrapper) cd(path string) error {
	fmt.Printf("cd: %s\n", path)
	return nil
}

func (w *Wrapper) exists(path string) (bool, error) {
	fmt.Printf("exists: %s\n", path)
	return false, nil
}

func (w *Wrapper) isDir(path string) (bool, error) {
	fmt.Printf("isDir: %s\n", path)
	return false, nil
}
