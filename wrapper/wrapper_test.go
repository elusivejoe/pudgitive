package wrapper

import (
	"path"
	"testing"

	"github.com/elusivejoe/pudgitive/database"

	"github.com/stretchr/testify/assert"
)

func getWrapper(t *testing.T) *Wrapper {
	database := database.NewDatabase(path.Join(t.TempDir(), "testdb"))

	t.Cleanup(func() {
		database.Close()
	})

	return NewWrapper(database)
}

func TestAdmin(t *testing.T) {
	wrapper := getWrapper(t)

	err := wrapper.InitRoot("")
	assert.EqualError(t, err, "root name cannot be empty")

	rootName := "Awesome File System!"

	err = wrapper.InitRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")

	err = wrapper.InitRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")

	err = wrapper.InitRoot(rootName)

	assert.EqualError(t, err, "wrapper: root 'Awesome File System!' already exists")

	err = wrapper.OpenRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot("Some Other Root")
	assert.EqualError(t, err, "wrapper: unable to find root 'Some Other Root'")
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")
}

func TestMkDir(t *testing.T) {
	wrapper := getWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_mkdir"))
	assert.Nil(t, wrapper.OpenRoot("test_mkdir"))

	ok, err := wrapper.Exists("test dir")
	assert.Nil(t, err)
	assert.False(t, ok)

	meta, err := wrapper.MkDir("test dir")
	assert.Nil(t, err)
	assert.True(t, meta.Attributes().IsDirectory())
	assert.Equal(t, meta.EntityName(), "test dir")

	ok, err = wrapper.Exists("test dir")
	assert.Nil(t, err)
	assert.True(t, ok)
}
