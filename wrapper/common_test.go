package wrapper

import (
	"path"
	"testing"

	"github.com/elusivejoe/pudgitive/database"

	"github.com/stretchr/testify/assert"
)

func createWrapper(t *testing.T) *Wrapper {
	database := database.NewDatabase(path.Join(t.TempDir(), "testdb"))

	t.Cleanup(func() {
		database.Close()
	})

	return NewWrapper(database)
}

func TestPositionTrimming(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_pos_trim"))
	assert.Nil(t, wrapper.OpenRoot("test_pos_trim"))

	wrapper.pwd = "a/b/c"
	assert.Equal(t, "/a/b/c", trimPosition(wrapper, "test_pos_trim/a/b/c", true))
	assert.Equal(t, "/d/e", trimPosition(wrapper, "test_pos_trim/a/b/c/d/e", false))
}
