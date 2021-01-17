package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/wrapper"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	empty := ""
	path, err := wrapper.NewCheckedPath(empty)
	assert.Nil(t, path)
	assert.EqualError(t, err, "empty path")

	slashesOnly := "///////"
	path, err = wrapper.NewCheckedPath(slashesOnly)
	assert.Nil(t, path)
	assert.EqualError(t, err, "empty path")

	curious := "//////folder1/////folll\\der2////Spa   Ces//"
	path, err = wrapper.NewCheckedPath(curious)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/folder1/folll\\der2/Spa   Ces", path.Path())
	assert.Equal(t, 3, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folll\\der2", path.Parts()[1])
	assert.Equal(t, "Spa   Ces", path.Parts()[2])

	normalAbs := "/folder1/folder2/3folder"
	path, err = wrapper.NewCheckedPath(normalAbs)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, normalAbs, path.Path())
	assert.Equal(t, 3, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
	assert.Equal(t, "3folder", path.Parts()[2])

	normalRel := "folder1/folder2"
	path, err = wrapper.NewCheckedPath(normalRel)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.False(t, path.IsAbs())
	assert.Equal(t, normalRel, path.Path())
	assert.Equal(t, 2, len(path.Parts()))

	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
}
