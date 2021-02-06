package wrapper_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkDirRel(t *testing.T) {
	wrapper := createWrapper(t)

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

func TestMkDirRelDots(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_mkdir"))
	assert.Nil(t, wrapper.OpenRoot("test_mkdir"))

	descriptors, err := wrapper.MkDir("test dir/another dir")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(descriptors))
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)
	assert.Equal(t, "test dir", descriptors[0].Meta.Name)
	assert.Equal(t, "another dir", descriptors[1].Meta.Name)

	err = wrapper.Cd("test dir/another dir")
	assert.Nil(t, err)

	descriptors, err = wrapper.MkDir("a/b/c/../../b/c/d/e/../../../../../x/c/v")
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
	assert.Equal(t, "x", descriptors[5].Meta.Name)
	assert.Equal(t, "c", descriptors[6].Meta.Name)
	assert.Equal(t, "v", descriptors[7].Meta.Name)

	ok, err := wrapper.Exists("a")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/c")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/c/d")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("a/b/c/d/e")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("x")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("x/c")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.Exists("x/c/v")
	assert.Nil(t, err)
	assert.True(t, ok)
}
