package utils_tests

import (
	"testing"

	"github.com/elusivejoe/pudgitive/utils"

	"github.com/stretchr/testify/assert"
)

func TestDbKeyTrim(t *testing.T) {
	root := "test_pos_trim"
	pwd := "a/b/c"

	assert.Equal(t, "/a/b/c", utils.TrimDbKey(root, pwd, "test_pos_trim/a/b/c", true))
	assert.Equal(t, "/d/e", utils.TrimDbKey(root, pwd, "test_pos_trim/a/b/c/d/e", false))
}
