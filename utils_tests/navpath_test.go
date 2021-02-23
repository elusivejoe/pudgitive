package utils_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elusivejoe/pudgitive/utils"
)

func TestNavPath(t *testing.T) {
	navPath, err := utils.NewNavPath(nil)
	assert.Nil(t, navPath)
	assert.EqualError(t, err, "nil path provided")

	navPath, err = utils.NewNavPath(utils.NewNormPath("/"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.DestList()))
	assert.True(t, navPath.DestList()[0].IsAbs())
	assert.Equal(t, "/", navPath.DestList()[0].Path())

	navPath, err = utils.NewNavPath(utils.NewNormPath(""))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.DestList()))
	assert.False(t, navPath.DestList()[0].IsAbs())
	assert.Equal(t, "", navPath.DestList()[0].Path())

	navPath, err = utils.NewNavPath(utils.NewNormPath("/ab/c/d"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.DestList()))
	assert.True(t, navPath.DestList()[0].IsAbs())
	assert.Equal(t, "/ab/c/d", navPath.DestList()[0].Path())

}

func TestNavPathGoUp(t *testing.T) {
	navPath, err := utils.NewNavPath(utils.NewNormPath("/.."))
	assert.Nil(t, navPath)
	assert.EqualError(t, err, "cannot go any higher")

	navPath, err = utils.NewNavPath(utils.NewNormPath("/ab/c/d/../e/f"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 2, len(navPath.DestList()))
	assert.True(t, navPath.DestList()[0].IsAbs())
	assert.True(t, navPath.DestList()[1].IsAbs())
	assert.Equal(t, "/ab/c/d", navPath.DestList()[0].Path())
	assert.Equal(t, "/ab/c/e/f", navPath.DestList()[1].Path())
	assert.Equal(t, "/ab/c/e/f", navPath.FinalDest().Path())
}

func TestNavPathGoUpOverlap(t *testing.T) {
	navPath, err := utils.NewNavPath(utils.NewNormPath("/a/b/c/d/../../../b/c/d"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.DestList()))
	assert.True(t, navPath.DestList()[0].IsAbs())
	assert.Equal(t, "/a/b/c/d", navPath.DestList()[0].Path())

	navPath, err = utils.NewNavPath(utils.NewNormPath("a/b/c/../../b/c/d/e/../../../../../x/c/v"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 2, len(navPath.DestList()))
	assert.False(t, navPath.DestList()[0].IsAbs())
	assert.False(t, navPath.DestList()[1].IsAbs())
	assert.Equal(t, "a/b/c/d/e", navPath.DestList()[0].Path())
	assert.Equal(t, "x/c/v", navPath.DestList()[1].Path())
	assert.Equal(t, "x/c/v", navPath.FinalDest().Path())
}
