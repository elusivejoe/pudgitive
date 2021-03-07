package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"

	"github.com/stretchr/testify/assert"
)

func TestLsAbs(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	err := wrapper.Cd("/dir_4/dir_5/dir_6")
	assert.Nil(t, err)

	metas, err := wrapper.Ls("/", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 4, len(metas))

	metas, err = wrapper.Ls(".", 0, 0, true)
	assert.Nil(t, err)
	assert.Nil(t, metas)

	err = wrapper.Cd("/")
	assert.Nil(t, err)

	metas, err = wrapper.Ls(".", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 4, len(metas))

	metas, err = wrapper.Ls("", 0, 0, true)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 4, len(metas))

	assert.Equal(t, "dir_1", metas[0].Name)
	assert.True(t, metas[0].IsDir)

	assert.Equal(t, "dir_2", metas[1].Name)
	assert.True(t, metas[1].IsDir)

	assert.Equal(t, "dir_3", metas[2].Name)
	assert.True(t, metas[2].IsDir)

	assert.Equal(t, "dir_4", metas[3].Name)
	assert.True(t, metas[3].IsDir)
}

func TestLsDesc(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	metas, err := wrapper.Ls("/", 0, 0, false)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 4, len(metas))

	assert.Equal(t, "dir_4", metas[0].Name)
	assert.True(t, metas[0].IsDir)

	assert.Equal(t, "dir_3", metas[1].Name)
	assert.True(t, metas[1].IsDir)

	assert.Equal(t, "dir_2", metas[2].Name)
	assert.True(t, metas[2].IsDir)

	assert.Equal(t, "dir_1", metas[3].Name)
	assert.True(t, metas[3].IsDir)
}

func TestLsOffsLimAsc(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	metas, err := wrapper.Ls("/", 2, 2, true)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 2, len(metas))

	assert.Equal(t, "dir_3", metas[0].Name)
	assert.True(t, metas[0].IsDir)

	assert.Equal(t, "dir_4", metas[1].Name)
	assert.True(t, metas[1].IsDir)
}

func TestLsOffsLimDesc(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	wrapper.MkDir("/dir_1")
	wrapper.MkDir("/dir_2")
	wrapper.MkDir("/dir_3")
	wrapper.MkDir("/dir_4/dir_5")
	wrapper.MkDir("/dir_4/dir_5/dir_6")

	metas, err := wrapper.Ls("/", 2, 2, false)
	assert.Nil(t, err)
	assert.NotNil(t, metas)
	assert.Equal(t, 2, len(metas))

	assert.Equal(t, "dir_2", metas[0].Name)
	assert.True(t, metas[0].IsDir)

	assert.Equal(t, "dir_1", metas[1].Name)
	assert.True(t, metas[1].IsDir)
}
