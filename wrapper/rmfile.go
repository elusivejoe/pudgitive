package wrapper

import "fmt"

func (w *Wrapper) RmFile(path string) error {
	fmt.Printf("RmFile: %s\n", path)
	return nil
}
