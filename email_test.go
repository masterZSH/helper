package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmail(t *testing.T) {
	assert.Equal(t, true, IsEmail("foo@bar.com"))
	assert.Equal(t, true, IsEmail("foo_s@bar.com"))
	assert.Equal(t, false, IsEmail("foo.s@bar.com"))
	assert.Equal(t, false, IsEmail("foo_sbar.com"))
}
