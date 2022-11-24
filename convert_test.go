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
