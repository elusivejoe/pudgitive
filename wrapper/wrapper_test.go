package wrapper

import (
	"testing"

	"github.com/elusivejoe/pudgitive/database"

	"github.com/stretchr/testify/assert"
)

func TestNavigation(t *testing.T) {
	wrapper := NewWrapper(database.NewDatabase("../tmp/testdb"))

	_, err := wrapper.Ls("/")
	assert.Nil(t, err)

	_, err = wrapper.Exists("test")
	assert.Nil(t, err)

	_, err = wrapper.IsDir("test")
	assert.Nil(t, err)

	err = wrapper.Cd("test")
	assert.Nil(t, err)
}

func TestAlteration(t *testing.T) {
	wrapper := NewWrapper(database.NewDatabase("../tmp/testdb"))

	exists, err := wrapper.Exists("test")
	assert.Nil(t, err)
	assert.False(t, exists)

	wrapper.MkDir("test", false)

	exists, err = wrapper.Exists("test")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.Cd("test")
	assert.Nil(t, err)

	err = wrapper.Mv("test", "test_1")
	assert.Nil(t, err)

	exists, err = wrapper.Exists("test")
	assert.Nil(t, err)
	assert.False(t, exists)

	exists, err = wrapper.Exists("test_1")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.Cd("test_1")
	assert.Nil(t, err)

	file, err := wrapper.MkFile("test_file.txt")
	assert.Nil(t, err)
	assert.False(t, file.Attributes().IsDir())
	assert.Equal(t, file.Name(), "test_file.txt")

	err = wrapper.Mv("test_file.txt", "test.txt")
	assert.Nil(t, err)

	exists, err = wrapper.Exists("test.txt")
	assert.Nil(t, err)
	assert.True(t, exists)

	file, err = wrapper.MkFile("test_1.txt")
	assert.Nil(t, err)
	assert.False(t, file.Attributes().IsDir())
	assert.Equal(t, file.Name(), "test_1.txt")

	err = wrapper.Cd("..")
	assert.Nil(t, err)

	exists, err = wrapper.Exists("test_1")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.RmDir("test_1", false)
	assert.NotNil(t, err)

	err = wrapper.RmDir("test_1", true)
	assert.Nil(t, err)

	exists, err = wrapper.Exists("test_1")
	assert.Nil(t, err)
	assert.False(t, exists)
}
