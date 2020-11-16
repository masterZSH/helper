package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sArr []string = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func TestByteCount(t *testing.T) {
	testCommonSize(t, B)
	testCommonSize(t, KB)
	testCommonSize(t, MB)
	testCommonSize(t, GB)
	testCommonSize(t, TB)
	testCommonSize(t, PB)
}

func TestGetPhoneCipherText(t *testing.T) {
	assert.Equal(t, "", GetPhoneCipherText("aaa"))
	assert.Equal(t, "131****2312", GetPhoneCipherText("13112312312"))
}

func testCommonSize(t *testing.T, s ByteSize) {
	ns := s + 1
	assert.Equal(t, "1.0 "+sArr[ns], ByteCount(1024, s))
	n := 1.1
	n = n * unit
	ex := int(n)
	assert.Equal(t, "1.1 "+sArr[ns], ByteCount(ex, s))
}
