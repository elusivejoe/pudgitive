package wrapper_tests

import (
	"path"
	"testing"

	"github.com/elusivejoe/pudgitive/wrapper"

	"github.com/elusivejoe/pudgitive/database"
)

func createWrapper(t *testing.T) *wrapper.Wrapper {
	database := database.NewDatabase(path.Join(t.TempDir(), "testdb"))

	t.Cleanup(func() {
		database.Close()
	})

	return wrapper.NewWrapper(database)
}
