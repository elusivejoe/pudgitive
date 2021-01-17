package wrapper_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmin(t *testing.T) {
	wrapper := createWrapper(t)

	err := wrapper.InitRoot("")
	assert.EqualError(t, err, "root name cannot be empty")

	rootName := "Awesome File System!"

	err = wrapper.InitRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")

	err = wrapper.InitRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")

	err = wrapper.InitRoot(rootName)

	assert.EqualError(t, err, "wrapper: root 'Awesome File System!' already exists")

	err = wrapper.OpenRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot("Some Other Root")
	assert.EqualError(t, err, "wrapper: unable to find root 'Some Other Root'")
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot("Another Root")
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), rootName)

	err = wrapper.DeleteRoot(rootName)
	assert.Nil(t, err)
	assert.Equal(t, wrapper.CurrentRoot(), "")
}
