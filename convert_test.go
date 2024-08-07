package helper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilSliceToEmptySlice(t *testing.T) {
	var foo []interface{}
	r, err := json.Marshal(foo)
	assert.Nil(t, err)
	assert.Equal(t, "null", r)
	result := NilSliceToEmptySlice(foo)
	r, err = json.Marshal(result)
	assert.Nil(t, err)
	assert.Equal(t, "[]", r)
}

// BenchmarkSliceByteToString-8   	1000000000	         0.2276 ns/op	       0 B/op	       0 allocs/op
// BenchmarkStringToSliceByte-8   	1000000000	         0.2293 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSliceByteToString(b *testing.B) {
	bs := []byte("123")
	for i := 0; i < b.N; i++ {
		SliceByteToString(bs)
	}
}

func BenchmarkStringToSliceByte(b *testing.B) {
	str := "123"
	for i := 0; i < b.N; i++ {
		StringToSliceByte(str)
	}
}

func BenchmarkSliceByteToStringDefault(b *testing.B) {
	bs := []byte("123")
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}

// 5.139 ns/op
func BenchmarkStringToSliceByteDefault(b *testing.B) {
	str := "123"
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}

func BenchmarkStringToSliceByteV2(b *testing.B) {
	str := "123"
	for i := 0; i < b.N; i++ {
		StringToSliceByteV2(str)
	}
}
