package wrapper

import "fmt"

func (w *Wrapper) RmDir(path string, recursive bool) error {
	fmt.Printf("RmDir: %s recursive: %t\n", path, recursive)
	return nil
}
