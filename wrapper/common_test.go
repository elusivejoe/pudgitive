package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionTrimming(t *testing.T) {
	root := "test_pos_trim"
	pwd := "a/b/c"

	assert.Equal(t, "/a/b/c", trimPosition(root, pwd, "test_pos_trim/a/b/c", true))
	assert.Equal(t, "/d/e", trimPosition(root, pwd, "test_pos_trim/a/b/c/d/e", false))
}
