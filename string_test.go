package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsChinese(t *testing.T) {
	assert.Equal(t, true, IsChinese("测试喉中文"))

	assert.Equal(t, false, IsChinese("测试foo中文"))
	assert.Equal(t, false, IsChinese("foo"))
	assert.Equal(t, false, IsChinese("---中文"))
}
