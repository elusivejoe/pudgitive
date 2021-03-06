package testutils

import (
	"path"
	"testing"

	"github.com/recoilme/pudge"

	"github.com/elusivejoe/pudgitive/wrapper"

	"github.com/elusivejoe/pudgitive/database"
)

func NewWrapper(t *testing.T) (*wrapper.Wrapper, *pudge.Db) {
	database := database.NewDatabase(path.Join(t.TempDir(), "testdb"))

	t.Cleanup(func() {
		database.Close()
	})

	wrapper := wrapper.NewWrapper(database)
	rootName := "test_root"
	err := wrapper.InitRoot(rootName)

	if err != nil {
		panic("failed to initialize root filesystem")
	}

	err = wrapper.OpenRoot(rootName)

	if err != nil {
		panic("failed to open root filesystem")
	}

	return wrapper, database
}
