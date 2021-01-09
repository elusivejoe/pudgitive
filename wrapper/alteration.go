package wrapper

import (
	"fmt"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) MkDir(path string, parents bool) (meta.Meta, error) {
	fmt.Printf("MkDir: %s %t\n", path, parents)
	return meta.NewMeta(path, true), nil
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
