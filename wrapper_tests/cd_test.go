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

	metas, err := wrapper.Ls("/", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 3, len(metas))

	assert.Equal(t, "dir_1", metas[0].Name)
	assert.True(t, metas[0].Attrs.IsDir)

	assert.Equal(t, "dir_2", metas[1].Name)
	assert.True(t, metas[1].Attrs.IsDir)

	assert.Equal(t, "dir_3", metas[2].Name)
	assert.True(t, metas[2].Attrs.IsDir)

	err = wrapper.Cd("dir_1")
	assert.Nil(t, err)
}
