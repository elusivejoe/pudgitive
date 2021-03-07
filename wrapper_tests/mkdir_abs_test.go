package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"

	"github.com/stretchr/testify/assert"
)

func TestMkDirAbs(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

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

	metas, err := wrapper.MkDir("/dir_1/dir_2/dir_3")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(metas))

	for _, meta := range metas {
		assert.True(t, meta.IsDir)
	}

	assert.Equal(t, "dir_1", metas[0].Name)
	assert.Equal(t, "dir_2", metas[1].Name)
	assert.Equal(t, "dir_3", metas[2].Name)

	metas, err = wrapper.MkDir("/dir_1/dir_2/dir_4")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(metas))
	assert.True(t, metas[0].IsDir)
	assert.Equal(t, "dir_4", metas[0].Name)

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
	wrapper, _ := testutils.NewWrapper(t)

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

	metas, err := wrapper.MkDir("/a/b/c/../d/e")
	assert.Nil(t, err)
	assert.NotNil(t, metas)

	for _, meta := range metas {
		assert.True(t, meta.IsDir)
	}

	assert.Equal(t, "a", metas[0].Name)
	assert.Equal(t, "b", metas[1].Name)
	assert.Equal(t, "c", metas[2].Name)
	assert.Equal(t, "d", metas[3].Name)
	assert.Equal(t, "e", metas[4].Name)

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
