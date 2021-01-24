package pathUtils_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/pathUtils"

	"github.com/stretchr/testify/assert"
)

func TestNormPath(t *testing.T) {
	empty := ""
	path, err := pathUtils.NewNormPath(empty)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.False(t, path.IsAbs())
	assert.Equal(t, "", path.Path())
	assert.Equal(t, 1, len(path.Parts()))
	assert.Equal(t, "", path.Parts()[0])

	curious := "//////folder1/////folll\\der2////Spa   Ces//"
	path, err = pathUtils.NewNormPath(curious)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/folder1/folll\\der2/Spa   Ces", path.Path())
	assert.Equal(t, 3, len(path.Parts()))
	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folll\\der2", path.Parts()[1])
	assert.Equal(t, "Spa   Ces", path.Parts()[2])

	regularAbs := "/folder1/folder2/3folder"
	path, err = pathUtils.NewNormPath(regularAbs)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, regularAbs, path.Path())
	assert.Equal(t, 3, len(path.Parts()))
	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
	assert.Equal(t, "3folder", path.Parts()[2])

	normalRel := "folder1/folder2"
	path, err = pathUtils.NewNormPath(normalRel)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.False(t, path.IsAbs())
	assert.Equal(t, normalRel, path.Path())
	assert.Equal(t, 2, len(path.Parts()))
	assert.Equal(t, "folder1", path.Parts()[0])
	assert.Equal(t, "folder2", path.Parts()[1])
}

func TestManyDots(t *testing.T) {
	withManyDots := "/././././././.."
	path, err := pathUtils.NewNormPath(withManyDots)
	assert.Nil(t, path)
	assert.EqualError(t, err, "cannot go higher than root")

	withManyDots = "./././././../ab"
	path, err = pathUtils.NewNormPath(withManyDots)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.Equal(t, "../ab", path.Path())
	assert.Equal(t, 2, len(path.Parts()))
	assert.Equal(t, "..", path.Parts()[0])
	assert.Equal(t, "ab", path.Parts()[1])

	withManyDots = "/./..a a./../.././..a a../......./.a/a./a"
	path, err = pathUtils.NewNormPath(withManyDots)
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/..a a./../../..a a../......./.a/a./a", path.Path())
	assert.Equal(t, 8, len(path.Parts()))
	assert.Equal(t, "..a a.", path.Parts()[0])
	assert.Equal(t, "..", path.Parts()[1])
	assert.Equal(t, "..", path.Parts()[2])
	assert.Equal(t, "..a a..", path.Parts()[3])
	assert.Equal(t, ".......", path.Parts()[4])
	assert.Equal(t, ".a", path.Parts()[5])
	assert.Equal(t, "a.", path.Parts()[6])
	assert.Equal(t, "a", path.Parts()[7])
}

func TestNormRootOnly(t *testing.T) {
	path, err := pathUtils.NewNormPath("///////")
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/", path.Path())
	assert.Equal(t, 0, len(path.Parts()))

	path, err = pathUtils.NewNormPath("/")
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/", path.Path())
	assert.Equal(t, 0, len(path.Parts()))

	path, err = pathUtils.NewNormPath("/.")
	assert.Nil(t, err)
	assert.NotNil(t, path)
	assert.True(t, path.IsAbs())
	assert.Equal(t, "/", path.Path())
	assert.Equal(t, 0, len(path.Parts()))

	path, err = pathUtils.NewNormPath("/..")
	assert.Nil(t, path)
	assert.EqualError(t, err, "cannot go higher than root")
}
