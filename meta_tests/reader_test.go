package meta_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/meta"

	"github.com/stretchr/testify/assert"

	"github.com/elusivejoe/pudgitive/testutils"
)

func TestReadMeta(t *testing.T) {
	wrapper, db := testutils.NewWrapper(t)

	wrapper.InitRoot("read_meta")
	wrapper.OpenRoot("read_meta")

	metaInfo, err := meta.ReadMeta(db, "abc")
	assert.NotNil(t, err)
	assert.Empty(t, metaInfo)

	wrapper.MkDir("/a/b/cd/ef/the dir/another dir")
	metaInfo, err = meta.ReadMeta(db, "read_meta/a/b/cd/ef/the dir")
	assert.Nil(t, err)
	assert.True(t, metaInfo.Attrs.IsDir)
	assert.Equal(t, "the dir", metaInfo.Name)
}
