package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elusivejoe/pudgitive/database"
)

func TestNavigation(t *testing.T) {
	wrapper := NewWrapper(&database.DummyDatabase{})

	_, err := wrapper.ls("/")
	assert.Nil(t, err)

	_, err = wrapper.exists("test")
	assert.Nil(t, err)

	_, err = wrapper.isDir("test")
	assert.Nil(t, err)

	err = wrapper.cd("test")
	assert.Nil(t, err)
}

func TestAlteration(t *testing.T) {
	wrapper := NewWrapper(&database.DummyDatabase{})

	exists, err := wrapper.exists("test")
	assert.Nil(t, err)
	assert.False(t, exists)

	wrapper.mkDir("test", false)

	exists, err = wrapper.exists("test")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.cd("test")
	assert.Nil(t, err)

	err = wrapper.mv("test", "test_1")
	assert.Nil(t, err)

	exists, err = wrapper.exists("test")
	assert.Nil(t, err)
	assert.False(t, exists)

	exists, err = wrapper.exists("test_1")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.cd("test_1")
	assert.Nil(t, err)

	file, err := wrapper.mkFile("test_file.txt")
	assert.Nil(t, err)
	assert.False(t, file.IsDir())
	assert.Equal(t, file.Name(), "test_file.txt")

	err = wrapper.mv("test_file.txt", "test.txt")
	assert.Nil(t, err)

	exists, err = wrapper.exists("test.txt")
	assert.Nil(t, err)
	assert.True(t, exists)

	file, err = wrapper.mkFile("test_1.txt")
	assert.Nil(t, err)
	assert.False(t, file.IsDir())
	assert.Equal(t, file.Name(), "test_1.txt")

	err = wrapper.cd("..")
	assert.Nil(t, err)

	exists, err = wrapper.exists("test_1")
	assert.Nil(t, err)
	assert.True(t, exists)

	err = wrapper.rmDir("test_1", false)
	assert.NotNil(t, err)

	err = wrapper.rmDir("test_1", true)
	assert.Nil(t, err)

	exists, err = wrapper.exists("test_1")
	assert.Nil(t, err)
	assert.False(t, exists)
}
