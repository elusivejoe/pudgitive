package wrapper

import (
	"path"
	"testing"

	"github.com/elusivejoe/pudgitive/database"

	"github.com/stretchr/testify/assert"
)

func createWrapper(t *testing.T) *Wrapper {
	database := database.NewDatabase(path.Join(t.TempDir(), "testdb"))

	t.Cleanup(func() {
		database.Close()
	})

	return NewWrapper(database)
}

func TestEndpointAssembling(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_endpoint"))
	assert.Nil(t, wrapper.OpenRoot("test_endpoint"))

	path, err := NewCheckedPath("/")
	assert.Nil(t, err)
	assert.True(t, path.isAbs)

	endpoint, err := wrapper.assembleEndpoint(path)
	assert.Nil(t, err)
	assert.Equal(t, "test_endpoint", endpoint)
}

func TestPositionTrimming(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_pos_trim"))
	assert.Nil(t, wrapper.OpenRoot("test_pos_trim"))

	path, _ := NewCheckedPath("/")
	endpoint, _ := wrapper.assembleEndpoint(path)
	endpoint += "/dir1/dir2"
	assert.Equal(t, "test_pos_trim/dir1/dir2", endpoint)

	trimmed := wrapper.trimPosition(endpoint)
	assert.Equal(t, "/dir1/dir2", trimmed)
}
