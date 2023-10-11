package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneStr(t *testing.T) {
	sourceStr := "foo"
	actual := CloneStr(sourceStr)
	assert.Equal(t, sourceStr, actual)
}
