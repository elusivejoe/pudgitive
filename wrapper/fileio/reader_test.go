package fileio

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	_, db := testutils.NewWrapper(t)
	writer, err := NewReader(db, 0)
	assert.Nil(t, err)
	assert.NotNil(t, writer)

	out := make([]byte, 64)
	n, err := writer.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 11, n)
	assert.Equal(t, "test string", string(out))
}
