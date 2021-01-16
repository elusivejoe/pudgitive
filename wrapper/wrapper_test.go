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

	ok, err = wrapper.Exists("test dir 2")
	assert.Nil(t, err)
	assert.False(t, ok)

	descriptors, err := wrapper.MkDir("test dir")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(descriptors))
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)
	assert.Equal(t, "test dir", descriptors[0].Meta.Name)

	ok, err = wrapper.Exists("test dir")
	assert.Nil(t, err)
	assert.True(t, ok)

	descriptors, err = wrapper.MkDir("test dir 2")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(descriptors))
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)
	assert.Equal(t, "test dir 2", descriptors[0].Meta.Name)

	ok, err = wrapper.Exists("test dir 2")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestMkDirAbs(t *testing.T) {
	wrapper := getWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_mkdir"))
	assert.Nil(t, wrapper.OpenRoot("test_mkdir"))

	ok, err := wrapper.Exists("dir_1")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("dir_1/dir_2")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("dir_1/dir_2/dir_3")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("dir_1/dir_2/dir_4")
	assert.Nil(t, err)
	assert.False(t, ok)

	descriptors, err := wrapper.MkDir("/dir_1/dir_2/dir_3")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(descriptors))

	for _, descriptor := range descriptors {
		assert.True(t, descriptor.Meta.Attrs.IsDir)
	}

	assert.Equal(t, "dir_1", descriptors[0].Meta.Name)
	assert.Equal(t, "dir_2", descriptors[1].Meta.Name)
	assert.Equal(t, "dir_3", descriptors[2].Meta.Name)

	descriptors, err = wrapper.MkDir("/dir_1/dir_2/dir_4")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(descriptors))
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)
	assert.Equal(t, "dir_4", descriptors[0].Meta.Name)

	ok, err = wrapper.Exists("/dir_1")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("/dir_1/dir_2")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("/dir_1/dir_2/dir_3")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("/dir_1/dir_2/dir_4")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestPath(t *testing.T) {
	empty := ""
	path, err := NewCheckedPath(empty)
	assert.Nil(t, path)
	assert.EqualError(t, err, "empty path")

	slashesOnly := "///////"
	path, err = NewCheckedPath(slashesOnly)
	assert.Nil(t, path)
	assert.EqualError(t, err, "empty path")

	curious := "//////folder1/////folll\\der2////Spa   Ces//"
	path, err = NewCheckedPath(curious)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/folder1/folll\\der2/Spa   Ces", path.Path())
	assert.Equal(t, 3, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folll\\der2", path.Parts()[1])
	assert.Equal(t, "Spa   Ces", path.Parts()[2])

	normalAbs := "/folder1/folder2/3folder"
	path, err = NewCheckedPath(normalAbs)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, normalAbs, path.Path())
	assert.Equal(t, 3, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
	assert.Equal(t, "3folder", path.Parts()[2])

	normalRel := "folder1/folder2"
	path, err = NewCheckedPath(normalRel)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.False(t, path.IsAbs())
	assert.Equal(t, normalRel, path.Path())
	assert.Equal(t, 2, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
}
