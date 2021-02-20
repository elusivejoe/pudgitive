package wrapper

import "fmt"

func (w *Wrapper) Mv(src string, dst string) error {
	fmt.Printf("Mv: from %s to: %s\n", src, dst)
	return nil
}
