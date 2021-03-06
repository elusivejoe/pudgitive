package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"

	"github.com/stretchr/testify/assert"
)

func TestMkDirRel(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	ok, err := wrapper.Exists("test dir")
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = wrapper.Exists("test dir 2")
	assert.Nil(t, err)
	assert.False(t, ok)

	metas, err := wrapper.MkDir("test dir")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(metas))
	assert.True(t, metas[0].Attrs.IsDir)
	assert.Equal(t, "test dir", metas[0].Name)

	ok, err = wrapper.Exists("test dir")
	assert.Nil(t, err)
	assert.True(t, ok)

	metas, err = wrapper.MkDir("test dir 2")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(metas))
	assert.True(t, metas[0].Attrs.IsDir)
	assert.Equal(t, "test dir 2", metas[0].Name)

	ok, err = wrapper.Exists("test dir 2")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestMkDirRelDots(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	metas, err := wrapper.MkDir("test dir/another dir")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(metas))
	assert.True(t, metas[0].Attrs.IsDir)
	assert.True(t, metas[1].Attrs.IsDir)
	assert.Equal(t, "test dir", metas[0].Name)
	assert.Equal(t, "another dir", metas[1].Name)

	err = wrapper.Cd("test dir/another dir")
	assert.Nil(t, err)

	metas, err = wrapper.MkDir("a/b/c/../../b/c/d/e/../../../../../x/c/v")
	assert.Nil(t, err)
	assert.NotNil(t, metas)

	for _, meta := range metas {
		assert.True(t, meta.Attrs.IsDir)
	}

	assert.Equal(t, "a", metas[0].Name)
	assert.Equal(t, "b", metas[1].Name)
	assert.Equal(t, "c", metas[2].Name)
	assert.Equal(t, "d", metas[3].Name)
	assert.Equal(t, "e", metas[4].Name)
	assert.Equal(t, "x", metas[5].Name)
	assert.Equal(t, "c", metas[6].Name)
	assert.Equal(t, "v", metas[7].Name)

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
