package fileio

import (
	"testing"

	"github.com/elusivejoe/pudgitive/testutils"
	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	_, db := testutils.NewWrapper(t)

	writer, _ := NewWriter(db, 0, 5)
	writer.Write([]byte("t"))
	writer.Write([]byte("es"))
	writer.Write([]byte("t stri"))
	writer.Write([]byte("ng"))

	reader, err := NewReader(db, 0)
	assert.Nil(t, err)
	assert.NotNil(t, reader)

	out := make([]byte, 64)
	n, err := reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 11, n)
	assert.Equal(t, "test string", string(out[:n]))

	out = make([]byte, 4)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
	assert.Equal(t, "", string(out[:n]))
}

func TestReadParts(t *testing.T) {
	_, db := testutils.NewWrapper(t)

	writer, _ := NewWriter(db, 0, 5)
	writer.Write([]byte("t"))
	writer.Write([]byte("es"))
	writer.Write([]byte("t stri"))
	writer.Write([]byte("ng"))

	reader, err := NewReader(db, 0)
	assert.Nil(t, err)
	assert.NotNil(t, reader)

	out := make([]byte, 3)
	n, err := reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, "tes", string(out[:n]))

	out = make([]byte, 3)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, "t s", string(out[:n]))

	out = make([]byte, 1)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, "t", string(out[:n]))

	out = make([]byte, 4)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "ring", string(out[:n]))
	assert.True(t, reader.Eof())

	out = make([]byte, 10)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}

func TestReadAligned(t *testing.T) {
	_, db := testutils.NewWrapper(t)

	writer, _ := NewWriter(db, 0, 5)
	writer.Write([]byte("test "))
	writer.Write([]byte("str"))
	writer.Write([]byte("in"))
	writer.Write([]byte("g abc"))

	reader, err := NewReader(db, 0)
	assert.Nil(t, err)
	assert.NotNil(t, reader)

	out := make([]byte, 5)
	n, err := reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "test ", string(out[:n]))

	out = make([]byte, 3)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, "str", string(out[:n]))

	out = make([]byte, 2)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, "in", string(out[:n]))

	out = make([]byte, 5)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "g abc", string(out[:n]))
	assert.True(t, reader.Eof())

	out = make([]byte, 10)
	n, err = reader.Read(out)
	assert.Nil(t, err)
	assert.Equal(t, 0, n)
}
