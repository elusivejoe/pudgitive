package wrapper_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"

	"github.com/stretchr/testify/assert"
)

func TestAdmin(t *testing.T) {
	wrapper, _ := testutils.NewWrapper(t)

	err := wrapper.InitRoot("")
	assert.EqualError(t, err, "root name cannot be empty")

	rootName := "Awesome File System!"

	err = wrapper.InitRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, "test_root", wrapper.CurrentRoot())

	err = wrapper.InitRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, "test_root", wrapper.CurrentRoot())

	err = wrapper.InitRoot(rootName)

	assert.EqualError(t, err, "root 'Awesome File System!' already exists")

	err = wrapper.OpenRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, rootName, wrapper.CurrentRoot())

	err = wrapper.DeleteRoot("Some Other Root")
	assert.EqualError(t, err, "unable to find root 'Some Other Root'")
	assert.Equal(t, rootName, wrapper.CurrentRoot())

	err = wrapper.DeleteRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, rootName, wrapper.CurrentRoot())

	err = wrapper.DeleteRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, "", wrapper.CurrentRoot())
}
