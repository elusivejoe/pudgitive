package wrapper_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkDirAbs(t *testing.T) {
	wrapper := createWrapper(t)

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

func TestMkDirAbsDots(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_mkdir"))
	assert.Nil(t, wrapper.OpenRoot("test_mkdir"))

	ok, err := wrapper.Exists("a")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("a/b")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("a/b/c")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("a/b/d")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("a/b/d/e")
	assert.Nil(t, err)
	assert.False(t, ok)

	descriptors, err := wrapper.MkDir("/a/b/c/../d/e")
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)

	for _, descriptor := range descriptors {
		assert.True(t, descriptor.Meta.Attrs.IsDir)
	}

	assert.Equal(t, "a", descriptors[0].Meta.Name)
	assert.Equal(t, "b", descriptors[1].Meta.Name)
	assert.Equal(t, "c", descriptors[2].Meta.Name)
	assert.Equal(t, "d", descriptors[3].Meta.Name)
	assert.Equal(t, "e", descriptors[4].Meta.Name)

	ok, err = wrapper.Exists("a")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/c")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/d")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/d/e")
	assert.Nil(t, err)
	assert.True(t, ok)
}
