package utils_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/utils"
	"github.com/stretchr/testify/assert"
)

func TestResolveAbsolute(t *testing.T) {
	path := utils.NewNormPath("/a/b/c/d")
	resolved := utils.ResolveAbsolute("/a/b", path)

	assert.True(t, resolved.IsAbs())
	assert.Equal(t, "/a/b/c/d", resolved.Path())

	path = utils.NewNormPath("c/d")
	resolved = utils.ResolveAbsolute("/a/b", path)

	assert.True(t, resolved.IsAbs())
	assert.Equal(t, "/a/b/c/d", resolved.Path())
}
