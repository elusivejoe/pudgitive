package wrapper

import (
	"errors"
	"fmt"
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
		return fmt.Errorf("wrapper: root '%s' already exists", rootName)
	}

	w.db.Set(rootName, rootMeta{Name: rootName, Magic: "pudgitive"})

	return nil
}

func (w *Wrapper) verifyRoot(rootName string) error {
	ok, err := w.db.Has(rootName)

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("wrapper: unable to find root '%s'", rootName)
	}

	root := rootMeta{}
	w.db.Get(rootName, &root)

	if root.Magic != "pudgitive" || root.Name != rootName {
		return fmt.Errorf("wrapper: broken root '%s' %v", rootName, root)
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

	w.db.Delete(key)

	return nil
}
