package wrapper_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCd(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_cd"))
	assert.Nil(t, wrapper.OpenRoot("test_cd"))

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_1/dir_2")
	wrapper.MkDir("/dir_1/dir_3")

	descriptors, err := wrapper.Ls("/", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, descriptors)
	assert.Equal(t, 3, len(descriptors))

	assert.Equal(t, "/dir_1", descriptors[0].Path)
	assert.Equal(t, "dir_1", descriptors[0].Meta.Name)
	assert.True(t, descriptors[0].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_2", descriptors[1].Path)
	assert.Equal(t, "dir_2", descriptors[1].Meta.Name)
	assert.True(t, descriptors[1].Meta.Attrs.IsDir)

	assert.Equal(t, "/dir_3", descriptors[2].Path)
	assert.Equal(t, "dir_3", descriptors[2].Meta.Name)
	assert.True(t, descriptors[2].Meta.Attrs.IsDir)

	err = wrapper.Cd("dir_1")
	assert.Nil(t, err)
}
