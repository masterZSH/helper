package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCaptcha(t *testing.T) {
	testCases := []struct {
		val uint8
		len int
	}{
		{
			val: 4,
			len: 4,
		},
		{
			val: 6,
			len: 6,
		},
		{
			val: 8,
			len: 8,
		},
	}

	for _, v := range testCases {
		result := GenerateCaptcha(int(v.val))
		assert.Equal(t, v.len, len(result))
	}

}

func BenchmarkGenCaptcha(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateCaptcha(4)
	}
}
