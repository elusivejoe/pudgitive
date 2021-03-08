package fileio

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elusivejoe/pudgitive/testutils"
)

func TestWriter(t *testing.T) {
	_, db := testutils.NewWrapper(t)
	writer, err := NewWriter(db, 0, 32)
	assert.Nil(t, err)
	assert.NotNil(t, writer)

	n, err := writer.Write([]byte("test string"))
	assert.Nil(t, err)
	assert.Equal(t, 11, n)
}
