package wrapper

import (
	"errors"
	"fmt"

	"github.com/elusivejoe/pudgitive/meta"
)

func (w *Wrapper) CurrentRoot() string {
	return w.root
}

func (w *Wrapper) InitRoot(rootName string) error {
	if len(rootName) == 0 {
		return errors.New("root name cannot be empty")
	}

	ok, err := w.db.Has(rootName)

	if err != nil {
		return err
	}

	if ok {
		return fmt.Errorf("root '%s' already exists", rootName)
	}

	return w.db.Set(rootName, meta.RootMeta{Name: rootName, Magic: "pudgitive"})
}

func (w *Wrapper) verifyRoot(rootName string) error {
	ok, err := w.db.Has(rootName)

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("unable to find root '%s'", rootName)
	}

	root := meta.RootMeta{}
	err = w.db.Get(rootName, &root)

	if err != nil {
		return err
	}

	if root.Magic != "pudgitive" || root.Name != rootName {
		return fmt.Errorf("broken root '%s' %v", rootName, root)
	}

	return nil
}

func (w *Wrapper) OpenRoot(key string) error {
	if err := w.verifyRoot(key); err != nil {
		return err
	}

	w.root = key

	return nil
}

func (w *Wrapper) DeleteRoot(key string) error {
	if err := w.verifyRoot(key); err != nil {
		return err
	}

	if w.root == key {
		w.root = ""
	}

	return w.db.Delete(key)
}
