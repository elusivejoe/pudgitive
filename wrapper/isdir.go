package wrapper

import "fmt"

func (w *Wrapper) IsDir(path string) (bool, error) {
	fmt.Printf("IsDir: %s\n", path)
	return false, nil
}
