package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPhoneNumber(t *testing.T) {
	assert.Equal(t, true, IsPhoneNumber("13112312312"))
	assert.Equal(t, false, IsPhoneNumber("11112312312"))
	assert.Equal(t, false, IsPhoneNumber("foo"))
}
