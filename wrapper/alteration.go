package wrapper

import (
	"fmt"

	"github.com/elusivejoe/pudgitive/tree"
)

func (w *Wrapper) mkDir(path string, parents bool) (tree.Node, error) {
	fmt.Printf("mkDir: %s %t", path, parents)
	return tree.NewNode(path, true), nil
}

func (w *Wrapper) mkFile(path string) (tree.Node, error) {
	fmt.Printf("mkFile: %s", path)
	return tree.NewNode(path, false), nil
}

func (w *Wrapper) rmDir(path string, recursive bool) error {
	fmt.Printf("rmDir: %s recursive: %t\n", path, recursive)
	return nil
}

func (w *Wrapper) rmFile(path string) error {
	fmt.Printf("rmFile: %s\n", path)
	return nil
}

func (w *Wrapper) mv(src string, dst string) error {
	fmt.Printf("mv: from %s to: %s\n", src, dst)
	return nil
}
