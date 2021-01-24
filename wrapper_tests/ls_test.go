package wrapper_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLsDefaults(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_ls"))
	assert.Nil(t, wrapper.OpenRoot("test_ls"))

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	descriptors, err := wrapper.Ls("/", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 4, len(descriptors))

	descriptors, err = wrapper.Ls(".", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 4, len(descriptors))

	descriptors, err = wrapper.Ls("", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 4, len(descriptors))

	assert.Equal(t, "/dir_1", descriptors[0].Path)
	assert.Equal(t, "dir_1", descriptors[0].Meta.Name)
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_2", descriptors[1].Path)
	assert.Equal(t, "dir_2", descriptors[1].Meta.Name)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_3", descriptors[2].Path)
	assert.Equal(t, "dir_3", descriptors[2].Meta.Name)
	assert.True(t, descriptors[2].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_4", descriptors[3].Path)
	assert.Equal(t, "dir_4", descriptors[3].Meta.Name)
	assert.True(t, descriptors[3].Meta.Attrs.IsDir)
}

func TestLsDesc(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_ls"))
	assert.Nil(t, wrapper.OpenRoot("test_ls"))

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	descriptors, err := wrapper.Ls("/", 0, 0, false)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 4, len(descriptors))

	assert.Equal(t, "/dir_4", descriptors[0].Path)
	assert.Equal(t, "dir_4", descriptors[0].Meta.Name)
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_3", descriptors[1].Path)
	assert.Equal(t, "dir_3", descriptors[1].Meta.Name)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_2", descriptors[2].Path)
	assert.Equal(t, "dir_2", descriptors[2].Meta.Name)
	assert.True(t, descriptors[2].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_1", descriptors[3].Path)
	assert.Equal(t, "dir_1", descriptors[3].Meta.Name)
	assert.True(t, descriptors[3].Meta.Attrs.IsDir)
}

func TestLsOffsLimAsc(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_ls"))
	assert.Nil(t, wrapper.OpenRoot("test_ls"))

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	descriptors, err := wrapper.Ls("/", 2, 2, true)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 2, len(descriptors))

	assert.Equal(t, "/dir_3", descriptors[0].Path)
	assert.Equal(t, "dir_3", descriptors[0].Meta.Name)
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_4", descriptors[1].Path)
	assert.Equal(t, "dir_4", descriptors[1].Meta.Name)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)
}

func TestLsOffsLimDesc(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_ls"))
	assert.Nil(t, wrapper.OpenRoot("test_ls"))

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	descriptors, err := wrapper.Ls("/", 2, 2, false)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 2, len(descriptors))

	assert.Equal(t, "/dir_2", descriptors[0].Path)
	assert.Equal(t, "dir_2", descriptors[0].Meta.Name)
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_1", descriptors[1].Path)
	assert.Equal(t, "dir_1", descriptors[1].Meta.Name)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)
}
