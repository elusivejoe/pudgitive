package wrapper

import (
	"path"
	"testing"

	"github.com/elusivejoe/pudgitive/pathUtils"

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

	path := pathUtils.NewNormPath("/")
	assert.True(t, path.IsAbs())

	endpoint := assembleEndpoint(wrapper, path)
	assert.Equal(t, "test_endpoint", endpoint)
}

func TestPositionTrimming(t *testing.T) {
	wrapper := createWrapper(t)

	assert.Nil(t, wrapper.InitRoot("test_pos_trim"))
	assert.Nil(t, wrapper.OpenRoot("test_pos_trim"))

	path := pathUtils.NewNormPath("/")
	endpoint := assembleEndpoint(wrapper, path)
	endpoint += "/dir1/dir2"
	assert.Equal(t, "test_pos_trim/dir1/dir2", endpoint)

	trimmed := trimPosition(wrapper, endpoint)
	assert.Equal(t, "/dir1/dir2", trimmed)
}
