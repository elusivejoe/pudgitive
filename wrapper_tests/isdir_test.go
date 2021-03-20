package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"
	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	ok, err := wrapper.IsDir("/")
	assert.Nil(t, err)
	assert.False(t, ok)

	wrapper.MkDir("/a/b/c/d/../../e")

	ok, err = wrapper.IsDir("a")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.IsDir("a/b")
	assert.Nil(t, err)
	assert.True(t, ok)

	assert.Nil(t, wrapper.Cd("/a/b/c"))

	ok, err = wrapper.IsDir("d")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.IsDir("../e")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = wrapper.IsDir("x")
	assert.EqualError(t, err, "Error: key not found")
	assert.False(t, ok)
}
