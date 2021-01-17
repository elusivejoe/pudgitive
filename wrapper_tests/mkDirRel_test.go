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
