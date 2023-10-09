package helper

import (
	"crypto/rand"
	"unsafe"
)

// GenerateCaptcha 生成数字验证码
// GenerateCaptcha(4) => 1253
// GenerateCaptcha(4) => 142536
func GenerateCaptcha(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	for k, v := range b {
		b[k] = (v & 9) + 48
	}
	return *(*string)(unsafe.Pointer(&b))
}
