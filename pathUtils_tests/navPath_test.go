package pathUtils_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elusivejoe/pudgitive/pathUtils"
)

func TestNavPath(t *testing.T) {
	navPath, err := pathUtils.NewNavPath(nil)
	assert.Nil(t, navPath)
	assert.EqualError(t, err, "nil path provided")

	navPath, err = pathUtils.NewNavPath(pathUtils.NewNormPath("/ab/c/d"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.AllDestinations()))
	assert.True(t, navPath.AllDestinations()[0].IsAbs())
	assert.Equal(t, "/ab/c/d", navPath.AllDestinations()[0].Path())

}

func TestNavPathGoUp(t *testing.T) {
	navPath, err := pathUtils.NewNavPath(pathUtils.NewNormPath("/.."))
	assert.Nil(t, navPath)
	assert.EqualError(t, err, "cannot go any higher")

	navPath, err = pathUtils.NewNavPath(pathUtils.NewNormPath("/ab/c/d/../e/f"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 2, len(navPath.AllDestinations()))
	assert.True(t, navPath.AllDestinations()[0].IsAbs())
	assert.True(t, navPath.AllDestinations()[1].IsAbs())
	assert.Equal(t, "/ab/c/d", navPath.AllDestinations()[0].Path())
	assert.Equal(t, "/ab/c/e/f", navPath.AllDestinations()[1].Path())
	assert.Equal(t, "/ab/c/e/f", navPath.FinalDestination().Path())
}

func TestNavPathGoUpOverlap(t *testing.T) {
	navPath, err := pathUtils.NewNavPath(pathUtils.NewNormPath("/a/b/c/d/../../../b/c/d"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 1, len(navPath.AllDestinations()))
	assert.True(t, navPath.AllDestinations()[0].IsAbs())
	assert.Equal(t, "/a/b/c/d", navPath.AllDestinations()[0].Path())

	navPath, err = pathUtils.NewNavPath(pathUtils.NewNormPath("a/b/c/../../b/c/d/e/../../../../../x/c/v"))
	assert.Nil(t, err)
	assert.NotNil(t, navPath)
	assert.Equal(t, 2, len(navPath.AllDestinations()))
	assert.False(t, navPath.AllDestinations()[0].IsAbs())
	assert.False(t, navPath.AllDestinations()[1].IsAbs())
	assert.Equal(t, "a/b/c/d/e", navPath.AllDestinations()[0].Path())
	assert.Equal(t, "x/c/v", navPath.AllDestinations()[1].Path())
	assert.Equal(t, "x/c/v", navPath.FinalDestination().Path())
}
